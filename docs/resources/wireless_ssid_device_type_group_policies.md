---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_device_type_group_policies Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless SSID Device Type Group Policies configuration.
---

# meraki_wireless_ssid_device_type_group_policies (Resource)

This resource can manage the `Wireless SSID Device Type Group Policies` configuration.

## Example Usage

```terraform
resource "meraki_wireless_ssid_device_type_group_policies" "example" {
  network_id = "L_123456"
  number     = "0"
  enabled    = true
  device_type_policies = [
    {
      device_policy = "Allowed"
      device_type   = "Android"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `device_type_policies` (Attributes List) List of device type policies. (see [below for nested schema](#nestedatt--device_type_policies))
- `enabled` (Boolean) If true, the SSID device type group policies are enabled.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--device_type_policies"></a>
### Nested Schema for `device_type_policies`

Required:

- `device_policy` (String) The device policy. Can be one of `Allowed`, `Blocked` or `Group policy`
  - Choices: `Allowed`, `Blocked`, `Group policy`
- `device_type` (String) The device type. Can be one of `Android`, `BlackBerry`, `Chrome OS`, `iPad`, `iPhone`, `iPod`, `Mac OS X`, `Windows`, `Windows Phone`, `B&N Nook` or `Other OS`
  - Choices: `Android`, `B&N Nook`, `BlackBerry`, `Chrome OS`, `Mac OS X`, `Other OS`, `Windows`, `Windows Phone`, `iPad`, `iPhone`, `iPod`

Optional:

- `group_policy_id` (Number) ID of the group policy object.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_wireless_ssid_device_type_group_policies.example "<network_id>,<number>"
```
