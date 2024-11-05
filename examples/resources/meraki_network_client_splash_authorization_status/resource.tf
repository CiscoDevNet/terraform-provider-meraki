resource "meraki_network_client_splash_authorization_status" "example" {
  network_id            = "L_123456"
  client_id             = "1.2.3.4"
  ssids_0_is_authorized = true
  ssids_2_is_authorized = false
}
