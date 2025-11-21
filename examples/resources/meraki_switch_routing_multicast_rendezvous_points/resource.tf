resource "meraki_switch_routing_multicast_rendezvous_points" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    interface_ip = "192.168.1.2"
    multicast_group = "Any"
  }]
}
