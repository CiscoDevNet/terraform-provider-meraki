resource "meraki_device_management_interface" "example" {
  serial                  = "1234-1234-1234"
  wan1_static_gateway_ip  = "1.2.3.1"
  wan1_static_ip          = "1.2.3.4"
  wan1_static_subnet_mask = "255.255.255.0"
  wan1_using_static_ip    = true
  wan1_vlan               = 7
  wan1_static_dns         = ["1.2.3.2"]
}
