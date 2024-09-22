---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_settings Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless settings configuration.
---

# meraki_wireless_settings (Data Source)

This data source can read the `Wireless settings` configuration.

## Example Usage

```terraform
data "meraki_wireless_settings" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object
- `ipv6_bridge_enabled` (Boolean) Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
- `led_lights_on` (Boolean) Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
- `location_analytics_enabled` (Boolean) Toggle for enabling or disabling location analytics for your network
- `meshing_enabled` (Boolean) Toggle for enabling or disabling meshing in a network
- `named_vlans_pool_dhcp_monitoring_duration` (Number) The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
- `named_vlans_pool_dhcp_monitoring_enabled` (Boolean) Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP.
- `upgrade_strategy` (String) The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.