resource "meraki_network_group_policy" "example" {
  splash_auth_settings        = "bypass"
  vlan_tagging_settings       = "custom"
  vlan_tagging_vlan_id        = "1"
  bonjour_forwarding_settings = "custom"
  bonjour_forwarding_rules = [
    {
      description = "A simple bonjour rule"
      vlan_id     = "1"
      services    = ["All Services"]
    }
  ]
  name                                  = "No video streaming"
  scheduling_monday_to                  = "17:00"
  scheduling_monday_active              = true
  scheduling_monday_from                = "09:00"
  scheduling_tuesday_active             = true
  scheduling_tuesday_from               = "09:00"
  scheduling_tuesday_to                 = "17:00"
  scheduling_wednesday_to               = "17:00"
  scheduling_wednesday_active           = true
  scheduling_wednesday_from             = "09:00"
  scheduling_thursday_active            = true
  scheduling_thursday_from              = "09:00"
  scheduling_thursday_to                = "17:00"
  scheduling_friday_active              = true
  scheduling_friday_from                = "09:00"
  scheduling_friday_to                  = "17:00"
  scheduling_saturday_to                = "17:00"
  scheduling_saturday_active            = true
  scheduling_saturday_from              = "09:00"
  scheduling_sunday_from                = "09:00"
  scheduling_sunday_to                  = "17:00"
  scheduling_sunday_active              = true
  scheduling_enabled                    = true
  bandwidth_settings                    = "custom"
  bandwidth_bandwidth_limits_limit_up   = 1000000
  bandwidth_bandwidth_limits_limit_down = 1000000
  firewall_and_traffic_shaping_l7_firewall_rules = [
    {
      policy = "deny"
      type   = "host"
      value  = "google.com"
    }
  ]
  firewall_and_traffic_shaping_settings = "custom"
  firewall_and_traffic_shaping_traffic_shaping_rules = [
    {
      definitions = [
        {
          type  = "host"
          value = "google.com"
        }
      ]
      per_client_bandwidth_limits_settings                    = "custom"
      per_client_bandwidth_limits_bandwidth_limits_limit_up   = 1000000
      per_client_bandwidth_limits_bandwidth_limits_limit_down = 1000000
      dscp_tag_value                                          = 0
      pcp_tag_value                                           = 0
      priority                                                = "normal"
    }
  ]
  firewall_and_traffic_shaping_l3_firewall_rules = [
    {
      protocol  = "tcp"
      dest_port = "443"
      dest_cidr = "192.168.1.0/24"
      comment   = "Allow TCP traffic to subnet with HTTP servers."
      policy    = "allow"
    }
  ]
  content_filtering_allowed_url_patterns_settings     = "network default"
  content_filtering_allowed_url_patterns_patterns     = ["http://www.example.com"]
  content_filtering_blocked_url_patterns_settings     = "append"
  content_filtering_blocked_url_patterns_patterns     = ["http://www.example.com"]
  content_filtering_blocked_url_categories_settings   = "override"
  content_filtering_blocked_url_categories_categories = ["meraki:contentFiltering/category/1"]
  network_id                                          = ""
}
