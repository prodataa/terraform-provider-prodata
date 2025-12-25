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

# Create a local network using provider defaults for region and project_id
resource "prodata_local_network" "main" {
  name    = "terraform-network"
  cidr    = "10.0.0.0/24"
  gateway = "10.0.0.1"
}

output "network" {
  value = {
    id      = prodata_local_network.main.id
    name    = prodata_local_network.main.name
    cidr    = prodata_local_network.main.cidr
    gateway = prodata_local_network.main.gateway
  }
}

# Create a local network with explicit region and project_id (overrides provider defaults)
resource "prodata_local_network" "backend" {
  region     = "UZ5"
  project_id = 89
  name       = "backend-network"
  cidr       = "10.0.1.0/24"
  gateway    = "10.0.1.1"
}