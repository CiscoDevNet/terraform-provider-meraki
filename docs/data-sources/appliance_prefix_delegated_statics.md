---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_prefix_delegated_statics Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Prefix Delegated Static configuration in bulk.
---

# meraki_appliance_prefix_delegated_statics (Data Source)

This data source can read the `Appliance Prefix Delegated Static` configuration in bulk.

## Example Usage

```terraform
data "meraki_appliance_prefix_delegated_statics" "example" {
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

- `description` (String) A name or description for the prefix
- `id` (String) The id of the object
- `origin_interfaces` (List of String) Interfaces associated with the prefix
- `origin_type` (String) Type of the origin
- `prefix` (String) A static IPv6 prefix
