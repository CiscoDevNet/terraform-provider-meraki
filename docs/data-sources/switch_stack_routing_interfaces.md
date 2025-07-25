---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_stack_routing_interfaces Data Source - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This data source can read the Switch Stack Routing Interface configuration in bulk.
---

# meraki_switch_stack_routing_interfaces (Data Source)

This data source can read the `Switch Stack Routing Interface` configuration in bulk.

## Example Usage

```terraform
data "meraki_switch_stack_routing_interfaces" "example" {
  network_id      = "L_123456"
  switch_stack_id = "1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `switch_stack_id` (String) Switch stack ID

### Read-Only

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `default_gateway` (String) The next hop for any traffic that isn`t going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
- `id` (String) The id of the object
- `interface_ip` (String) The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch`s management IP.
- `ipv6_address` (String) The IPv6 address of the interface. Required if assignmentMode is `static`. Must not be included if assignmentMode is `eui-64`.
- `ipv6_assignment_mode` (String) The IPv6 assignment mode for the interface. Can be either `eui-64` or `static`.
- `ipv6_gateway` (String) The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.
- `ipv6_prefix` (String) The IPv6 prefix of the interface. Required if IPv6 object is included.
- `multicast_routing` (String) Enable multicast support if, multicast routing between VLANs is required. Options are, `disabled`, `enabled` or `IGMP snooping querier`. Default is `disabled`.
- `name` (String) A friendly name or description for the interface or VLAN.
- `ospf_settings_area` (String) The OSPF area to which this interface should belong. Can be either `ospfDisabled` or the identifier of an existing OSPF area. Defaults to `ospfDisabled`.
- `ospf_settings_cost` (Number) The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
- `ospf_settings_is_passive_enabled` (Boolean) When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
- `subnet` (String) The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
- `vlan_id` (Number) The VLAN this routed interface is on. VLAN must be between 1 and 4094.
