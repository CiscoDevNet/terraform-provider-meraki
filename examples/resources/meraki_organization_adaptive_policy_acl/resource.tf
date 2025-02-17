resource "meraki_organization_adaptive_policy_acl" "example" {
  organization_id = "123456"
  description     = "Blocks sensitive web traffic"
  ip_version      = "ipv6"
  name            = "Block sensitive web traffic"
  rules = [
    {
      dst_port        = "22-30"
      log             = true
      policy          = "deny"
      protocol        = "tcp"
      src_port        = "1,33"
      tcp_established = true
    }
  ]
}
