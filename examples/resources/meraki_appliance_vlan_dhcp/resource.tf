resource "meraki_appliance_vlan_dhcp" "example" {
  network_id                = "L_123456"
  vlan_id                   = "1234"
  dhcp_boot_filename        = "sample.file"
  dhcp_boot_next_server     = "1.2.3.4"
  dhcp_boot_options_enabled = true
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "1 day"
  dns_nameservers           = "upstream_dns"
  fixed_ip_assignments = {
    "22:33:44:55:66:77" = {
      ip   = "192.168.1.100"
      name = "Some client name"
    }
  }
  mandatory_dhcp_enabled = true
}
