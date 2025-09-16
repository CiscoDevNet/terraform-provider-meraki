resource "meraki_network_meraki_auth_users" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    account_type           = "802.1X"
    email                  = "miles321@meraki.com"
    email_password_to_user = false
    is_admin               = false
    name                   = "Miles Meraki"
    password               = "Cisco123456&"
    authorizations = [
      {
        expires_at  = "2018-03-13T00:00:00.090210Z"
        ssid_number = 1
      }
    ]
  }]
}
