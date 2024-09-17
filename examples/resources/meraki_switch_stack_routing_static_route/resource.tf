resource "meraki_switch_stack_routing_static_route" "example" {
  network_id                      = "L_123456"
  switch_stack_id                 = "1234"
  advertise_via_ospf_enabled      = false
  name                            = "My route"
  next_hop_ip                     = "192.168.1.1"
  prefer_over_ospf_routes_enabled = false
  subnet                          = "192.168.2.0/24"
}
