---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_air_marshal_rule Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless Air Marshal Rule configuration.
---

# meraki_wireless_air_marshal_rule (Resource)

This resource can manage the `Wireless Air Marshal Rule` configuration.

## Example Usage

```terraform
resource "meraki_wireless_air_marshal_rule" "example" {
  network_id   = "L_123456"
  type         = "allow"
  match_string = "00:11:22:33:44:55"
  match_type   = "bssid"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `type` (String) Indicates if this rule will allow, block, or alert.
  - Choices: `alert`, `allow`, `block`

### Optional

- `match_string` (String) The string used to match.
- `match_type` (String) The type of match.
  - Choices: `bssid`, `contains`, `exact`, `wildcard`

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_wireless_air_marshal_rule.example "<network_id>,<id>"
```
