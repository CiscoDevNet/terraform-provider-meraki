---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_stacks Data Source - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This data source can read the Switch Stack configuration in bulk.
---

# meraki_switch_stacks (Data Source)

This data source can read the `Switch Stack` configuration in bulk.

## Example Usage

```terraform
data "meraki_switch_stacks" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String) The id of the object
- `name` (String) The name of the new stack
- `serials` (Set of String) An array of switch serials to be added into the new stack
