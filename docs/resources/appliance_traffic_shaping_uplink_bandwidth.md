---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_traffic_shaping_uplink_bandwidth Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance Traffic Shaping Uplink Bandwidth configuration.
---

# meraki_appliance_traffic_shaping_uplink_bandwidth (Resource)

This resource can manage the `Appliance Traffic Shaping Uplink Bandwidth` configuration.

## Example Usage

```terraform
resource "meraki_appliance_traffic_shaping_uplink_bandwidth" "example" {
  network_id          = "L_123456"
  cellular_limit_down = 100000
  cellular_limit_up   = 100000
  wan1_limit_down     = 100000
  wan1_limit_up       = 100000
  wan2_limit_down     = 100000
  wan2_limit_up       = 100000
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `cellular_limit_down` (Number) The maximum download limit (integer, in Kbps). null indicates no limit
- `cellular_limit_up` (Number) The maximum upload limit (integer, in Kbps). null indicates no limit
- `wan1_limit_down` (Number) The maximum download limit (integer, in Kbps). null indicates no limit
- `wan1_limit_up` (Number) The maximum upload limit (integer, in Kbps). null indicates no limit
- `wan2_limit_down` (Number) The maximum download limit (integer, in Kbps). null indicates no limit
- `wan2_limit_up` (Number) The maximum upload limit (integer, in Kbps). null indicates no limit

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_traffic_shaping_uplink_bandwidth.example "<network_id>"
```
