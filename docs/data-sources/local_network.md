---
page_title: "prodata_local_network Data Source - ProData Provider"
description: |-
  Lookup a ProData local network by ID.
---

# prodata_local_network (Data Source)

Lookup a ProData local network by its unique identifier.

## Example Usage

### Basic Usage

```terraform
data "prodata_local_network" "example" {
  id = 114530
}

output "network_name" {
  value = data.prodata_local_network.example.name
}

output "network_cidr" {
  value = data.prodata_local_network.example.cidr
}
```

### With Region Override

```terraform
data "prodata_local_network" "example" {
  id     = 114530
  region = "UZ5"
}
```

### Check Link Status

```terraform
data "prodata_local_network" "example" {
  id = 114530
}

output "is_linked" {
  value = data.prodata_local_network.example.linked
}

output "network_info" {
  value = {
    name    = data.prodata_local_network.example.name
    cidr    = data.prodata_local_network.example.cidr
    gateway = data.prodata_local_network.example.gateway
    linked  = data.prodata_local_network.example.linked
  }
}
```

## Schema

### Required

- `id` (Number) The unique identifier of the local network.

### Optional

- `region` (String) Region ID override. If not specified, uses the provider's default region.
- `project_id` (Number) Project ID override. If not specified, uses the provider's default project ID.

### Read-Only

- `name` (String) The name of the local network.
- `cidr` (String) The CIDR block of the local network.
- `gateway` (String) The gateway IP address of the local network.
- `linked` (Boolean) `true` if the local network is linked to an instance, `false` otherwise.
