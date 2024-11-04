---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_ports Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  This data source can read the Appliance Ports configuration.
---

# meraki_appliance_ports (Data Source)

This data source can read the `Appliance Ports` configuration.

## Example Usage

```terraform
data "meraki_appliance_ports" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `items` (Attributes List) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `access_policy` (String) The name of the policy. Only applicable to Access ports.
- `allowed_vlans` (String) Comma-delimited list of the VLAN ID`s allowed on the port, or `all` to permit all VLAN`s on the port.
- `drop_untagged_traffic` (Boolean) Whether the trunk port can drop all untagged traffic.
- `enabled` (Boolean) The status of the port
- `id` (String) The id of the object
- `number` (Number) Number of the port
- `type` (String) The type of the port: `access` or `trunk`.
- `vlan` (Number) Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.