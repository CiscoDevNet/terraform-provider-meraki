resource "meraki_appliance_site_to_site_vpn" "example" {
  network_id                = "L_123456"
  confirmed                 = true
  mode                      = "hub"
  nat_traversal_mode        = "manual"
  nat_traversal_public_ip   = "203.0.113.8"
  nat_traversal_public_port = 4500
  sgt_enabled               = true
  subnet_nat_is_allowed     = true
  host_translations = [
    {
      name           = "Host 1"
      local_address  = "192.168.128.10"
      remote_address = "10.128.0.10"
    }
  ]
  subnets = [
    {
      local_subnet      = "192.168.128.0/24"
      use_vpn           = true
      nat_enabled       = true
      nat_remote_subnet = "10.128.0.0/24"
    }
  ]
}
