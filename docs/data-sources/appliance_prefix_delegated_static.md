---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_prefix_delegated_static Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Prefix Delegated Static configuration.
---

# meraki_appliance_prefix_delegated_static (Data Source)

This data source can read the `Appliance Prefix Delegated Static` configuration.

## Example Usage

```terraform
data "meraki_appliance_prefix_delegated_static" "example" {
  id         = "12345678"
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object
- `network_id` (String) Network ID

### Read-Only

- `description` (String) A name or description for the prefix
- `origin_interfaces` (List of String) Interfaces associated with the prefix
- `origin_type` (String) Type of the origin
- `prefix` (String) A static IPv6 prefix
