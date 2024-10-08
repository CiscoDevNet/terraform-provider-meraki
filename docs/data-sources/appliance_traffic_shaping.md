---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_traffic_shaping Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance Traffic Shaping configuration.
---

# meraki_appliance_traffic_shaping (Data Source)

This data source can read the `Appliance Traffic Shaping` configuration.

## Example Usage

```terraform
data "meraki_appliance_traffic_shaping" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `global_bandwidth_limit_down` (Number) The download bandwidth limit in Kbps. (0 represents no limit.)
- `global_bandwidth_limit_up` (Number) The upload bandwidth limit in Kbps. (0 represents no limit.)
- `id` (String) The id of the object
