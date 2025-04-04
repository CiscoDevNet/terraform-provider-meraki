---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_camera_custom_analytics Data Source - terraform-provider-meraki"
subcategory: "Cameras"
description: |-
  This data source can read the Camera Custom Analytics configuration.
---

# meraki_camera_custom_analytics (Data Source)

This data source can read the `Camera Custom Analytics` configuration.

## Example Usage

```terraform
data "meraki_camera_custom_analytics" "example" {
  serial = "1234-ABCD-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) Device serial

### Read-Only

- `artifact_id` (String) The ID of the custom analytics artifact
- `enabled` (Boolean) Enable custom analytics
- `id` (String) The id of the object
- `parameters` (Attributes List) Parameters for the custom analytics workload (see [below for nested schema](#nestedatt--parameters))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Read-Only:

- `name` (String) Name of the parameter
- `value` (String) Value of the parameter
