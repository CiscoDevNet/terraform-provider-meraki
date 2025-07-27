resource "meraki_switch_link_aggregations" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    switch_ports = [
      {
        port_id = "1"
        serial  = "Q234-ABCD-0001"
      }
    ]
  }]
}
