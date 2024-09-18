resource "meraki_organization_adaptive_policy_settings" "example" {
  organization_id  = "123456"
  enabled_networks = ["L_11111111"]
}
