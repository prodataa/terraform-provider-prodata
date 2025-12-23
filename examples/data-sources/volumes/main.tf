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

# List all volumes
data "prodata_volumes" "all" {}

output "all_volumes" {
  value = data.prodata_volumes.all.volumes
}

output "volume_count" {
  value = length(data.prodata_volumes.all.volumes)
}

# Filter volumes using locals
locals {
  attached_volumes  = [for vol in data.prodata_volumes.all.volumes : vol if vol.in_use]
  available_volumes = [for vol in data.prodata_volumes.all.volumes : vol if !vol.in_use]
  hdd_volumes       = [for vol in data.prodata_volumes.all.volumes : vol if vol.type == "HDD"]
  ssd_volumes       = [for vol in data.prodata_volumes.all.volumes : vol if vol.type == "SSD"]
}

output "attached_volume_ids" {
  value = [for vol in local.attached_volumes : vol.id]
}

output "available_volume_names" {
  value = [for vol in local.available_volumes : vol.name]
}

output "total_hdd_storage_gb" {
  value = sum([for vol in local.hdd_volumes : vol.size])
}

output "total_ssd_storage_gb" {
  value = sum([for vol in local.ssd_volumes : vol.size])
}
