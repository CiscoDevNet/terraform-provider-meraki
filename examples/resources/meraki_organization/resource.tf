resource "meraki_organization" "example" {
  name = "Dev"
  management_details = [
    {
      name  = "MSP ID"
      value = "123456"
    }
  ]
}
