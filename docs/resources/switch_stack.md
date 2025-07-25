---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_stack Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch Stack configuration.
---

# meraki_switch_stack (Resource)

This resource can manage the `Switch Stack` configuration.

## Example Usage

```terraform
resource "meraki_switch_stack" "example" {
  network_id = "L_123456"
  name       = "A cool stack"
  serials    = ["QBZY-XWVU-TSRQ"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the new stack
- `network_id` (String) Network ID
- `serials` (Set of String) An array of switch serials to be added into the new stack

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_switch_stack.example "<network_id>,<id>"
```
