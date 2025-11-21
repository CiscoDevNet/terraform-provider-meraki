resource "meraki_wireless_billing" "example" {
  network_id = "L_123456"
  currency = "USD"
  plans = [
    {
      price = 5
      time_limit = "1 hour"
      bandwidth_limits_limit_down = 1000000
      bandwidth_limits_limit_up = 1000000
    }
  ]
}
