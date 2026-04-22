resource "meraki_switch_stack_routing_interface" "example" {
  network_id            = "L_123456"
  switch_stack_id       = "1234"
  candidate_uplink_v4   = true
  default_gateway       = "192.168.1.1"
  interface_ip          = "192.168.1.2"
  mode                  = "vlan"
  multicast_routing     = "disabled"
  name                  = "L3 interface"
  static_v4_dns1        = "8.8.8.8"
  static_v4_dns2        = "8.8.4.4"
  subnet                = "192.168.1.0/24"
  uplink_v4             = true
  uplink_v6             = true
  vlan_id               = 100
  ipv6_address          = "1:2:3:4::1"
  ipv6_assignment_mode  = "static"
  ipv6_candidate_uplink = true
  ipv6_gateway          = "1:2:3:4::2"
  ipv6_prefix           = "1:2:3:4::/64"
  ipv6_static_v6_dns1   = "2001:db8::1234"
  ipv6_static_v6_dns2   = "2001:db8::8888"
  ospf_settings_area    = "ospfDisabled"
}
