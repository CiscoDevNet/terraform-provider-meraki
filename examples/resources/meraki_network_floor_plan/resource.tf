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
