resource "meraki_switch_link_aggregation" "example" {
  network_id = "L_123456"
  switch_ports = [
    {
      port_id = "1"
      serial  = "Q234-ABCD-0001"
    }
  ]
  switch_profile_ports = [
    {
      port_id = "2"
      profile = "1234"
    }
  ]
}