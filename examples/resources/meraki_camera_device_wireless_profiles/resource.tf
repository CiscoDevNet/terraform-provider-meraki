resource "meraki_camera_device_wireless_profiles" "example" {
  serial        = "1234-ABCD-1234"
  ids_backup    = "1"
  ids_primary   = "3"
  ids_secondary = "2"
}
