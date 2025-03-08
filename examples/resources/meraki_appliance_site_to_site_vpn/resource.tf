resource "meraki_appliance_site_to_site_vpn" "example" {
  network_id            = "L_123456"
  mode                  = "hub"
  subnet_nat_is_allowed = true
  subnets = [
    {
      local_subnet      = "192.168.128.0/24"
      use_vpn           = true
      nat_enabled       = true
      nat_remote_subnet = "192.168.2.0/24"
    }
  ]
}
