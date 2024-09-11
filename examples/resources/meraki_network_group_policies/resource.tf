resource "meraki_network_group_policies" "example" {
  network_id                  = "L_123456"
  name                        = "test_group_policy"
  bonjour_forwarding_settings = "custom"
  bonjour_forwarding_rules = [
    {
      description = "a simple bonjour rule"
      services    = ["All Services"]
      vlan_id     = "2"
    }
  ]
  bandwidth_bandwidth_limits_limit_down               = 1000000
  bandwidth_bandwidth_limits_limit_up                 = 1000000
  bandwidth_settings                                  = "custom"
  content_filtering_allowed_url_patterns_settings     = "network default"
  content_filtering_blocked_url_categories_categories = ["meraki:contentFiltering/category/1"]
  content_filtering_blocked_url_categories_settings   = "override"
  content_filtering_blocked_url_patterns_patterns     = ["http://www.betting.com"]
  content_filtering_blocked_url_patterns_settings     = "append"
}
