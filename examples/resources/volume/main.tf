terraform {
  required_providers {
    prodata = {
      source = "pro-data/prodata"
    }
  }
}

provider "prodata" {
  # Configuration options (region and project_id can be set here as defaults)
}

# Create an HDD volume using provider defaults for region and project_id
resource "prodata_volume" "main" {
  name = "terraform-volume"
  type = "HDD"
  size = 10
}

output "volume" {
  value = {
    id   = prodata_volume.main.id
    name = prodata_volume.main.name
    type = prodata_volume.main.type
    size = prodata_volume.main.size
  }
}

# Create an SSD volume with explicit region and project_id (overrides provider defaults)
resource "prodata_volume" "ssd" {
  region     = "UZ5"
  project_id = 89
  name       = "fast-storage"
  type       = "SSD"
  size       = 20
}
