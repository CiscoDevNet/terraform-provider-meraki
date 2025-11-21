resource "meraki_network_netflow" "example" {
  network_id = "L_123456"
  collector_ip = "1.2.3.4"
  collector_port = 1234
  reporting_enabled = true
}
