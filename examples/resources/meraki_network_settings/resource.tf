resource "meraki_network_settings" "example" {
  network_id                                = "L_123456"
  local_status_page_enabled                 = false
  remote_status_page_enabled                = false
  local_status_page_authentication_enabled  = false
  local_status_page_authentication_password = "miles123"
  named_vlans_enabled                       = false
  secure_port_enabled                       = false
}
