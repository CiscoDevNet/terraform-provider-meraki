resource "meraki_switch_routing_multicast" "example" {
  network_id                                               = "L_123456"
  default_settings_flood_unknown_multicast_traffic_enabled = true
  default_settings_igmp_snooping_enabled                   = true
  overrides = [
    {
      flood_unknown_multicast_traffic_enabled = true
      igmp_snooping_enabled                   = true
      switches                                = ["Q234-ABCD-0001"]
    }
  ]
}
