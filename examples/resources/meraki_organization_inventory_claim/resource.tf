resource "meraki_organization_inventory_claim" "example" {
  organization_id = "123456"
  licenses = [
    {
      key  = "Z2XXXXXXXXXX"
      mode = "addDevices"
    }
  ]
  orders  = ["4CXXXXXXX"]
  serials = ["1234-1234-1234"]
}
