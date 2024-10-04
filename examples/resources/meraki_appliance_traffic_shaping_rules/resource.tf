resource "meraki_appliance_traffic_shaping_rules" "example" {
  network_id            = "L_123456"
  default_rules_enabled = true
  rules = [
    {
      dscp_tag_value                      = 0
      priority                            = "normal"
      per_client_bandwidth_limit_settings = "custom"
      per_client_bandwidth_limit_down     = 1000000
      per_client_bandwidth_limit_up       = 1000000
      definitions = [
        {
          type  = "host"
          value = "google.com"
        }
      ]
    }
  ]
}
