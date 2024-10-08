---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization Data Source - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This data source can read the Organization configuration.
---

# meraki_organization (Data Source)

This data source can read the `Organization` configuration.

## Example Usage

```terraform
data "meraki_organization" "example" {
  id = "12345678"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the organization

### Read-Only

- `management_details` (Attributes List) Details related to organization management, possibly empty (see [below for nested schema](#nestedatt--management_details))

<a id="nestedatt--management_details"></a>
### Nested Schema for `management_details`

Read-Only:

- `name` (String) Name of management data
- `value` (String) Value of management data
