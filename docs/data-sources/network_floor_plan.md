---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_floor_plan Data Source - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This data source can read the Network Floor Plan configuration.
---

# meraki_network_floor_plan (Data Source)

This data source can read the `Network Floor Plan` configuration.

## Example Usage

```terraform
data "meraki_network_floor_plan" "example" {
  id         = "12345678"
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `id` (String) The id of the object
- `name` (String) The name of your floor plan.

### Read-Only

- `bottom_left_corner_lat` (Number) Latitude
- `bottom_left_corner_lng` (Number) Longitude
- `bottom_right_corner_lat` (Number) Latitude
- `bottom_right_corner_lng` (Number) Longitude
- `center_lat` (Number) Latitude
- `center_lng` (Number) Longitude
- `image_contents` (String) The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
- `top_left_corner_lat` (Number) Latitude
- `top_left_corner_lng` (Number) Longitude
- `top_right_corner_lat` (Number) Latitude
- `top_right_corner_lng` (Number) Longitude