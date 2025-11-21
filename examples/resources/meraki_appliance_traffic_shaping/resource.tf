resource "meraki_appliance_traffic_shaping" "example" {
  network_id = "L_123456"
  global_bandwidth_limit_down = 5120
  global_bandwidth_limit_up = 2048
}
