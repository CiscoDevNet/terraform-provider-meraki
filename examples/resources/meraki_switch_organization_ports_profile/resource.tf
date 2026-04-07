resource "meraki_switch_organization_ports_profile" "example" {
  organization_id            = "123456"
  description                = "IP Phones for all office workers"
  is_organization_wide       = true
  name                       = "Phone"
  port_access_policy_type    = "Open"
  port_allowed_vlans         = "1-100"
  port_dai_trusted           = false
  port_isolation_enabled     = false
  port_peer_sgt_capable      = false
  port_poe_enabled           = true
  port_rstp_enabled          = true
  port_storm_control_enabled = true
  port_stp_guard             = "disabled"
  port_type                  = "access"
  port_udld                  = "Alert only"
  port_vlan                  = 10
  port_voice_vlan            = 20
  tags                       = ["tag1"]
}
