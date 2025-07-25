---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_floor_plan Resource - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This resource can manage the Network Floor Plan configuration.
---

# meraki_network_floor_plan (Resource)

This resource can manage the `Network Floor Plan` configuration.

## Example Usage

```terraform
resource "meraki_network_floor_plan" "example" {
  network_id              = "L_123456"
  floor_number            = 0
  image_contents          = "R0lGODdhAQABAIEAAP///wAAAAAAAAAAACwAAAAAAQABAAAIBAABBAQAOw=="
  name                    = "HQ Floor Plan"
  bottom_left_corner_lat  = 37.770040510499996
  bottom_left_corner_lng  = -122.38714009525
  bottom_right_corner_lat = 37.770040510499996
  bottom_right_corner_lng = -121.38714009525
  top_left_corner_lat     = 38.770040510499996
  top_left_corner_lng     = -122.38714009525
  top_right_corner_lat    = 38.770040510499996
  top_right_corner_lng    = -121.38714009525
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image_contents` (String) The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
- `name` (String) The name of your floor plan.
- `network_id` (String) Network ID

### Optional

- `bottom_left_corner_lat` (Number) Latitude
- `bottom_left_corner_lng` (Number) Longitude
- `bottom_right_corner_lat` (Number) Latitude
- `bottom_right_corner_lng` (Number) Longitude
- `center_lat` (Number) Latitude
- `center_lng` (Number) Longitude
- `floor_number` (Number) The floor number of the floors within the building
- `top_left_corner_lat` (Number) Latitude
- `top_left_corner_lng` (Number) Longitude
- `top_right_corner_lat` (Number) Latitude
- `top_right_corner_lng` (Number) Longitude

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_network_floor_plan.example "<network_id>,<id>"
```
