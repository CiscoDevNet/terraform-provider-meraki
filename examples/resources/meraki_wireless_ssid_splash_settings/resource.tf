resource "meraki_wireless_ssid_splash_settings" "example" {
  network_id = "L_123456"
  number = "0"
  allow_simultaneous_logins = false
  block_all_traffic_before_sign_on = false
  controller_disconnection_behavior = "default"
  redirect_url = "https://example.com"
  splash_timeout = 1440
  use_redirect_url = true
  use_splash_url = false
  welcome_message = "Welcome!"
  billing_prepaid_access_fast_login_enabled = false
  billing_reply_to_email_address = "user@email.com"
  billing_free_access_duration_in_minutes = 20
  billing_free_access_enabled = false
  guest_sponsorship_duration_in_minutes = 30
  guest_sponsorship_guest_can_request_timeframe = false
  self_registration_authorization_type = "admin"
  self_registration_enabled = true
}
