resource "meraki_switch_settings" "example" {
  network_id                        = "L_123456"
  use_combined_power                = false
  vlan                              = 1
  mac_blocklist_enabled             = true
  uplink_client_sampling_enabled    = false
  uplink_selection_candidates       = "all"
  uplink_selection_failback_enabled = false
}
