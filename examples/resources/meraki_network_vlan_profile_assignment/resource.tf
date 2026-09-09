resource "meraki_network_vlan_profile_assignment" "example" {
  network_id         = "L_123456"
  vlan_profile_iname = "Profile1"
  serials            = ["Q234-ABCD-5678"]
}
