---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_dscp_to_cos_mappings Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch DSCP to CoS mappings configuration.
---

# meraki_switch_dscp_to_cos_mappings (Resource)

This resource can manage the `Switch DSCP to CoS mappings` configuration.

## Example Usage

```terraform
resource "meraki_switch_dscp_to_cos_mappings" "example" {
  network_id = "L_123456"
  mappings = [
    {
      cos   = 1
      dscp  = 1
      title = "Video"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mappings` (Attributes List) An array of DSCP to CoS mappings. An empty array will reset the mappings to default. (see [below for nested schema](#nestedatt--mappings))
- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--mappings"></a>
### Nested Schema for `mappings`

Required:

- `cos` (Number) The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
- `dscp` (Number) The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.

Optional:

- `title` (String) Label for the mapping (optional).

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_switch_dscp_to_cos_mappings.example "<network_id>"
```
