resource "meraki_appliance_port" "example" {
  network_id = "L_123456"
  port_id = "12"
  access_policy = "open"
  drop_untagged_traffic = false
  enabled = true
  type = "access"
  vlan = 1
}
