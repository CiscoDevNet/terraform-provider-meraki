---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_mtu Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch MTU configuration.
---

# meraki_switch_mtu (Resource)

This resource can manage the `Switch MTU` configuration.

## Example Usage

```terraform
resource "meraki_switch_mtu" "example" {
  network_id       = "L_123456"
  default_mtu_size = 9578
  overrides = [
    {
      mtu_size        = 1500
      switch_profiles = ["1284392014819"]
      switches        = ["Q234-ABCD-0001"]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `default_mtu_size` (Number) MTU size for the entire network. Default value is 9578.
- `overrides` (Attributes List) Override MTU size for individual switches or switch templates. An empty array will clear overrides. (see [below for nested schema](#nestedatt--overrides))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--overrides"></a>
### Nested Schema for `overrides`

Required:

- `mtu_size` (Number) MTU size for the switches or switch templates.

Optional:

- `switch_profiles` (List of String) List of switch template IDs. Applicable only for template network.
- `switches` (List of String) List of switch serials. Applicable only for switch network.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_switch_mtu.example "<network_id>"
```