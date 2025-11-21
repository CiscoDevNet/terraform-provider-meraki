resource "meraki_switch_routing_static_route" "example" {
  serial = "1234-ABCD-1234"
  name = "My route"
  next_hop_ip = "192.168.1.1"
  subnet = "192.168.2.0/24"
}
