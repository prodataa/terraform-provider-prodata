---
page_title: "ProData Provider"
description: |-
  The ProData provider enables seamless infrastructure management for ProData Cloud resources through Terraform.
---

# ProData Provider

The ProData provider allows you to manage your ProData Cloud infrastructure using Terraform's declarative configuration language. With this provider, you can automate the provisioning, configuration, and lifecycle management of ProData resources.

## Requirements

| Requirement                                                      | Version |
| ---------------------------------------------------------------- | ------- |
| [Terraform](https://developer.hashicorp.com/terraform/downloads) | >= 1.0  |
| ProData Provider                                                 | >= 1.0  |

## Getting Started

Before you begin, ensure you have:

1. A [ProData Cloud account](https://my.pro-data.tech)
2. API credentials (API Key ID and Secret Key)
3. Your Project ID
4. Selected region (`UZ-5`, `UZ-3`, or `KZ-1`)

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

## Argument Reference

The following arguments are supported:

| Argument         | Type   | Required | Description                                                                                          |
| ---------------- | ------ | -------- | ---------------------------------------------------------------------------------------------------- |
| `api_base_url`   | String | Yes      | The base URL for the ProData API. See [Regional Availability](#regional-availability) for endpoints. |
| `api_key_id`     | String | Yes      | Your API Key ID for authentication.                                                                  |
| `api_secret_key` | String | Yes      | Your API Secret Key for authentication. **Sensitive** - will not appear in logs.                     |
| `region`         | String | Yes      | The region where resources will be created. Valid values: `UZ-5`, `UZ-3`, `KZ-1`.                    |
| `project_id`     | String | Yes      | Your ProData project ID.                                                                             |

### Environment Variables

All arguments can be set via environment variables. This is the recommended approach for credentials.

| Argument         | Environment Variable     |
| ---------------- | ------------------------ |
| `api_base_url`   | `PRODATA_API_BASE_URL`   |
| `api_key_id`     | `PRODATA_API_KEY_ID`     |
| `api_secret_key` | `PRODATA_API_SECRET_KEY` |
| `region`         | `PRODATA_REGION`         |
| `project_id`     | `PRODATA_PROJECT_ID`     |

> **Note:** Values set in the provider block override environment variables.

## Obtaining API Credentials

To generate API credentials:

1. Log in to your [ProData Cloud console](https://my.pro-data.tech)
2. Navigate to **Account** → **Access Keys**
3. Click **Generate Key**
4. Copy your API Key ID and Secret Key

> **Important:** The Secret Key is only displayed once. Store it securely immediately after creation.

## Regional Availability

ProData Cloud is available in the following regions:

| Region Code | Location   | API Endpoint                 |
| ----------- | ---------- | ---------------------------- |
| `UZ-5`      | Uzbekistan | `https://my.pro-data.tech`   |
| `UZ-3`      | Uzbekistan | `https://my.pro-data.tech`   |
| `KZ-1`      | Kazakhstan | `https://kz-1.pro-data.tech` |

## Troubleshooting

### Authentication Errors

If you receive authentication errors:

1. Verify your API Key ID and Secret Key are correct
2. Check that your API key has not expired or been revoked
3. Ensure you're using the correct `api_base_url` for your region

### Region Errors

If resources fail to create due to region issues:

1. Confirm the region code matches your project's location
2. Verify the `api_base_url` corresponds to your selected region

## Support

| Channel   | Link                                                       |
| --------- | ---------------------------------------------------------- |
| Help Desk | [helpdesk.pro-data.tech](https://helpdesk.pro-data.tech)   |
| Telegram  | [@PRO_DATA_Support_Bot](https://t.me/PRO_DATA_Support_Bot) |
