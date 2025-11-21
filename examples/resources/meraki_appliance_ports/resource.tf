resource "meraki_appliance_ports" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    port_id = "12"
    access_policy = "open"
    drop_untagged_traffic = false
    enabled = true
    type = "access"
    vlan = 1
  }]
}
