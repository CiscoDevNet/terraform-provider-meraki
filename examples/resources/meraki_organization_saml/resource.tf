resource "meraki_organization_saml" "example" {
  organization_id        = "123456"
  enabled                = true
  sp_initiated_idp_id    = "uu3H_bx28Nnd"
  sp_initiated_subdomain = "example_subdomain"
}
