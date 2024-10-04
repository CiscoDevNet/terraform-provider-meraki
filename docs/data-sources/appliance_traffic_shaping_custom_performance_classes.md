---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_traffic_shaping_custom_performance_classes Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Traffic Shaping Custom Performance Class configuration.
---

# meraki_appliance_traffic_shaping_custom_performance_classes (Data Source)

This data source can read the `Appliance Traffic Shaping Custom Performance Class` configuration.

## Example Usage

```terraform
data "meraki_appliance_traffic_shaping_custom_performance_classes" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `items` (Attributes List) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String) The id of the object
- `max_jitter` (Number) Maximum jitter in milliseconds
- `max_latency` (Number) Maximum latency in milliseconds
- `max_loss_percentage` (Number) Maximum percentage of packet loss
- `name` (String) Name of the custom performance class