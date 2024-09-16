resource "meraki_switch_routing_multicast_rendezvous_point" "example" {
  network_id      = "L_123456"
  interface_ip    = "192.168.1.2"
  multicast_group = "Any"
}
