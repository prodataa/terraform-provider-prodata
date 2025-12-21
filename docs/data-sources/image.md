---
page_title: "prodata_image Data Source - ProData Provider"
description: |-
  Lookup ProData images (OS templates or custom images) by slug or name.
---

# prodata_image (Data Source)

Lookup a ProData image by slug (OS templates) or name (custom images).

## Example Usage

### Lookup OS Template by Slug

```terraform
data "prodata_image" "ubuntu" {
  slug = "ubuntu-22.04"
}

output "ubuntu_image_id" {
  value = data.prodata_image.ubuntu.id
}
```

### Lookup Custom Image by Name

```terraform
data "prodata_image" "my_image" {
  name = "my-custom-image"
}

output "custom_image_id" {
  value = data.prodata_image.my_image.id
}
```

### With Region Override

```terraform
data "prodata_image" "debian" {
  slug   = "debian-12"
  region = "KZ-1"
}
```

## Schema

### Optional

Exactly one of `name` or `slug` must be specified.

- `name` (String) Image name. Use for looking up custom images. Conflicts with `slug`.
- `slug` (String) Image slug (e.g., `ubuntu-22.04`, `debian-12`). Use for looking up OS templates. Conflicts with `name`.
- `region` (String) Region ID override. If not specified, uses the provider's default region.
- `project_id` (Number) Project ID override. If not specified, uses the provider's default project ID.

### Read-Only

- `id` (Number) The unique identifier of the image.
- `is_custom` (Boolean) `true` if this is a custom image, `false` if it's an OS template.
