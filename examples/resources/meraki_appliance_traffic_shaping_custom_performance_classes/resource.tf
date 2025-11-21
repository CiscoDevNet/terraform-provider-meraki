resource "meraki_appliance_traffic_shaping_custom_performance_classes" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    max_jitter = 100
    max_latency = 100
    max_loss_percentage = 5
    name = "myCustomPerformanceClass"
  }]
}
