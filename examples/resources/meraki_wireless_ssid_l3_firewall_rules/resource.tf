resource "meraki_wireless_ssid_l3_firewall_rules" "example" {
  network_id = "L_123456"
  number = "0"
  allow_lan_access = true
  rules = [
    {
      comment = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr = "Any"
      dest_port = "443"
      policy = "allow"
      protocol = "tcp"
      ip_version = "ipv4"
    }
  ]
}
