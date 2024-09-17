resource "meraki_switch_access_policy" "example" {
  network_id                               = "L_123456"
  access_policy_type                       = "Hybrid authentication"
  guest_port_bouncing                      = true
  guest_vlan_id                            = 100
  host_mode                                = "Single-Host"
  increase_access_speed                    = false
  name                                     = "Access policy #1"
  radius_accounting_enabled                = true
  radius_coa_support_enabled               = false
  radius_group_attribute                   = "11"
  radius_testing_enabled                   = true
  url_redirect_walled_garden_enabled       = true
  voice_vlan_clients                       = true
  dot1x_control_direction                  = "inbound"
  radius_failed_auth_vlan_id               = 100
  radius_re_authentication_interval        = 120
  radius_cache_enabled                     = false
  radius_cache_timeout                     = 24
  radius_critical_auth_data_vlan_id        = 100
  radius_critical_auth_suspend_port_bounce = true
  radius_critical_auth_voice_vlan_id       = 101
  radius_accounting_servers = [
    {
      host   = "1.2.3.4"
      port   = 22
      secret = "secret"
    }
  ]
  radius_servers = [
    {
      host   = "1.2.3.4"
      port   = 22
      secret = "secret"
    }
  ]
  url_redirect_walled_garden_ranges = ["192.168.1.0/24"]
}
