resource "meraki_wireless_alternate_management_interface" "example" {
  network_id = "L_123456"
  enabled = true
  vlan_id = 100
  access_points = [
    {
      alternate_management_ip = "1.2.3.4"
      dns1 = "8.8.8.8"
      dns2 = "8.8.4.4"
      gateway = "1.2.3.5"
      serial = "Q234-ABCD-5678"
      subnet_mask = "255.255.255.0"
    }
  ]
  protocols = ["radius"]
}
