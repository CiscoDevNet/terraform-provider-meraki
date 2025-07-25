---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization Resource - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This resource can manage the Organization configuration.
---

# meraki_organization (Resource)

This resource can manage the `Organization` configuration.

## Example Usage

```terraform
resource "meraki_organization" "example" {
  name = "My organization"
  management_details = [
    {
      name  = "MSP ID"
      value = "123456"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the organization

### Optional

- `management_details` (Attributes List) Details related to organization management, possibly empty (see [below for nested schema](#nestedatt--management_details))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--management_details"></a>
### Nested Schema for `management_details`

Required:

- `name` (String) Name of management data
- `value` (String) Value of management data

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_organization.example "<id>"
```
