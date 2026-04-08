resource "meraki_organization_policy_object_groups" "example" {
  organization_id = "123456"
  items = [{
    category   = "NetworkObjectGroup"
    name       = "Web Servers Group"
    object_ids = [100]
  }]
}
