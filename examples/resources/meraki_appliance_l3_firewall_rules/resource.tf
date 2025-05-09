resource "meraki_appliance_l3_firewall_rules" "example" {
  network_id          = "L_123456"
  syslog_default_rule = false
  rules = [
    {
      comment        = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr      = "192.168.1.0/24"
      dest_port      = "443"
      policy         = "allow"
      protocol       = "tcp"
      src_cidr       = "Any"
      src_port       = "Any"
      syslog_enabled = false
    }
  ]
}
