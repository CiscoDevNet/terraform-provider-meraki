---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_stp Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch STP configuration.
---

# meraki_switch_stp (Resource)

This resource can manage the `Switch STP` configuration.

## Example Usage

```terraform
resource "meraki_switch_stp" "example" {
  network_id   = "L_123456"
  rstp_enabled = true
  stp_bridge_priority = [
    {
      stp_priority    = 4096
      stacks          = ["789102"]
      switch_profiles = ["1098"]
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

- `rstp_enabled` (Boolean) The spanning tree protocol status in network
- `stp_bridge_priority` (Attributes List) STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings. (see [below for nested schema](#nestedatt--stp_bridge_priority))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--stp_bridge_priority"></a>
### Nested Schema for `stp_bridge_priority`

Required:

- `stp_priority` (Number) STP priority for switch, stacks, or switch templates

Optional:

- `stacks` (List of String) List of stack IDs
- `switch_profiles` (List of String) List of switch template IDs
- `switches` (List of String) List of switch serial numbers

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_switch_stp.example "<network_id>"
```