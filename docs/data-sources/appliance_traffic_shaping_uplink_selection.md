---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_traffic_shaping_uplink_selection Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Traffic Shaping Uplink Selection configuration.
---

# meraki_appliance_traffic_shaping_uplink_selection (Data Source)

This data source can read the `Appliance Traffic Shaping Uplink Selection` configuration.

## Example Usage

```terraform
data "meraki_appliance_traffic_shaping_uplink_selection" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `active_active_auto_vpn_enabled` (Boolean) Toggle for enabling or disabling active-active AutoVPN
- `default_uplink` (String) The default uplink. Must be a WAN interface `wanX`
- `failover_and_failback_immediate_enabled` (Boolean) Toggle for enabling or disabling immediate WAN failover and failback
- `id` (String) The id of the object
- `load_balancing_enabled` (Boolean) Toggle for enabling or disabling load balancing
- `vpn_traffic_uplink_preferences` (Attributes List) Array of uplink preference rules for VPN traffic (see [below for nested schema](#nestedatt--vpn_traffic_uplink_preferences))
- `wan_traffic_uplink_preferences` (Attributes List) Array of uplink preference rules for WAN traffic. Note: these preferences are shared (merged) with meraki_appliance_sdwan_internet_policies resource. It is recommended to only use one resource for these preferences, not both at the same time: Deleting this resource clears preferences created in meraki_appliance_sdwan_internet_policies, which isn't detected as a change by the provider. The same happens the other way around, but the change is detected in uplink_selection on a subsequent apply. (see [below for nested schema](#nestedatt--wan_traffic_uplink_preferences))

<a id="nestedatt--vpn_traffic_uplink_preferences"></a>
### Nested Schema for `vpn_traffic_uplink_preferences`

Read-Only:

- `builtin_performance_class_name` (String) Name of builtin performance class, must be present when performanceClass type is `builtin`, and value must be one of: `VoIP`
- `custom_performance_class_id` (String) ID of created custom performance class, must be present when performanceClass type is `custom`
- `fail_over_criterion` (String) Fail over criterion for this uplink preference rule. Must be one of: `poorPerformance` or `uplinkDown`
- `performance_class_type` (String) Type of this performance class. Must be one of: `builtin` or `custom`
- `preferred_uplink` (String) Preferred uplink for uplink preference rule. Must be one of: `wan1`, `wan2`, `bestForVoIP`, `loadBalancing` or `defaultUplink`, or any other valid uplink(`wanX`) if it applies to the network
- `traffic_filters` (Attributes List) Array of traffic filters for this uplink preference rule (see [below for nested schema](#nestedatt--vpn_traffic_uplink_preferences--traffic_filters))

<a id="nestedatt--vpn_traffic_uplink_preferences--traffic_filters"></a>
### Nested Schema for `vpn_traffic_uplink_preferences.traffic_filters`

Read-Only:

- `destination_cidr` (String) CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')
- `destination_fqdn` (String) FQDN format address. Currently only availabe in `destination` of `vpnTrafficUplinkPreference` object. E.g.: `www.google.com`
- `destination_host` (Number) Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.
- `destination_network` (String) Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.
- `destination_port` (String) E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'
- `destination_vlan` (Number) VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
- `id` (String) ID of this applicationCategory or application type traffic filter. E.g.: 'meraki:layer7/category/1', 'meraki:layer7/application/4'
- `protocol` (String) Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp`, `icmp6` or `any`
- `source_cidr` (String) CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')
- `source_host` (Number) Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.
- `source_network` (String) Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.
- `source_port` (String) E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'
- `source_vlan` (Number) VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
- `type` (String) Type of this traffic filter. Must be one of: `applicationCategory`, `application` or `custom`



<a id="nestedatt--wan_traffic_uplink_preferences"></a>
### Nested Schema for `wan_traffic_uplink_preferences`

Read-Only:

- `preferred_uplink` (String) Preferred uplink for uplink preference rule. Must be one of: `wan1` or `wan2`, or any other valid uplink(`wanX`) if it applies to the network
- `traffic_filters` (Attributes List) Array of traffic filters for this uplink preference rule (see [below for nested schema](#nestedatt--wan_traffic_uplink_preferences--traffic_filters))

<a id="nestedatt--wan_traffic_uplink_preferences--traffic_filters"></a>
### Nested Schema for `wan_traffic_uplink_preferences.traffic_filters`

Read-Only:

- `destination_cidr` (String) CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')
- `destination_port` (String) E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'
- `protocol` (String) Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp6` or `any`
- `source_cidr` (String) CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')
- `source_host` (Number) Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.
- `source_port` (String) E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'
- `source_vlan` (Number) VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
- `type` (String) Type of this traffic filter. Must be one of: `custom`
