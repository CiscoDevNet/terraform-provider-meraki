resource "meraki_wireless_network_bluetooth_settings" "example" {
  network_id                  = "L_123456"
  advertising_enabled         = true
  major                       = 1
  major_minor_assignment_mode = "Non-unique"
  minor                       = 1
  scanning_enabled            = true
  uuid                        = "00000000-0000-0000-0000-000000000000"
}
