resource "meraki_switch_dhcp_server_policy_arp_inspection_trusted_server" "example" {
  network_id   = "L_123456"
  mac          = "00:11:22:33:44:55"
  vlan         = 100
  ipv4_address = "1.2.3.4"
}
