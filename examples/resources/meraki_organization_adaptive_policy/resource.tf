resource "meraki_organization_adaptive_policy" "example" {
  organization_id        = "123456"
  last_entry_rule        = "allow"
  destination_group_id   = "333"
  destination_group_name = "IoT Servers"
  destination_group_sgt  = 51
  source_group_id        = "222"
  source_group_name      = "IoT Devices"
  source_group_sgt       = 50
  acls = [
    {
      id   = "444"
      name = "Block web"
    }
  ]
}
