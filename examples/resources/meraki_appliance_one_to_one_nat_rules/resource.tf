resource "meraki_appliance_one_to_one_nat_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      lan_ip = "192.168.128.22"
      name = "Service behind NAT"
      public_ip = "146.12.3.33"
      uplink = "internet1"
        allowed_inbound = [
          {
          protocol = "tcp"
          allowed_ips = ["10.82.112.0/24"]
          destination_ports = ["80"]
          }
        ]
    }
  ]
}
