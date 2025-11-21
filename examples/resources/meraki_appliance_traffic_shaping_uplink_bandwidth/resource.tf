resource "meraki_appliance_traffic_shaping_uplink_bandwidth" "example" {
  network_id = "L_123456"
  cellular_limit_down = 100000
  cellular_limit_up = 100000
  wan1_limit_down = 100000
  wan1_limit_up = 100000
  wan2_limit_down = 100000
  wan2_limit_up = 100000
}
