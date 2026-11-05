resource "meraki_appliance_uplinks_nat" "example" {
  network_id = "L_123456"
  uplinks = [
    {
      interface   = "wan1"
      nat_enabled = false
    }
  ]
}
