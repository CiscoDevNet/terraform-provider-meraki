resource "meraki_appliance_vpn_bgp" "example" {
  network_id                      = "L_123456"
  as_number                       = 64515
  enabled                         = true
  ibgp_hold_timer                 = 120
  local_as_number                 = 64514
  priority_route                  = "eBGP"
  router_id                       = "10.15.10.2"
  ipv6_single_peering_enabled     = true
  tunnel_down_termination_enabled = true
  neighbors = [
    {
      allow_transit                                 = true
      ebgp_hold_timer                               = 180
      ebgp_multihop                                 = 2
      ip                                            = "10.10.10.22"
      multi_exit_discriminator                      = 2
      next_hop_ip                                   = "1.2.3.4"
      receive_limit                                 = 120
      remote_as_number                              = 64343
      source_interface                              = "wan1"
      weight                                        = 10
      authentication_password                       = "abc123"
      explicit_outbound_route_advertisement_enabled = true
      filter_out                                    = ["10.0.0.0/8"]
      ipv6_address                                  = "2002::1234:abcd:ffff:c0a8:101"
      ttl_security_enabled                          = false
      community_out                                 = ["64515:100"]
      filter_in                                     = ["10.0.0.0/8"]
      path_prepend                                  = [1]
    }
  ]
}
