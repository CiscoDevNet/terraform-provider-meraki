resource "meraki_wireless_ethernet_port_profile_assign" "example" {
  network_id = "L_123456"
  profile_id = "1001"
  serials = ["Q234-ABCD-0001"]
}
