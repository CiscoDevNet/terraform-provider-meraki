resource "meraki_appliance_vlans" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    appliance_ip = "192.168.1.2"
    dhcp_boot_options_enabled = false
    dhcp_handling = "Run a DHCP server"
    dhcp_lease_time = "1 day"
    dns_nameservers = "upstream_dns"
    vlan_id = "1234"
    name = "My VLAN"
    subnet = "192.168.1.0/24"
    ipv6_enabled = true
    mandatory_dhcp_enabled = true
  }]
}
