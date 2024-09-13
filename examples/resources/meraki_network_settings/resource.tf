resource "meraki_network_settings" "example" {
  remote_status_page_enabled                = true
  local_status_page_authentication_enabled  = false
  local_status_page_authentication_password = "miles123"
  secure_port_enabled                       = false
  named_vlans_enabled                       = true
  local_status_page_enabled                 = true
  network_id                                = "L_123456"
}
