resource "meraki_cellular_gateway_connectivity_monitoring_destinations" "example" {
  network_id = "L_123456"
  destinations = [
    {
      default = true
      description = "Google"
      ip = "1.2.3.4"
    }
  ]
}
