resource "meraki_wireless_radio_settings" "example" {
  serial                             = "1234-1234-1234"
  five_ghz_settings_channel          = 64
  five_ghz_settings_channel_width    = 20
  two_four_ghz_settings_channel      = 11
  two_four_ghz_settings_target_power = 14
}
