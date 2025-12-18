package datasources

import (
	"context"
	"fmt"

	"terraform-provider-prodata/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ImageDataSource{}

type ImageDataSource struct {
	client *client.Client
}

type ImageDataSourceModel struct {
	// Input - one of these required
	Name types.String `tfsdk:"name"`
	Slug types.String `tfsdk:"slug"`

	// Computed output
	ID       types.Int64 `tfsdk:"id"`
	IsCustom types.Bool  `tfsdk:"is_custom"`
}

func ProDataImageDataSource() datasource.DataSource {
	return &ImageDataSource{}
}

func (d *ImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image"
}

func (d *ImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get information about a ProData image (OS template or custom image) for use in other resources.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the image. Used for custom images lookup.",
				Optional:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "The slug of the image. Used for OS template lookup (e.g., `ubuntu-22.04`, `debian-11`).",
				Optional:            true,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the image.",
				Computed:            true,
			},
			"is_custom": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a custom image (`true`) or OS template (`false`).",
				Computed:            true,
			},
		},
	}
}

func (d *ImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	d.client = c
}

func (d *ImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ImageDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading image data source", map[string]interface{}{
		"name": data.Name.ValueString(),
		"slug": data.Slug.ValueString(),
	})

	// Call API based on what's provided
	var image *client.Image
	var err error

	if !data.Slug.IsNull() {
		tflog.Debug(ctx, "Looking up image by slug", map[string]interface{}{"slug": data.Slug.ValueString()})
		image, err = d.client.GetImageBySlug(ctx, data.Slug.ValueString())
	} else if !data.Name.IsNull() {
		tflog.Debug(ctx, "Looking up image by name", map[string]interface{}{"name": data.Name.ValueString()})
		image, err = d.client.GetImageByName(ctx, data.Name.ValueString())
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Image",
			fmt.Sprintf("Could not read image: %s", err.Error()),
		)
		return
	}

	if image == nil {
		resp.Diagnostics.AddError(
			"Image Not Found",
			"No image found with the specified criteria.",
		)
		return
	}

	// Map response to model
	data.ID = types.Int64Value(image.ID)
	data.IsCustom = types.BoolValue(image.IsCustom)

	tflog.Debug(ctx, "Successfully read image", map[string]interface{}{
		"id":        image.ID,
		"is_custom": image.IsCustom,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
