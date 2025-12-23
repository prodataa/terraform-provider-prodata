---
page_title: "prodata_volume Data Source - ProData Provider"
description: |-
  Lookup a ProData volume by ID.
---

# prodata_volume (Data Source)

Lookup a ProData volume by its unique identifier.

## Example Usage

### Basic Usage

```terraform
data "prodata_volume" "example" {
  id = 285098
}

output "volume_name" {
  value = data.prodata_volume.example.name
}

output "volume_size" {
  value = data.prodata_volume.example.size
}
```

### With Region Override

```terraform
data "prodata_volume" "example" {
  id     = 285098
  region = "UZ-5"
}
```

### Check Attachment Status

```terraform
data "prodata_volume" "example" {
  id = 285098
}

output "is_attached" {
  value = data.prodata_volume.example.in_use
}

output "attached_instance" {
  value = data.prodata_volume.example.in_use ? data.prodata_volume.example.attached_id : null
}
```

## Schema

### Required

- `id` (Number) The unique identifier of the volume.

### Optional

- `region` (String) Region ID override. If not specified, uses the provider's default region.
- `project_id` (Number) Project ID override. If not specified, uses the provider's default project ID.

### Read-Only

- `name` (String) The name of the volume.
- `type` (String) The type of the volume (e.g., HDD, SSD).
- `size` (Number) The size of the volume in GB.
- `in_use` (Boolean) `true` if the volume is attached to an instance, `false` otherwise.
- `attached_id` (Number) The ID of the instance the volume is attached to, or `null` if not attached.
