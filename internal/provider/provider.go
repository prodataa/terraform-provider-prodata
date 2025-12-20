package provider

import (
	"context"
	"os"

	"terraform-provider-prodata/internal/client"
	"terraform-provider-prodata/internal/provider/datasources"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &ProDataProvider{}

type ProDataProvider struct {
	version string
}

type ProDataProviderModel struct {
	APIBaseURL   types.String `tfsdk:"api_base_url"`
	APIKeyID     types.String `tfsdk:"api_key_id"`
	APISecretKey types.String `tfsdk:"api_secret_key"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ProDataProvider{version: version}
	}
}

func (p *ProDataProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "prodata"
	resp.Version = p.version
}

func (p *ProDataProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage ProData Cloud resources.",
		Attributes: map[string]schema.Attribute{
			"api_base_url": schema.StringAttribute{
				MarkdownDescription: "ProData API base URL. Env: `PRODATA_API_BASE_URL`.",
				Optional:            true,
			},
			"api_key_id": schema.StringAttribute{
				MarkdownDescription: "API Key ID. Env: `PRODATA_API_KEY_ID`.",
				Optional:            true,
			},
			"api_secret_key": schema.StringAttribute{
				MarkdownDescription: "API Secret Key. Env: `PRODATA_API_SECRET`.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *ProDataProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ProDataProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build config with env fallbacks
	cfg := client.Config{
		APIBaseURL:   envOrValue("PRODATA_API_BASE_URL", data.APIBaseURL),
		APIKeyID:     envOrValue("PRODATA_API_KEY_ID", data.APIKeyID),
		APISecretKey: envOrValue("PRODATA_API_SECRET", data.APISecretKey),
	}

	// Validate required fields
	if cfg.APIBaseURL == "" {
		resp.Diagnostics.AddAttributeError(path.Root("api_base_url"), "Missing API Base URL",
			"Set api_base_url in config or PRODATA_API_BASE_URL environment variable.")
	}
	if cfg.APIKeyID == "" {
		resp.Diagnostics.AddAttributeError(path.Root("api_key_id"), "Missing API Key ID",
			"Set api_key_id in config or PRODATA_API_KEY_ID environment variable.")
	}
	if cfg.APISecretKey == "" {
		resp.Diagnostics.AddAttributeError(path.Root("api_secret_key"), "Missing API Secret",
			"Set api_secret_key in config or PRODATA_API_SECRET environment variable.")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := client.New(cfg)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create client", err.Error())
		return
	}

	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *ProDataProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *ProDataProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewImageDataSource,
	}
}

// envOrValue returns the config value if set, otherwise falls back to env var
func envOrValue(envKey string, value types.String) string {
	if !value.IsNull() {
		return value.ValueString()
	}
	return os.Getenv(envKey)
}
