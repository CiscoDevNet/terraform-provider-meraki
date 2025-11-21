resource "meraki_appliance_traffic_shaping_custom_performance_class" "example" {
  network_id = "L_123456"
  max_jitter = 100
  max_latency = 100
  max_loss_percentage = 5
  name = "myCustomPerformanceClass"
}
