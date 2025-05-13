resource "meraki_network_group_policy" "example" {
  network_id                  = "L_123456"
  name                        = "No video streaming"
  splash_auth_settings        = "bypass"
  bandwidth_settings          = "custom"
  bandwidth_limit_down        = 1000000
  bandwidth_limit_up          = 1000000
  bonjour_forwarding_settings = "custom"
  bonjour_forwarding_rules = [
    {
      description = "A simple bonjour rule"
      vlan_id     = "1"
      services    = ["All Services"]
    }
  ]
  firewall_and_traffic_shaping_settings = "custom"
  l3_firewall_rules = [
    {
      comment   = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr = "192.168.1.0/24"
      dest_port = "443"
      policy    = "allow"
      protocol  = "tcp"
    }
  ]
  l7_firewall_rules = [
    {
      policy = "deny"
      type   = "host"
      value  = "google.com"
    }
  ]
  traffic_shaping_rules = [
    {
      dscp_tag_value                                          = 0
      pcp_tag_value                                           = 0
      priority                                                = "normal"
      per_client_bandwidth_limits_settings                    = "custom"
      per_client_bandwidth_limits_bandwidth_limits_limit_down = 1000000
      per_client_bandwidth_limits_bandwidth_limits_limit_up   = 1000000
      definitions = [
        {
          type  = "host"
          value = "google.com"
        }
      ]
    }
  ]
  scheduling_enabled          = true
  scheduling_friday_active    = true
  scheduling_friday_from      = "09:00"
  scheduling_friday_to        = "17:00"
  scheduling_monday_active    = true
  scheduling_monday_from      = "09:00"
  scheduling_monday_to        = "17:00"
  scheduling_saturday_active  = true
  scheduling_saturday_from    = "09:00"
  scheduling_saturday_to      = "17:00"
  scheduling_sunday_active    = true
  scheduling_sunday_from      = "09:00"
  scheduling_sunday_to        = "17:00"
  scheduling_thursday_active  = true
  scheduling_thursday_from    = "09:00"
  scheduling_thursday_to      = "17:00"
  scheduling_tuesday_active   = true
  scheduling_tuesday_from     = "09:00"
  scheduling_tuesday_to       = "17:00"
  scheduling_wednesday_active = true
  scheduling_wednesday_from   = "09:00"
  scheduling_wednesday_to     = "17:00"
  vlan_tagging_settings       = "custom"
  vlan_tagging_vlan_id        = "1"
  force_delete                = true
}
