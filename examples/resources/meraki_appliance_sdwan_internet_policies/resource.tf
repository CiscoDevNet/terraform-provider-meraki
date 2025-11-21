resource "meraki_appliance_sdwan_internet_policies" "example" {
  network_id = "L_123456"
  wan_traffic_uplink_preferences = [
    {
      fail_over_criterion = "poorPerformance"
      preferred_uplink = "wan1"
      builtin_performance_class_name = "VoIP"
      custom_performance_class_id = "123456"
      performance_class_type = "custom"
        traffic_filters = [
          {
          type = "custom"
          protocol = "tcp"
          destination_cidr = "any"
          destination_port = "any"
            destination_applications = [
              {
                id = "meraki:layer7/application/3"
                name = "DNS"
                type = "major"
              }
            ]
          source_cidr = "any"
          source_host = 254
          source_port = "1-1024"
          source_vlan = 10
          }
        ]
    }
  ]
}
