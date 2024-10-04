resource "meraki_appliance_single_lan" "example" {
  network_id             = "L_123456"
  appliance_ip           = "192.168.128.1"
  subnet                 = "192.168.128.0/24"
  ipv6_enabled           = true
  mandatory_dhcp_enabled = false
}
