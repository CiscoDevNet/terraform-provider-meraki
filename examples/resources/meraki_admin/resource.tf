resource "meraki_admin" "example" {
  organization_id = "123456"
  email           = "miles@meraki.com"
  name            = "Miles Meraki"
  org_access      = "none"
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
