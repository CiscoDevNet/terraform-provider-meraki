resource "meraki_wireless_zigbee_device" "example" {
  organization_id = "123456"
  device_id = "1"
  channel = "11"
  enrolled = true
}
