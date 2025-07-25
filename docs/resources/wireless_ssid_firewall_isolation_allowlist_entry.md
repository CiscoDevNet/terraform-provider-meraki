---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_firewall_isolation_allowlist_entry Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless SSID Firewall Isolation Allowlist Entry configuration.
---

# meraki_wireless_ssid_firewall_isolation_allowlist_entry (Resource)

This resource can manage the `Wireless SSID Firewall Isolation Allowlist Entry` configuration.

## Example Usage

```terraform
resource "meraki_wireless_ssid_firewall_isolation_allowlist_entry" "example" {
  organization_id = "123456"
  description     = "Example mac address"
  client_mac      = "A1:B2:C3:D4:E5:F6"
  network_id      = "N_123"
  ssid_number     = 2
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_mac` (String) L2 Isolation mac address
- `network_id` (String) The ID of network
- `organization_id` (String) Organization ID
- `ssid_number` (Number) The number of SSID

### Optional

- `description` (String) The description of mac address

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_wireless_ssid_firewall_isolation_allowlist_entry.example "<organization_id>,<id>"
```
