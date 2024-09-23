resource "meraki_wireless_ssid_splash_settings" "example" {
  network_id                                    = "L_123456"
  number                                        = "0"
  allow_simultaneous_logins                     = false
  block_all_traffic_before_sign_on              = false
  controller_disconnection_behavior             = "default"
  redirect_url                                  = "https://example.com"
  splash_timeout                                = 1440
  use_redirect_url                              = true
  welcome_message                               = "Welcome!"
  guest_sponsorship_duration_in_minutes         = 30
  guest_sponsorship_guest_can_request_timeframe = false
}
