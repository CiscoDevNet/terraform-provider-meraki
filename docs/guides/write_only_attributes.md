---
subcategory: "Guides"
page_title: "Write-Only Attributes"
description: |-
    Write-Only Attributes
---

# Write-Only Attributes

Starting with Terraform 1.11, the Terraform Plugin Framework supports [write-only attributes](https://developer.hashicorp.com/terraform/language/resources/ephemeral/write-only) — attributes whose values are sent to the API but never persisted in state. This provider uses write-only attributes for all sensitive string attributes (passwords, secrets, passphrases, PSKs, shared secrets, and similar credentials).

## How It Works

Every sensitive string attribute now has two optional sibling attributes with `_wo` and `_wo_version` suffixes. For example, the `meraki_network_snmp` resource has:

| Attribute | Purpose |
|---|---|
| `community_string` | Original attribute — value is stored in state (marked as sensitive) |
| `community_string_wo` | Write-only variant — value is sent to the API but **never stored in state** |
| `community_string_wo_version` | Version tracker — increment this value to re-send `community_string_wo` on the next apply |

When `community_string_wo` is set, it takes precedence over `community_string`. Both attributes map to the same API field.

## Usage

### Using Write-Only Attributes (Recommended)

```hcl
resource "meraki_network_snmp" "example" {
  network_id                  = meraki_network.example.id
  access                      = "community"
  community_string_wo         = var.snmp_community
  community_string_wo_version = 1
}
```

The value of `community_string_wo` is sent to the Meraki API during create and update, but is not stored in the Terraform state file. To force the value to be re-sent (for example, after rotating a credential), increment `community_string_wo_version`.

### Using the Original Sensitive Attribute

The original sensitive attribute continues to work as before:

```hcl
resource "meraki_network_snmp" "example" {
  network_id       = meraki_network.example.id
  access           = "community"
  community_string = var.snmp_community
}
```

The value is stored in state (marked as sensitive so it is redacted in plan output and logs), but it is still present in the state file on disk.

### Nested Attributes

Write-only siblings are also available inside nested blocks. For example, RADIUS server secrets:

```hcl
resource "meraki_wireless_ssid" "example" {
  network_id = meraki_network.example.id
  number     = "0"
  name       = "Corp Wi-Fi"

  radius_servers = [{
    host              = "10.0.0.1"
    port              = 1812
    secret_wo         = var.radius_secret
    secret_wo_version = 1
  }]
}
```

## Terraform Version Requirements

- **Terraform 1.11+**: Full support for write-only attributes. Both `_wo` and the original sensitive attribute work.
- **Terraform 1.10 and earlier**: Write-only attributes are not supported. Use the original sensitive attribute instead. The `_wo` and `_wo_version` attributes are silently ignored.
