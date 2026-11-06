resource "meraki_appliance_vpn_site_to_site_hub_vrfs" "example" {
  network_id     = "L_123456"
  hub_network_id = "L_234567"
  vrfs = [
    {
      id = "12345"
    }
  ]
}
