---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_location_scanning Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless Location Scanning configuration.
---

# meraki_wireless_location_scanning (Data Source)

This data source can read the `Wireless Location Scanning` configuration.

## Example Usage

```terraform
data "meraki_wireless_location_scanning" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `api_enabled` (Boolean) Enable push API for scanning events, analytics must be enabled
- `enabled` (Boolean) Collect location and scanning analytics
- `id` (String) The id of the object
