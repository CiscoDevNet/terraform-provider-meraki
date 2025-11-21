resource "meraki_switch_dhcp_server_policy" "example" {
  network_id = "L_123456"
  default_policy = "block"
  alerts_email_enabled = true
  arp_inspection_enabled = true
  allowed_servers = ["00:50:56:00:00:01"]
  blocked_servers = ["00:50:56:00:00:03"]
}
