resource "meraki_organization_license" "example" {
  organization_id = "123456"
  license_id      = "123"
  device_serial   = "Q234-ABCD-5678"
}
