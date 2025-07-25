---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_traffic_shaping_custom_performance_classes Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance Traffic Shaping Custom Performance Class configuration in bulk.
  This bulk resource uses the following attributes to uniquely identify each object. Changing any of these attributes will cause the object to be deleted and recreated.
  name
---

# meraki_appliance_traffic_shaping_custom_performance_classes (Resource)

This resource can manage the `Appliance Traffic Shaping Custom Performance Class` configuration in bulk.

This bulk resource uses the following attributes to uniquely identify each object. Changing any of these attributes will cause the object to be deleted and recreated.

- `name`

## Example Usage

```terraform
resource "meraki_appliance_traffic_shaping_custom_performance_classes" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    max_jitter          = 100
    max_latency         = 100
    max_loss_percentage = 5
    name                = "myCustomPerformanceClass"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))
- `network_id` (String) Network ID
- `organization_id` (String) The organization ID

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Required:

- `name` (String) Name of the custom performance class

Optional:

- `max_jitter` (Number) Maximum jitter in milliseconds
- `max_latency` (Number) Maximum latency in milliseconds
- `max_loss_percentage` (Number) Maximum percentage of packet loss

Read-Only:

- `id` (String) The id of the item

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_traffic_shaping_custom_performance_classes.example "<organization_id>,<network_id>"
```
