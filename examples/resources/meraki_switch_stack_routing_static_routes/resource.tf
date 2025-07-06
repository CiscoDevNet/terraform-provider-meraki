resource "meraki_switch_stack_routing_static_routes" "example" {
  network_id      = "L_123456"
  switch_stack_id = "1234"
  organization_id = "123456"
  items = [{
    name        = "My route"
    next_hop_ip = "192.168.1.1"
    subnet      = "192.168.2.0/24"
  }]
}
