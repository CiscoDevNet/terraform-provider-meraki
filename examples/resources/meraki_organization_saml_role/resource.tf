resource "meraki_organization_saml_role" "example" {
  organization_id = "123456"
  org_access      = "none"
  role            = "myrole"
  networks = [
    {
      access = "full"
      id     = "N_24329156"
    }
  ]
  tags = [
    {
      access = "read-only"
      tag    = "west"
    }
  ]
}
