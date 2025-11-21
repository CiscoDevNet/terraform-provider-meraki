resource "meraki_cellular_gateway_dhcp" "example" {
  network_id = "L_123456"
  dhcp_lease_time = "1 hour"
  dns_nameservers = "custom"
  dns_custom_nameservers = ["172.16.2.111"]
}
