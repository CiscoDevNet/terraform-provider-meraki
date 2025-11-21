resource "meraki_organization_licenses" "example" {
  organization_id = "123456"
  items = [{
    license_id = "123"
    device_serial = "Q234-ABCD-5678"
  }]
}
