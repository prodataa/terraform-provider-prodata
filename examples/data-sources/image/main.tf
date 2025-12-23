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

# Lookup an OS template by slug
data "prodata_image" "debian" {
  slug = "debian-11"
}

output "debian_image_id" {
  value = data.prodata_image.debian.id
}

# Lookup a custom image by name
data "prodata_image" "custom" {
  name = "my-custom-image"
}

output "custom_image" {
  value = {
    id        = data.prodata_image.custom.id
    is_custom = data.prodata_image.custom.is_custom
  }
}

# Lookup image in a specific region
data "prodata_image" "ubuntu_kz" {
  slug   = "ubuntu-22.04"
  region = "KZ-1"
}
