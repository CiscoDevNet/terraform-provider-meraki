resource "meraki_cellular_gateway_uplink" "example" {
  network_id                  = "L_123456"
  bandwidth_limits_limit_down = 1000000
  bandwidth_limits_limit_up   = 1000000
}
