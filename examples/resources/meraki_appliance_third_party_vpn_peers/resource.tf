resource "meraki_appliance_third_party_vpn_peers" "example" {
  organization_id = "123456"
  peers = [
    {
      ike_version = "2"
      is_route_based = false
      local_id = "myMXId@meraki.com"
      name = "Peer Name"
      priority_in_group = 1
      public_ip = "123.123.123.1"
      remote_id = "miles@meraki.com"
      secret = "Sample Password"
      group_active_active_tunnel = false
      group_number = 1
      group_failover_direct_to_internet = false
      ipsec_policies_child_lifetime = 28800
      ipsec_policies_ike_lifetime = 28800
      ipsec_policies_child_auth_algo = ["sha1"]
      ipsec_policies_child_cipher_algo = ["aes128"]
      ipsec_policies_child_pfs_group = ["disabled"]
      ipsec_policies_ike_auth_algo = ["sha1"]
      ipsec_policies_ike_cipher_algo = ["tripledes"]
      ipsec_policies_ike_diffie_hellman_group = ["group2"]
      ipsec_policies_ike_prf_algo = ["prfsha1"]
      network_tags = ["none"]
      private_subnets = ["192.168.1.0/24"]
    }
  ]
}
