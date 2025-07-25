---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_rf_profiles Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance RF Profile configuration in bulk.
  This bulk resource uses the following attributes to uniquely identify each object. Changing any of these attributes will cause the object to be deleted and recreated.
  name
---

# meraki_appliance_rf_profiles (Resource)

This resource can manage the `Appliance RF Profile` configuration in bulk.

This bulk resource uses the following attributes to uniquely identify each object. Changing any of these attributes will cause the object to be deleted and recreated.

- `name`

## Example Usage

```terraform
resource "meraki_appliance_rf_profiles" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    name                                      = "MX RF Profile"
    five_ghz_settings_ax_enabled              = true
    five_ghz_settings_min_bitrate             = 48
    per_ssid_settings_1_band_operation_mode   = "dual"
    per_ssid_settings_1_band_steering_enabled = true
    per_ssid_settings_2_band_operation_mode   = "dual"
    per_ssid_settings_2_band_steering_enabled = true
    per_ssid_settings_3_band_operation_mode   = "dual"
    per_ssid_settings_3_band_steering_enabled = true
    per_ssid_settings_4_band_operation_mode   = "dual"
    per_ssid_settings_4_band_steering_enabled = true
    two_four_ghz_settings_ax_enabled          = true
    two_four_ghz_settings_min_bitrate         = 12
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))
- `network_id` (String) Network ID
- `organization_id` (String) The organization ID

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Required:

- `name` (String) The name of the new profile. Must be unique. This param is required on creation.

Optional:

- `five_ghz_settings_ax_enabled` (Boolean) Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
- `five_ghz_settings_min_bitrate` (Number) Sets min bitrate (Mbps) of 5Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.
- `per_ssid_settings_1_band_operation_mode` (String) Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.
  - Choices: `2.4ghz`, `5ghz`, `6ghz`, `dual`, `multi`
- `per_ssid_settings_1_band_steering_enabled` (Boolean) Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
- `per_ssid_settings_2_band_operation_mode` (String) Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.
  - Choices: `2.4ghz`, `5ghz`, `6ghz`, `dual`, `multi`
- `per_ssid_settings_2_band_steering_enabled` (Boolean) Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
- `per_ssid_settings_3_band_operation_mode` (String) Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.
  - Choices: `2.4ghz`, `5ghz`, `6ghz`, `dual`, `multi`
- `per_ssid_settings_3_band_steering_enabled` (Boolean) Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
- `per_ssid_settings_4_band_operation_mode` (String) Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.
  - Choices: `2.4ghz`, `5ghz`, `6ghz`, `dual`, `multi`
- `per_ssid_settings_4_band_steering_enabled` (Boolean) Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
- `two_four_ghz_settings_ax_enabled` (Boolean) Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
- `two_four_ghz_settings_min_bitrate` (Number) Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 11.

Read-Only:

- `id` (String) The id of the item

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_rf_profiles.example "<organization_id>,<network_id>"
```
