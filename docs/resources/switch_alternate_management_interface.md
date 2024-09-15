---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_alternate_management_interface Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch Alternate Management Interface configuration.
---

# meraki_switch_alternate_management_interface (Resource)

This resource can manage the `Switch Alternate Management Interface` configuration.

## Example Usage

```terraform
resource "meraki_switch_alternate_management_interface" "example" {
  network_id = "L_123456"
  enabled    = true
  vlan_id    = 100
  protocols  = ["radius"]
  switches = [
    {
      alternate_management_ip = "1.2.3.4"
      gateway                 = "1.2.3.5"
      serial                  = "Q234-ABCD-5678"
      subnet_mask             = "255.255.255.0"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `enabled` (Boolean) Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
- `protocols` (List of String) Can be one or more of the following values: `radius`, `snmp` or `syslog`
- `switches` (Attributes List) Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put `switches` in the body when updating template networks. Also, an empty `switches` array will remove all previous assignments (see [below for nested schema](#nestedatt--switches))
- `vlan_id` (Number) Alternate management VLAN, must be between 1 and 4094

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--switches"></a>
### Nested Schema for `switches`

Required:

- `alternate_management_ip` (String) Switch alternative management IP. To remove a prior IP setting, provide an empty string
- `serial` (String) Switch serial number

Optional:

- `gateway` (String) Switch gateway must be in IP format. Only and must be specified for Polaris switches
- `subnet_mask` (String) Switch subnet mask must be in IP format. Only and must be specified for Polaris switches

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_switch_alternate_management_interface.example "<network_id>"
```