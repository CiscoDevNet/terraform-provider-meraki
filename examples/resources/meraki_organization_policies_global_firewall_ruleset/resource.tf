resource "meraki_organization_policies_global_firewall_ruleset" "example" {
  organization_id = "123456"
  description     = "Block all Social Media"
  name            = "Block Social Media"
}
