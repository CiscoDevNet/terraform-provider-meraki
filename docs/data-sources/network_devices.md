---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_devices Data Source - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This data source can read the Network Devices configuration in bulk.
---

# meraki_network_devices (Data Source)

This data source can read the `Network Devices` configuration in bulk.

## Example Usage

```terraform
data "meraki_network_devices" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `address` (String) Physical address of the device
- `beacon_id_params_major` (Number) The major number to be used in the beacon identifier
- `beacon_id_params_minor` (Number) The minor number to be used in the beacon identifier
- `beacon_id_params_uuid` (String) The UUID to be used in the beacon identifier
- `details` (Attributes List) Additional device information (see [below for nested schema](#nestedatt--items--details))
- `firmware` (String) Firmware version of the device
- `floor_plan_id` (String) The floor plan to associate to this device. null disassociates the device from the floorplan.
- `id` (String) The id of the object
- `lan_ip` (String) LAN IP address of the device
- `lat` (Number) Latitude of the device
- `lng` (Number) Longitude of the device
- `mac` (String) MAC address of the device
- `model` (String) Model of the device
- `name` (String) Name of the device
- `network_id` (String) ID of the network the device belongs to
- `notes` (String) Notes for the device, limited to 255 characters
- `serial` (String) Serial number of the device
- `tags` (List of String) List of tags assigned to the device

<a id="nestedatt--items--details"></a>
### Nested Schema for `items.details`

Read-Only:

- `name` (String) Additional property name
- `value` (String) Additional property value
