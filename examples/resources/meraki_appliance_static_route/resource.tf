resource "meraki_appliance_static_route" "example" {
  network_id = "L_123456"
  gateway_ip = "192.168.128.254"
  name       = "My route"
  subnet     = "5.5.5.0/24"
}
