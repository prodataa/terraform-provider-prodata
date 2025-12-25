---
page_title: "prodata_local_networks Data Source - ProData Provider"
description: |-
  List all available ProData local networks.
---

# prodata_local_networks (Data Source)

List all available ProData local networks in a project.

## Example Usage

### List All Local Networks

```terraform
data "prodata_local_networks" "all" {}

output "all_networks" {
  value = data.prodata_local_networks.all.local_networks
}
```

### With Region Override

```terraform
data "prodata_local_networks" "uz_networks" {
  region = "UZ5"
}

output "uz_network_count" {
  value = length(data.prodata_local_networks.uz_networks.local_networks)
}
```

### Filter Linked Networks with Local

```terraform
data "prodata_local_networks" "all" {}

locals {
  linked_networks   = [for net in data.prodata_local_networks.all.local_networks : net if net.linked]
  unlinked_networks = [for net in data.prodata_local_networks.all.local_networks : net if !net.linked]
}

output "linked_network_names" {
  value = [for net in local.linked_networks : net.name]
}

output "unlinked_network_names" {
  value = [for net in local.unlinked_networks : net.name]
}
```

### Get Network by Name

```terraform
data "prodata_local_networks" "all" {}

locals {
  backend_network = [for net in data.prodata_local_networks.all.local_networks : net if net.name == "backend-network"][0]
}

output "backend_network_id" {
  value = local.backend_network.id
}
```

## Schema

### Optional

- `region` (String) Region ID override. If not specified, uses the provider's default region.
- `project_id` (Number) Project ID override. If not specified, uses the provider's default project ID.

### Read-Only

- `local_networks` (List of Object) List of available local networks. Each local network has the following attributes:
  - `id` (Number) The unique identifier of the local network.
  - `name` (String) The name of the local network.
  - `cidr` (String) The CIDR block of the local network.
  - `gateway` (String) The gateway IP address of the local network.
  - `linked` (Boolean) `true` if the local network is linked to an instance, `false` otherwise.
