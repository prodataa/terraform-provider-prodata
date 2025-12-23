terraform {
  required_providers {
    prodata = {
      source = "pro-data/prodata"
    }
  }
}

provider "prodata" {
  # Configuration options
}

# List all available images
data "prodata_images" "all" {}

output "all_images" {
  value = data.prodata_images.all.images
}

output "image_count" {
  value = length(data.prodata_images.all.images)
}

# Filter custom images using locals
locals {
  custom_images = [for img in data.prodata_images.all.images : img if img.is_custom]
  os_templates  = [for img in data.prodata_images.all.images : img if !img.is_custom]
}

output "custom_image_names" {
  value = [for img in local.custom_images : img.name]
}

output "os_template_slugs" {
  value = [for img in local.os_templates : img.slug]
}

# List images in a specific region
data "prodata_images" "kz_images" {
  region = "KZ-1"
}
