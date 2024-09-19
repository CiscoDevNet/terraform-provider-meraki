resource "meraki_wireless_settings" "example" {
  network_id                                = "L_123456"
  ipv6_bridge_enabled                       = true
  led_lights_on                             = true
  location_analytics_enabled                = false
  meshing_enabled                           = true
  upgrade_strategy                          = "minimizeUpgradeTime"
  named_vlans_pool_dhcp_monitoring_duration = 5
  named_vlans_pool_dhcp_monitoring_enabled  = false
}
