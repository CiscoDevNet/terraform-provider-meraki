resource "meraki_switch_routing_static_route" "example" {
  serial                          = "1234-ABCD-1234"
  advertise_via_ospf_enabled      = false
  name                            = "My route"
  next_hop_ip                     = "192.168.1.1"
  prefer_over_ospf_routes_enabled = false
  subnet                          = "192.168.2.0/24"
}
