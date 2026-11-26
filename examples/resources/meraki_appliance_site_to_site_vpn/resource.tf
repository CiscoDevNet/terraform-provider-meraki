resource "meraki_appliance_site_to_site_vpn" "example" {
  network_id                = "L_123456"
  confirmed                 = true
  mode                      = "hub"
  nat_traversal_mode        = "manual"
  nat_traversal_public_ip   = "203.0.113.8"
  nat_traversal_public_port = 4500
  sgt_enabled               = true
  subnet_nat_is_allowed     = false
  subnets = [
    {
      local_subnet = "192.168.128.0/24"
      use_vpn      = true
    }
  ]
}
