resource "meraki_network_device_claim" "example" {
  network_id = "123456"
  details_by_device = [
    {
      serial = "Q234-ABCD-5678"
      details = [
        {
          name  = "username"
          value = "milesmeraki"
        }
      ]
    }
  ]
  serials = ["1234-1234-1234"]
}
