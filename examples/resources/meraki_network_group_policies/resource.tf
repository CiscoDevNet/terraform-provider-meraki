resource "meraki_network_group_policies" "example" {
  network_id                  = "L_123456"
  name                        = "test_group_policy"
  bonjour_forwarding_settings = "custom"
  bonjour_forwarding_rules = [
    {
      description = "a simple bonjour rule"
      services    = ["All Services"]
      vlan_id     = "2"
    }
  ]
}
