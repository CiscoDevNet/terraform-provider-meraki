resource "meraki_organization" "example" {
  name = "My organization"
  management_details = [
    {
      name  = "MSP ID"
      value = "123456"
    }
  ]
}
