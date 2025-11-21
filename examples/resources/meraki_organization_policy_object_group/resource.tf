resource "meraki_organization_policy_object_group" "example" {
  organization_id = "123456"
  category = "NetworkObjectGroup"
  name = "Web Servers - Datacenter 10"
  object_ids = [100]
}
