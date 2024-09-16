resource "meraki_network_vlan_profile" "example" {
  network_id = "L_123456"
  iname      = "Profile1"
  name       = "My VLAN profile name"
  vlan_groups = [
    {
      name     = "named-group-1"
      vlan_ids = "2,5-7"
    }
  ]
  vlan_names = [
    {
      name                     = "named-1"
      vlan_id                  = "1"
      adaptive_policy_group_id = "791"
    }
  ]
}
