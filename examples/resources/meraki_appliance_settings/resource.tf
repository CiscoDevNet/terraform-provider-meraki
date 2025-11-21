resource "meraki_appliance_settings" "example" {
  network_id = "L_123456"
  client_tracking_method = "MAC address"
  deployment_mode = "routed"
  dynamic_dns_enabled = true
  dynamic_dns_prefix = "test"
}
