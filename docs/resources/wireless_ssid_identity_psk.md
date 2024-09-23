---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_identity_psk Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless SSID Identity PSK configuration.
---

# meraki_wireless_ssid_identity_psk (Resource)

This resource can manage the `Wireless SSID Identity PSK` configuration.

## Example Usage

```terraform
resource "meraki_wireless_ssid_identity_psk" "example" {
  network_id      = "L_123456"
  number          = "0"
  expires_at      = "2018-02-11T00:00:00.090209Z"
  group_policy_id = "101"
  name            = "Sample Identity PSK"
  passphrase      = "Cisco123"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_policy_id` (String) The group policy to be applied to clients
- `name` (String) The name of the Identity PSK
- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `expires_at` (String) Timestamp for when the Identity PSK expires. Will not expire if left blank.
- `passphrase` (String) The passphrase for client authentication. If left blank, one will be auto-generated.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_wireless_ssid_identity_psk.example "<network_id>,<number>,<id>"
```