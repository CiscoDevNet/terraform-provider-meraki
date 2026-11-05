resource "meraki_organization_policies_global_group_policy_firewall_ruleset_assignment" "example" {
  organization_id = "123456"
  policy_id       = "789"
  priority        = 1
  ruleset_id      = "456"
}
