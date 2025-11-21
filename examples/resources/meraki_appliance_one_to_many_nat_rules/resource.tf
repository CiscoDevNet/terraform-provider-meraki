resource "meraki_appliance_one_to_many_nat_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      public_ip = "146.11.11.13"
      uplink = "internet1"
        port_rules = [
          {
          local_ip = "192.168.128.1"
          local_port = "443"
          name = "Rule 1"
          protocol = "tcp"
          public_port = "9443"
          allowed_ips = ["any"]
          }
        ]
    }
  ]
}
