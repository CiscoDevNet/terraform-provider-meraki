---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_device_type_group_policies Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless SSID Device Type Group Policies configuration.
---

# meraki_wireless_ssid_device_type_group_policies (Data Source)

This data source can read the `Wireless SSID Device Type Group Policies` configuration.

## Example Usage

```terraform
data "meraki_wireless_ssid_device_type_group_policies" "example" {
  network_id = "L_123456"
  number     = "0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Read-Only

- `device_type_policies` (Attributes List) List of device type policies. (see [below for nested schema](#nestedatt--device_type_policies))
- `enabled` (Boolean) If true, the SSID device type group policies are enabled.
- `id` (String) The id of the object

<a id="nestedatt--device_type_policies"></a>
### Nested Schema for `device_type_policies`

Read-Only:

- `device_policy` (String) The device policy. Can be one of `Allowed`, `Blocked` or `Group policy`
- `device_type` (String) The device type. Can be one of `Android`, `BlackBerry`, `Chrome OS`, `iPad`, `iPhone`, `iPod`, `Mac OS X`, `Windows`, `Windows Phone`, `B&N Nook` or `Other OS`
- `group_policy_id` (Number) ID of the group policy object.