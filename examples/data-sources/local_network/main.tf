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

# Lookup a local network by ID
data "prodata_local_network" "main" {
  id = 114530
}

output "network" {
  value = {
    id      = data.prodata_local_network.main.id
    name    = data.prodata_local_network.main.name
    cidr    = data.prodata_local_network.main.cidr
    gateway = data.prodata_local_network.main.gateway
    linked  = data.prodata_local_network.main.linked
  }
}

# Check if network is linked
output "network_status" {
  value = data.prodata_local_network.main.linked ? "linked to instance" : "available"
}
