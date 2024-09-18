---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_device Data Source - terraform-provider-meraki"
subcategory: "Devices"
description: |-
  This data source can read the Device configuration.
---

# meraki_device (Data Source)

This data source can read the `Device` configuration.

## Example Usage

```terraform
data "meraki_device" "example" {
  id     = "12345678"
  serial = "1234-ABCD-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) Switch serial

### Read-Only

- `address` (String) The address of a device
- `floor_plan_id` (String) The floor plan to associate to this device. null disassociates the device from the floorplan.
- `id` (String) The id of the object
- `lat` (Number) The latitude of a device
- `lng` (Number) The longitude of a device
- `move_map_marker` (Boolean) Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
- `name` (String) The name of a device
- `notes` (String) The notes for the device. String. Limited to 255 characters.
- `switch_profile_id` (String) The ID of a switch template to bind to the device (for available switch templates, see the `Switch Templates` endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
- `tags` (List of String) The list of tags of a device