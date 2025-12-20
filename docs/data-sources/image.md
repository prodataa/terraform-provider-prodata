---
page_title: "prodata_image Data Source"
description: |-
  Lookup ProData OS templates and custom images.
---

# prodata_image (Data Source)

Lookup ProData images by slug (OS templates) or name (custom images).

## Example Usage

### OS Template by Slug

```terraform
data "prodata_image" "ubuntu" {
  slug = "ubuntu-22.04"
}

output "image_id" {
  value = data.prodata_image.ubuntu.id
}
```

### Custom Image by Name

```terraform
data "prodata_image" "my_image" {
  name = "my-custom-image"
}

output "image_info" {
  value = {
    id        = data.prodata_image.my_image.id
    is_custom = data.prodata_image.my_image.is_custom
  }
}
```

## Schema

### Optional

You must specify exactly one of the following:

- `name` (String) Image name for custom images. Conflicts with `slug`.
- `slug` (String) Image slug for OS templates (e.g., `ubuntu-22.04`, `debian-11`). Conflicts with `name`.

### Read-Only

- `id` (Number) Image ID.
- `is_custom` (Boolean) Whether this is a custom image (`true`) or OS template (`false`).

## Need Help?

- **Help Desk**: [helpdesk.pro-data.tech](https://helpdesk.pro-data.tech)
- **Telegram**: [@PRO_DATA_Support_Bot](https://t.me/PRO_DATA_Support_Bot)
