resource "meraki_sensor_relationships" "example" {
  serial = "1234-ABCD-1234"
  livestream_related_devices = [
    {
      serial = ""
    }
  ]
}
