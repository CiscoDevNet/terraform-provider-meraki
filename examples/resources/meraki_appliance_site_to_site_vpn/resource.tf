resource "meraki_appliance_site_to_site_vpn" "example" {
  network_id = "L_123456"
  mode       = "hub"
  subnets = [
    {
      local_subnet = "192.168.128.0/24"
      use_vpn      = true
    }
  ]
}
