resource "meraki_appliance_site_to_site_vpn" "example" {
  network_id            = "L_123456"
  mode                  = "hub"
  subnet_nat_is_allowed = false
  host_translations = [
    {
      name           = "Host 1"
      local_address  = "192.168.1.10"
      remote_address = "72.168.2.10"
    }
  ]
  subnets = [
    {
      local_subnet = "192.168.128.0/24"
      use_vpn      = true
    }
  ]
}
