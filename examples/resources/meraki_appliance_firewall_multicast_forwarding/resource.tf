resource "meraki_appliance_firewall_multicast_forwarding" "example" {
  network_id = "L_123456"
  rules = [
    {
      address = "224.0.0.1"
      description = "test"
      vlan_ids = ["1"]
    }
  ]
}
