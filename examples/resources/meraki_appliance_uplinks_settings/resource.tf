resource "meraki_appliance_uplinks_settings" "example" {
  serial                                               = "1234-ABCD-1234"
  interfaces_wan1_enabled                              = true
  interfaces_wan1_pppoe_enabled                        = true
  interfaces_wan1_pppoe_authentication_enabled         = true
  interfaces_wan1_pppoe_authentication_password        = "password"
  interfaces_wan1_pppoe_authentication_username        = "username"
  interfaces_wan1_svis_ipv4_address                    = "9.10.11.10"
  interfaces_wan1_svis_ipv4_assignment_mode            = "static"
  interfaces_wan1_svis_ipv6_address                    = "2001:2:3::4"
  interfaces_wan1_svis_ipv6_assignment_mode            = "static"
  interfaces_wan1_transition_enabled                   = true
  interfaces_wan1_transition_mode                      = "MAP-E"
  interfaces_wan1_transition_authentication_enabled    = true
  interfaces_wan1_transition_authentication_password   = "password"
  interfaces_wan1_transition_authentication_username   = "4o6ddns"
  interfaces_wan1_transition_ipv4_local_address        = ""
  interfaces_wan1_transition_ipv6_interface_id         = ""
  interfaces_wan1_transition_ipv6_border_relay_address = ""
  interfaces_wan1_transition_ipv6_border_relay_mode    = "auto"
  interfaces_wan1_transition_isp_name                  = "OCN IPoE"
  interfaces_wan1_transition_map_draft_enabled         = false
  interfaces_wan1_transition_map_rules_source          = "manual"
  interfaces_wan1_transition_map_rules_entries = [
    {
      ea_bit_length        = 20
      ps_id_offset         = 6
      border_relay_address = "2001:db8:ffff::1"
      border_relay_prefix  = ""
      ipv4_prefix_address  = "192.0.2.0"
      ipv4_prefix_length   = 18
      ipv6_prefix_address  = "2001:db8:100::"
      ipv6_prefix_length   = 36
    }
  ]
  interfaces_wan1_vlan_tagging_enabled                 = true
  interfaces_wan1_vlan_tagging_vlan_id                 = 1
  interfaces_wan2_enabled                              = true
  interfaces_wan2_pppoe_enabled                        = true
  interfaces_wan2_pppoe_authentication_enabled         = true
  interfaces_wan2_pppoe_authentication_password        = "password"
  interfaces_wan2_pppoe_authentication_username        = "username"
  interfaces_wan2_svis_ipv4_address                    = "9.10.11.10"
  interfaces_wan2_svis_ipv4_assignment_mode            = "static"
  interfaces_wan2_svis_ipv6_address                    = "2001:2:3::4"
  interfaces_wan2_svis_ipv6_assignment_mode            = "static"
  interfaces_wan2_transition_enabled                   = true
  interfaces_wan2_transition_mode                      = "MAP-E"
  interfaces_wan2_transition_authentication_enabled    = true
  interfaces_wan2_transition_authentication_password   = "password"
  interfaces_wan2_transition_authentication_username   = "4o6ddns"
  interfaces_wan2_transition_ipv4_local_address        = ""
  interfaces_wan2_transition_ipv6_interface_id         = ""
  interfaces_wan2_transition_ipv6_border_relay_address = ""
  interfaces_wan2_transition_ipv6_border_relay_mode    = "auto"
  interfaces_wan2_transition_isp_name                  = "OCN IPoE"
  interfaces_wan2_transition_map_draft_enabled         = false
  interfaces_wan2_transition_map_rules_source          = "manual"
  interfaces_wan2_transition_map_rules_entries = [
    {
      ea_bit_length        = 20
      ps_id_offset         = 6
      border_relay_address = "2001:db8:ffff::1"
      border_relay_prefix  = ""
      ipv4_prefix_address  = "192.0.2.0"
      ipv4_prefix_length   = 18
      ipv6_prefix_address  = "2001:db8:100::"
      ipv6_prefix_length   = 36
    }
  ]
  interfaces_wan2_vlan_tagging_enabled = true
  interfaces_wan2_vlan_tagging_vlan_id = 1
}
