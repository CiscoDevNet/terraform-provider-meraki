resource "meraki_appliance_l7_firewall_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      policy = "deny"
      type = "host"
      value = "google.com"
    }
  ]
}
