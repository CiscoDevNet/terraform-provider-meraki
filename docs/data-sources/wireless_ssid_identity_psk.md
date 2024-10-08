---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_identity_psk Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless SSID Identity PSK configuration.
---

# meraki_wireless_ssid_identity_psk (Data Source)

This data source can read the `Wireless SSID Identity PSK` configuration.

## Example Usage

```terraform
data "meraki_wireless_ssid_identity_psk" "example" {
  id         = "12345678"
  network_id = "L_123456"
  number     = "0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the Identity PSK

### Read-Only

- `expires_at` (String) Timestamp for when the Identity PSK expires. Will not expire if left blank.
- `group_policy_id` (String) The group policy to be applied to clients
- `passphrase` (String) The passphrase for client authentication. If left blank, one will be auto-generated.
