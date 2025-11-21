resource "meraki_wireless_ssid_traffic_shaping_rules" "example" {
  network_id = "L_123456"
  number = "0"
  default_rules_enabled = true
  traffic_shaping_enabled = true
  rules = [
    {
      dscp_tag_value = 0
      pcp_tag_value = 0
      per_client_bandwidth_limits_settings = "custom"
      per_client_bandwidth_limits_bandwidth_limits_limit_down = 1000000
      per_client_bandwidth_limits_bandwidth_limits_limit_up = 1000000
        definitions = [
          {
          type = "host"
          value = "google.com"
          }
        ]
    }
  ]
}
