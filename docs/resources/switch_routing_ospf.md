---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_routing_ospf Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch Routing OSPF configuration.
---

# meraki_switch_routing_ospf (Resource)

This resource can manage the `Switch Routing OSPF` configuration.

## Example Usage

```terraform
resource "meraki_switch_routing_ospf" "example" {
  network_id                        = "L_123456"
  dead_timer_in_seconds             = 40
  enabled                           = true
  hello_timer_in_seconds            = 10
  md5_authentication_enabled        = true
  md5_authentication_key_id         = 1
  md5_authentication_key_passphrase = "abc1234"
  v3_dead_timer_in_seconds          = 40
  v3_enabled                        = true
  v3_hello_timer_in_seconds         = 10
  v3_areas = [
    {
      area_id   = "0"
      area_name = "V3 Backbone"
      area_type = "normal"
    }
  ]
  areas = [
    {
      area_id   = "0"
      area_name = "Backbone"
      area_type = "normal"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `areas` (Attributes List) OSPF areas (see [below for nested schema](#nestedatt--areas))
- `dead_timer_in_seconds` (Number) Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
- `enabled` (Boolean) Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.
- `hello_timer_in_seconds` (Number) Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
- `md5_authentication_enabled` (Boolean) Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.
- `md5_authentication_key_id` (Number) MD5 authentication key index. Key index must be between 1 to 255
- `md5_authentication_key_passphrase` (String) MD5 authentication passphrase
- `v3_areas` (Attributes List) OSPF v3 areas (see [below for nested schema](#nestedatt--v3_areas))
- `v3_dead_timer_in_seconds` (Number) Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
- `v3_enabled` (Boolean) Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
- `v3_hello_timer_in_seconds` (Number) Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--areas"></a>
### Nested Schema for `areas`

Required:

- `area_id` (String) OSPF area ID
- `area_name` (String) Name of the OSPF area
- `area_type` (String) Area types in OSPF. Must be one of: ['normal', 'stub', 'nssa']
  - Choices: `normal`, `nssa`, `stub`


<a id="nestedatt--v3_areas"></a>
### Nested Schema for `v3_areas`

Required:

- `area_id` (String) OSPF area ID
- `area_name` (String) Name of the OSPF area
- `area_type` (String) Area types in OSPF. Must be one of: ['normal', 'stub', 'nssa']
  - Choices: `normal`, `nssa`, `stub`

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_switch_routing_ospf.example "<network_id>"
```