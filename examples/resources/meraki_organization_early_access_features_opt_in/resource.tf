resource "meraki_organization_early_access_features_opt_in" "example" {
  organization_id         = "123456"
  short_name              = "has_vlan_db"
  limit_scope_to_networks = ["N_12345"]
}
