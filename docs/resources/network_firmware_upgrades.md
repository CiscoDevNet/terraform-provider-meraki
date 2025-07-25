---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_firmware_upgrades Resource - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This resource can manage the Network Firmware Upgrades configuration.
---

# meraki_network_firmware_upgrades (Resource)

This resource can manage the `Network Firmware Upgrades` configuration.

## Example Usage

```terraform
resource "meraki_network_firmware_upgrades" "example" {
  network_id                                                    = "L_123456"
  timezone                                                      = "America/Los_Angeles"
  products_appliance_participate_in_next_beta_release           = false
  products_appliance_next_upgrade_time                          = "2019-03-17T17:22:52Z"
  products_appliance_next_upgrade_to_version_id                 = "1001"
  products_camera_participate_in_next_beta_release              = false
  products_camera_next_upgrade_time                             = "2019-03-17T17:22:52Z"
  products_camera_next_upgrade_to_version_id                    = "1003"
  products_cellular_gateway_participate_in_next_beta_release    = false
  products_cellular_gateway_next_upgrade_time                   = "2019-03-17T17:22:52Z"
  products_cellular_gateway_next_upgrade_to_version_id          = "1004"
  products_secure_connect_participate_in_next_beta_release      = false
  products_secure_connect_next_upgrade_time                     = "2019-03-17T17:22:52Z"
  products_secure_connect_next_upgrade_to_version_id            = "1007"
  products_sensor_participate_in_next_beta_release              = false
  products_sensor_next_upgrade_time                             = "2019-03-17T17:22:52Z"
  products_sensor_next_upgrade_to_version_id                    = "1005"
  products_switch_participate_in_next_beta_release              = false
  products_switch_next_upgrade_time                             = "2019-03-17T17:22:52Z"
  products_switch_next_upgrade_to_version_id                    = "1002"
  products_switch_catalyst_participate_in_next_beta_release     = false
  products_switch_catalyst_next_upgrade_time                    = "2019-03-17T17:22:52Z"
  products_switch_catalyst_next_upgrade_to_version_id           = "1234"
  products_wireless_participate_in_next_beta_release            = false
  products_wireless_next_upgrade_time                           = "2019-03-17T17:22:52Z"
  products_wireless_next_upgrade_to_version_id                  = "1000"
  products_wireless_controller_participate_in_next_beta_release = false
  products_wireless_controller_next_upgrade_time                = "2019-03-17T17:22:52Z"
  products_wireless_controller_next_upgrade_to_version_id       = "1006"
  upgrade_window_day_of_week                                    = "sun"
  upgrade_window_hour_of_day                                    = "4:00"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `products_appliance_next_upgrade_time` (String) The time of the last successful upgrade
- `products_appliance_next_upgrade_to_version_id` (String) The version ID
- `products_appliance_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_camera_next_upgrade_time` (String) The time of the last successful upgrade
- `products_camera_next_upgrade_to_version_id` (String) The version ID
- `products_camera_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_cellular_gateway_next_upgrade_time` (String) The time of the last successful upgrade
- `products_cellular_gateway_next_upgrade_to_version_id` (String) The version ID
- `products_cellular_gateway_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_secure_connect_next_upgrade_time` (String) The time of the last successful upgrade
- `products_secure_connect_next_upgrade_to_version_id` (String) The version ID
- `products_secure_connect_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_sensor_next_upgrade_time` (String) The time of the last successful upgrade
- `products_sensor_next_upgrade_to_version_id` (String) The version ID
- `products_sensor_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_switch_catalyst_next_upgrade_time` (String) The time of the last successful upgrade
- `products_switch_catalyst_next_upgrade_to_version_id` (String) The version ID
- `products_switch_catalyst_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_switch_next_upgrade_time` (String) The time of the last successful upgrade
- `products_switch_next_upgrade_to_version_id` (String) The version ID
- `products_switch_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_wireless_controller_next_upgrade_time` (String) The time of the last successful upgrade
- `products_wireless_controller_next_upgrade_to_version_id` (String) The version ID
- `products_wireless_controller_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `products_wireless_next_upgrade_time` (String) The time of the last successful upgrade
- `products_wireless_next_upgrade_to_version_id` (String) The version ID
- `products_wireless_participate_in_next_beta_release` (Boolean) Whether or not the network wants beta firmware
- `timezone` (String) The timezone for the network
- `upgrade_window_day_of_week` (String) Day of the week
  - Choices: `fri`, `friday`, `mon`, `monday`, `sat`, `saturday`, `sun`, `sunday`, `thu`, `thursday`, `tue`, `tuesday`, `wed`, `wednesday`
- `upgrade_window_hour_of_day` (String) Hour of the day
  - Choices: `0:00`, `10:00`, `11:00`, `12:00`, `13:00`, `14:00`, `15:00`, `16:00`, `17:00`, `18:00`, `19:00`, `1:00`, `20:00`, `21:00`, `22:00`, `23:00`, `2:00`, `3:00`, `4:00`, `5:00`, `6:00`, `7:00`, `8:00`, `9:00`

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_network_firmware_upgrades.example "<network_id>"
```
