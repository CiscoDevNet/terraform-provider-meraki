data "meraki_organization" "org1" {
  name = "Org1"
}

resource "meraki_network" "tf_1" {
  organization_id = data.meraki_organization.org1.id
  name            = "TF-1"
  notes           = "My first network created by Terraform"
  product_types   = ["switch", "wireless"]
}

resource "meraki_network_device_claim" "switch" {
  network_id = meraki_network.tf_1.id
  serials    = ["1234-1234-1234"]
}

resource "meraki_device" "switch" {
  serial  = tolist(meraki_network_device_claim.switch.serials)[0]
  address = "1600 Pennsylvania Ave"
  name    = "TF-Switch-1"
  notes   = "My first switch configured by Terraform"
}

resource "meraki_switch_mtu" "mtu" {
  network_id       = meraki_network.tf_1.id
  default_mtu_size = 9100
}

resource "meraki_switch_port" "port_5_10" {
  count   = 6
  serial  = tolist(meraki_network_device_claim.switch.serials)[0]
  port_id = count.index + 5
  enabled = true
  name    = "My first switch ports configured by Terraform"
  type    = "access"
  vlan    = 10
}
