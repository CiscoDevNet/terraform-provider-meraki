resource "meraki_appliance_inbound_firewall_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      comment = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr = "2001::/64"
      dest_port = "443"
      policy = "allow"
      protocol = "tcp"
      src_cidr = "Any"
      src_port = "Any"
      syslog_enabled = false
    }
  ]
}
