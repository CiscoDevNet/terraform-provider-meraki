resource "meraki_wireless_location_scanning_receiver" "example" {
  organization_id = "123456"
  shared_secret   = "mysecretvalue"
  url             = "https://www.myreceiver.com"
  version         = "3"
  network_id      = "L_1234"
  radio_type      = "Wi-Fi"
}
