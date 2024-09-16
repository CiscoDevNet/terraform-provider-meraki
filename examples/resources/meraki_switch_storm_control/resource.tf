resource "meraki_switch_storm_control" "example" {
  network_id                = "L_123456"
  broadcast_threshold       = 30
  multicast_threshold       = 30
  unknown_unicast_threshold = 30
}
