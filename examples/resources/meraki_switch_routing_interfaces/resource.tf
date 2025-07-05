resource "meraki_switch_routing_interfaces" "example" {
  serial          = "1234-ABCD-1234"
  organization_id = "123456"
  items = [{
    default_gateway      = "192.168.1.1"
    interface_ip         = "192.168.1.2"
    multicast_routing    = "disabled"
    name                 = "L3 interface"
    subnet               = "192.168.1.0/24"
    vlan_id              = 100
    ipv6_address         = "1:2:3:4::1"
    ipv6_assignment_mode = "static"
    ipv6_gateway         = "1:2:3:4::2"
    ipv6_prefix          = "1:2:3:4::/64"
    ospf_settings_area   = "ospfDisabled"
  }]
}
