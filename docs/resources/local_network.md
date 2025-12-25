---
page_title: "prodata_local_network Resource - ProData Provider"
description: |-
  Manages a ProData local network.
---

# prodata_local_network (Resource)

Manages a ProData local network.

~> **Note:** Only the `name` attribute can be updated in-place. Changing `cidr`, `gateway`, `region`, or `project_id` will force the creation of a new local network (destroy and recreate).

## Example Usage

### Basic Usage (Using Provider Defaults)

```terraform
# Uses region and project_id from provider configuration
resource "prodata_local_network" "example" {
  name    = "my-network"
  cidr    = "10.0.0.0/24"
  gateway = "10.0.0.1"
}

output "network_id" {
  value = prodata_local_network.example.id
}
```

### With Explicit Region and Project

```terraform
resource "prodata_local_network" "example" {
  region     = "UZ5"
  project_id = 89
  name       = "my-network"
  cidr       = "10.0.0.0/24"
  gateway    = "10.0.0.1"
}
```

### Renaming a Local Network (In-Place Update)

```terraform
# Changing only the name will update the local network in-place without recreation
resource "prodata_local_network" "example" {
  name    = "renamed-network"  # Changed from "my-network" - updates in-place
  cidr    = "10.0.0.0/24"
  gateway = "10.0.0.1"
}
```

### Multiple Networks with Different CIDR Blocks

```terraform
resource "prodata_local_network" "frontend" {
  name    = "frontend-network"
  cidr    = "10.0.1.0/24"
  gateway = "10.0.1.1"
}

resource "prodata_local_network" "backend" {
  name    = "backend-network"
  cidr    = "10.0.2.0/24"
  gateway = "10.0.2.1"
}

resource "prodata_local_network" "database" {
  name    = "database-network"
  cidr    = "10.0.3.0/24"
  gateway = "10.0.3.1"
}
```

## Schema

### Required

- `name` (String) The name of the local network. **This is the only attribute that can be updated in-place.**
- `cidr` (String) The CIDR block for the local network (e.g., 10.0.0.0/24). Changing this forces a new resource.
- `gateway` (String) The gateway IP address for the local network (e.g., 10.0.0.1). Changing this forces a new resource.

### Optional

- `region` (String) Region where the local network will be created (e.g., UZ5). If not specified, uses the provider's default region. Changing this forces a new resource.
- `project_id` (Number) Project ID where the local network will be created. If not specified, uses the provider's default project_id. Changing this forces a new resource.

### Read-Only

- `id` (Number) The unique identifier of the local network.

## Import

Local networks cannot be imported as the API does not provide sufficient information to reconstruct the Terraform state.
