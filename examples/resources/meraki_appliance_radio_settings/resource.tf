resource "meraki_appliance_radio_settings" "example" {
  serial = "1234-ABCD-1234"
  five_ghz_settings_channel = 40
  five_ghz_settings_channel_width = 20
  five_ghz_settings_target_power = 15
  two_four_ghz_settings_channel = 11
  two_four_ghz_settings_target_power = 5
}
