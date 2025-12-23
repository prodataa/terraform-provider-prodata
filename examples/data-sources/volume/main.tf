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

# Lookup a volume by ID
data "prodata_volume" "main" {
  id = 649
}

output "volume" {
  value = {
    id          = data.prodata_volume.main.id
    name        = data.prodata_volume.main.name
    type        = data.prodata_volume.main.type
    size        = data.prodata_volume.main.size
    in_use      = data.prodata_volume.main.in_use
    attached_id = data.prodata_volume.main.attached_id
  }
}

# Check if volume is attached
output "volume_status" {
  value = data.prodata_volume.main.in_use ? "attached to instance ${data.prodata_volume.main.attached_id}" : "available"
}
