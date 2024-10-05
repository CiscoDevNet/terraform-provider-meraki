resource "meraki_appliance_vpn_bgp" "example" {
  network_id      = "L_123456"
  as_number       = 64515
  enabled         = true
  ibgp_hold_timer = 120
  neighbors = [
    {
      allow_transit           = true
      ebgp_hold_timer         = 180
      ebgp_multihop           = 2
      ip                      = "10.10.10.22"
      next_hop_ip             = "1.2.3.4"
      receive_limit           = 120
      remote_as_number        = 64343
      source_interface        = "wan1"
      authentication_password = "abc123"
      ipv6_address            = "2002::1234:abcd:ffff:c0a8:101"
      ttl_security_enabled    = false
    }
  ]
}
