resource "meraki_organization_wireless_devices_provisioning_deployments" "example" {
  organization_id = "123456"
  items = [{
    type               = "deploy"
    status             = "ready"
    network_id         = "N_24329156"
    devices_new_serial = "Q234-ABCD-5678"
    devices_new_tags   = ["tag1"]
  }]
}
