resource "meraki_switch_access_control_lists" "example" {
  network_id = "L_123456"
  rules = [
    {
      comment = "Deny SSH"
      dst_cidr = "172.16.3.0/24"
      dst_port = "22"
      ip_version = "ipv4"
      policy = "deny"
      protocol = "tcp"
      src_cidr = "10.1.10.0/24"
      src_port = "any"
      vlan = "10"
    }
  ]
}
