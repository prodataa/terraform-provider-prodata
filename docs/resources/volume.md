---
page_title: "prodata_volume Resource - ProData Provider"
description: |-
  Manages a ProData volume.
---

# prodata_volume (Resource)

Manages a ProData volume.

~> **Note:** Only the `name` attribute can be updated in-place. Changing `type`, `size`, `region`, or `project_id` will force the creation of a new volume (destroy and recreate).

## Example Usage

### Basic Usage (Using Provider Defaults)

```terraform
# Uses region and project_id from provider configuration
resource "prodata_volume" "example" {
  name = "my-volume"
  type = "HDD"
  size = 10
}

output "volume_id" {
  value = prodata_volume.example.id
}
```

### With Explicit Region and Project

```terraform
resource "prodata_volume" "example" {
  region     = "UZ5"
  project_id = 89
  name       = "my-volume"
  type       = "HDD"
  size       = 10
}
```

### Renaming a Volume (In-Place Update)

```terraform
# Changing only the name will update the volume in-place without recreation
resource "prodata_volume" "example" {
  name = "renamed-volume"  # Changed from "my-volume" - updates in-place
  type = "HDD"
  size = 10
}
```

### SSD Volume

```terraform
resource "prodata_volume" "ssd" {
  name = "fast-storage"
  type = "SSD"
  size = 50
}
```

### Multiple Volumes

```terraform
resource "prodata_volume" "data" {
  name = "data-volume"
  type = "HDD"
  size = 100
}

resource "prodata_volume" "logs" {
  name = "logs-volume"
  type = "HDD"
  size = 50
}
```

## Schema

### Required

- `name` (String) The name of the volume. **This is the only attribute that can be updated in-place.**
- `type` (String) The type of the volume (HDD or SSD). Changing this forces a new resource.
- `size` (Number) The size of the volume in GB. Changing this forces a new resource.

### Optional

- `region` (String) Region where the volume will be created (e.g., UZ5). If not specified, uses the provider's default region. Changing this forces a new resource.
- `project_id` (Number) Project ID where the volume will be created. If not specified, uses the provider's default project_id. Changing this forces a new resource.

### Read-Only

- `id` (Number) The unique identifier of the volume.

## Import

Volumes cannot be imported as the API does not provide sufficient information to reconstruct the Terraform state.
