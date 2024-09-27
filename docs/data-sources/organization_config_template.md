---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization_config_template Data Source - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This data source can read the Organization Config Template configuration.
---

# meraki_organization_config_template (Data Source)

This data source can read the `Organization Config Template` configuration.

## Example Usage

```terraform
data "meraki_organization_config_template" "example" {
  id              = "12345678"
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) Organization ID

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the configuration template

### Read-Only

- `copy_from_network_id` (String) The ID of the network or config template to copy configuration from. This attribute is only relevant when creating a new configuration template.
- `time_zone` (String) The timezone of the configuration template. For a list of allowed timezones, please see the `TZ` column in the table in this article. Not applicable if copying from existing network or template