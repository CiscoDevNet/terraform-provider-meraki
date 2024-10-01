resource "meraki_appliance_port_forwarding_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      lan_ip      = "192.168.128.1"
      local_port  = "442-443"
      name        = "Description of Port Forwarding Rule"
      protocol    = "tcp"
      public_port = "8100-8101"
      uplink      = "both"
      allowed_ips = ["any"]
    }
  ]
}
