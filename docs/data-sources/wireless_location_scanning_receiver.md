---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_location_scanning_receiver Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless Location Scanning Receiver configuration.
---

# meraki_wireless_location_scanning_receiver (Data Source)

This data source can read the `Wireless Location Scanning Receiver` configuration.

## Example Usage

```terraform
data "meraki_wireless_location_scanning_receiver" "example" {
  id              = "12345678"
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object
- `organization_id` (String) Organization ID

### Read-Only

- `network_id` (String) Network ID
- `radio_type` (String) Radio Type whether WiFi or Bluetooth
- `shared_secret` (String) Secret Value for Receiver
- `url` (String) Receiver Url
- `version` (String) Scanning API Version
