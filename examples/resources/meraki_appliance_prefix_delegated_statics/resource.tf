resource "meraki_appliance_prefix_delegated_statics" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    description = "Prefix on WAN 1 of Long Island Office network"
    prefix = "2002:db8:3c4d:15::/64"
    origin_type = "internet"
    origin_interfaces = ["wan1"]
  }]
}
