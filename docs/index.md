---
page_title: "ProData Provider"
description: |-
  The ProData provider enables seamless infrastructure management for ProData Cloud resources through Terraform.
---

# ProData Provider

The ProData provider allows you to manage your ProData Cloud infrastructure using Terraform's declarative configuration language. With this provider, you can automate the provisioning, configuration, and lifecycle management of ProData resources.

## Getting Started

To use the ProData provider, you'll need:

- A ProData Cloud account
- API credentials (API Access Key and Secret Key)
- Your project ID and preferred region

## Quick Start Example

```terraform
terraform {
  required_providers {
    prodata = {
      source  = "prodata-cloud/prodata"
      version = "~> 1.0"
    }
  }
}

provider "prodata" {
  # Base URL varies by location:
  # Uzbekistan: https://my.pro-data.tech
  # Kazakhstan: https://kz-1.pro-data.tech
  api_base_url   = "https://my.pro-data.tech"

  api_key_id     = var.prodata_api_key_id
  api_secret_key = var.prodata_api_secret_key

  region         = var.prodata_region
  project_id     = var.prodata_project_id
}

```

## Authentication

The ProData provider supports two authentication methods:

### Method 1: Provider Configuration Block

Configure credentials directly in your Terraform configuration:

```terraform
provider "prodata" {
  api_base_url   = "https://my.pro-data.tech"
  api_key_id     = "ak_UCSDcaxx..."
  api_secret_key = "sc_Xxdsfzzc..."
  region         = "UZ-5"
  project_id     = "123"
}
```

### Method 2: Environment Variables (Recommended)

Set credentials using environment variables for enhanced security:

```bash
export PRODATA_API_BASE_URL="https://my.pro-data.tech"
export PRODATA_API_KEY_ID="your-api-key-id"
export PRODATA_API_SECRET_KEY="your-api-secret-key"
export PRODATA_REGION="UZ-5"
export PRODATA_PROJECT_ID="your-project-id"
```

Then use the provider without explicit credentials:

```terraform
provider "prodata" {
  # Configuration will be loaded from environment variables
}
```

**Best Practice:** Use environment variables or secret management tools (like HashiCorp Vault) to avoid hardcoding sensitive credentials in your Terraform files.

## Configuration Reference

The following arguments are supported in the provider configuration:

### Required Arguments

- **`api_base_url`** (String) - The base URL for the ProData API endpoint.
  *Environment variable:* `PRODATA_API_BASE_URL`
  *Example:* `https://my.pro-data.tech`

- **`api_key_id`** (String) - Your ProData API Key ID used for authentication.
  *Environment variable:* `PRODATA_API_KEY_ID`
  *Example:* `ak_UCSDcax...`

- **`api_secret_key`** (String, Sensitive) - Your ProData API Secret Key used for authentication.
  *Environment variable:* `PRODATA_API_SECRET_KEY`
  *Security Note:* This value is marked as sensitive and will not appear in logs.

- **`region`** (String) - The ProData Cloud region where resources will be provisioned.
  *Environment variable:* `PRODATA_REGION`
  *Available regions:* `KZ-1`, `UZ-5`, `UZ-3`

- **`project_id`** (String) - The unique identifier for your ProData project.
  *Environment variable:* `PRODATA_PROJECT_ID`
  *Example:* `123`

## Obtaining API Credentials

To generate API credentials for the ProData provider:

1. Log in to your ProData Cloud console
2. Navigate to **Account** → **Access Keys**
3. Click **Generate Key**
4. Copy your API Key ID and Secret Key (the secret will only be shown once)
5. Store your credentials securely

## Regional Availability

ProData Cloud is available in the following regions:

| Region Code | Location   | API Endpoint                 |
| ----------- | ---------- | ---------------------------- |
| `UZ-5`      | Uzbekistan | `https://my.pro-data.tech`   |
| `UZ-3`      | Uzbekistan | `https://my.pro-data.tech`   |
| `KZ-1`      | Kazakhstan | `https://kz-1.pro-data.tech` |

## Support and Resources

- **Support Portal Help Desk:** [https://helpdesk.pro-data.tech/](https://helpdesk.pro-data.tech)
- **Support Telegram bot:** [https://t.me/PRO_DATA_Support_Bot](https://t.me/PRO_DATA_Support_Bot)

## Provider Development

This provider is maintained by the ProData team. For issues, feature requests, or contributions, visit our GitHub repository.
