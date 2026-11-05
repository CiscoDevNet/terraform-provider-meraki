resource "meraki_appliance_traffic_shaping_uplink_selection" "example" {
  network_id                              = "L_123456"
  default_uplink                          = "wan1"
  load_balancing_enabled                  = true
  failover_and_failback_immediate_enabled = true
  wan_traffic_uplink_preferences = [
    {
      preferred_uplink = "wan1"
      traffic_filters = [
        {
          type             = "custom"
          protocol         = "tcp"
          destination_cidr = "any"
          destination_port = "any"
          source_cidr      = "any"
          source_port      = "1-1024"
        }
      ]
    }
  ]
}
