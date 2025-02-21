resource "meraki_appliance_cellular_inbound_firewall_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      comment        = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr      = "Any"
      dest_port      = "443"
      policy         = "allow"
      protocol       = "tcp"
      src_cidr       = "Any"
      src_port       = "Any"
      syslog_enabled = false
    }
  ]
}
