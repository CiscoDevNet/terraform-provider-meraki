resource "meraki_devices" "example" {
  organization_id = "123456"
  items = [{
    serial = "1234-ABCD-1234"
    address = "1600 Pennsylvania Ave"
    lat = 37.4180951010362
    lng = -122.098531723022
    name = "My AP"
    notes = "My AP's note"
    tags = ["recently-added"]
  }]
}
