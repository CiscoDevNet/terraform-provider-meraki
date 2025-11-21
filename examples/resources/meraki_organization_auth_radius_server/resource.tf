resource "meraki_organization_auth_radius_server" "example" {
  organization_id = "123456"
  address = "1.2.3.4"
  name = "HQ RADIUS server"
  secret = "secret"
  modes = [
    {
      mode = "auth"
      port = 1812
    }
  ]
}
