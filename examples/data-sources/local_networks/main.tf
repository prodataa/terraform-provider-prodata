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

# List all local networks
data "prodata_local_networks" "all" {}

output "all_networks" {
  value = data.prodata_local_networks.all.local_networks
}

output "network_count" {
  value = length(data.prodata_local_networks.all.local_networks)
}

# Filter networks using locals
locals {
  linked_networks   = [for net in data.prodata_local_networks.all.local_networks : net if net.linked]
  unlinked_networks = [for net in data.prodata_local_networks.all.local_networks : net if !net.linked]
}

output "linked_network_ids" {
  value = [for net in local.linked_networks : net.id]
}

output "unlinked_network_names" {
  value = [for net in local.unlinked_networks : net.name]
}
