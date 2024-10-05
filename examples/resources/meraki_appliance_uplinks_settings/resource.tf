resource "meraki_appliance_uplinks_settings" "example" {
  serial                                        = "1234-ABCD-1234"
  interfaces_wan1_enabled                       = true
  interfaces_wan1_pppoe_enabled                 = true
  interfaces_wan1_pppoe_authentication_enabled  = true
  interfaces_wan1_pppoe_authentication_password = "password"
  interfaces_wan1_pppoe_authentication_username = "username"
  interfaces_wan1_svis_ipv4_address             = "9.10.11.10"
  interfaces_wan1_svis_ipv4_assignment_mode     = "static"
  interfaces_wan1_svis_ipv6_address             = "2001:2:3::4"
  interfaces_wan1_svis_ipv6_assignment_mode     = "static"
  interfaces_wan1_vlan_tagging_enabled          = true
  interfaces_wan1_vlan_tagging_vlan_id          = 1
  interfaces_wan2_enabled                       = true
  interfaces_wan2_pppoe_enabled                 = true
  interfaces_wan2_pppoe_authentication_enabled  = true
  interfaces_wan2_pppoe_authentication_password = "password"
  interfaces_wan2_pppoe_authentication_username = "username"
  interfaces_wan2_svis_ipv4_address             = "9.10.11.10"
  interfaces_wan2_svis_ipv4_assignment_mode     = "static"
  interfaces_wan2_svis_ipv6_address             = "2001:2:3::4"
  interfaces_wan2_svis_ipv6_assignment_mode     = "static"
  interfaces_wan2_vlan_tagging_enabled          = true
  interfaces_wan2_vlan_tagging_vlan_id          = 1
}
