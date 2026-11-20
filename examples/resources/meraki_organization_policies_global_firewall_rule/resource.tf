resource "meraki_organization_policies_global_firewall_rule" "example" {
  organization_id                      = "123456"
  description                          = "This is rule 1"
  enabled                              = true
  name                                 = "Allow developers"
  policy                               = "deny"
  priority                             = 1
  ruleset_id                           = "32"
  destinations_criteria_address_ranges = ["1.1.1.1"]
  destinations_match_criteria          = ["addressRanges"]
  sources_criteria_address_ranges      = ["1.1.1.1"]
  sources_match_criteria               = ["addressRanges"]
}
