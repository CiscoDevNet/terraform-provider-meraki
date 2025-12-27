resource "meraki_wireless_ssid_open_roaming" "example" {
  network_id = "L_123456"
  number     = "0"
  enabled    = false
  tenant_id  = "12345"
}
