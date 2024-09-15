resource "meraki_network_settings" "example" {
  network_id                                = "L_123456"
  local_status_page_enabled                 = true
  remote_status_page_enabled                = true
  local_status_page_authentication_enabled  = false
  local_status_page_authentication_password = "miles123"
  named_vlans_enabled                       = true
  secure_port_enabled                       = false
}
