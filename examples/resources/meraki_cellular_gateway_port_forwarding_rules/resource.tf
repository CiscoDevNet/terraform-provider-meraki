resource "meraki_cellular_gateway_port_forwarding_rules" "example" {
  serial = "1234-ABCD-1234"
  rules = [
    {
      access = "restricted"
      lan_ip = "172.31.128.5"
      local_port = "4"
      name = "test"
      protocol = "tcp"
      public_port = "11-12"
      allowed_ips = ["10.10.10.10"]
    }
  ]
}
