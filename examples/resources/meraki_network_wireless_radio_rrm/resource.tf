resource "meraki_network_wireless_radio_rrm" "example" {
  network_id                         = "L_123456"
  ai_enabled                         = true
  busy_hour_minimize_changes_enabled = true
  busy_hour_schedule_mode            = "automatic"
  busy_hour_schedule_manual_end      = "15:00"
  busy_hour_schedule_manual_start    = "10:00"
  channel_avoidance_enabled          = true
  fra_enabled                        = false
}
