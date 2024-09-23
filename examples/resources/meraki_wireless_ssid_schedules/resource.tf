resource "meraki_wireless_ssid_schedules" "example" {
  network_id = "L_123456"
  number     = "0"
  enabled    = true
  ranges = [
    {
      end_day    = "Tuesday"
      end_time   = "05:00:00"
      start_day  = "Tuesday"
      start_time = "01:00:00"
    }
  ]
}
