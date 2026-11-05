resource "meraki_organization_policies_global_group_policy" "example" {
  organization_id = "123456"
  description     = "My policy description"
  name            = "My policy"
}
