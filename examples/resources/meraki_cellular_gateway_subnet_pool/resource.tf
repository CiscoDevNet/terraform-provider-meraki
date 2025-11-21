resource "meraki_cellular_gateway_subnet_pool" "example" {
  network_id = "L_123456"
  cidr = "192.168.0.0/24"
  mask = 26
}
