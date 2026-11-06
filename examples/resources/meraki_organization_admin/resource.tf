resource "meraki_organization_admin" "example" {
  organization_id       = "123456"
  authentication_method = "Email"
  email                 = "miles@meraki.com"
  name                  = "Miles Meraki"
  org_access            = "none"
  networks = [
    {
      access                 = "full"
      id                     = "N_24329156"
      packet_capture_allowed = true
      switch_port_tags       = ["meraki-test"]
    }
  ]
  tags = [
    {
      access = "read-only"
      tag    = "west"
    }
  ]
}
