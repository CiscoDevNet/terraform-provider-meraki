resource "meraki_wireless_zigbee" "example" {
  network_id = "L_123456"
  enabled = false
  defaults_channel = "25"
  defaults_transmit_power_level = 10
  iot_controller_serial = "Q234-ABCD-5678"
  lock_management_address = "10.100.100.200"
  lock_management_password = "password"
  lock_management_username = "user"
}
