resource "meraki_organization_policies_global_group_policy_adaptive_policy_group_assignment" "example" {
  organization_id          = "123456"
  policy_id                = "123"
  adaptive_policy_group_id = "1234"
}
