resource "meraki_appliance_vlan_dhcp" "example" {
  network_id                = "L_123456"
  vlan_id                   = "1234"
  dhcp_boot_options_enabled = false
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "1 day"
  dns_nameservers           = "upstream_dns"
  mandatory_dhcp_enabled    = true
}
