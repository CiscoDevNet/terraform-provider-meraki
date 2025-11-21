resource "meraki_appliance_rf_profiles" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    name = "MX RF Profile"
    five_ghz_settings_ax_enabled = true
    five_ghz_settings_min_bitrate = 48
    per_ssid_settings_1_band_operation_mode = "dual"
    per_ssid_settings_1_band_steering_enabled = true
    per_ssid_settings_2_band_operation_mode = "dual"
    per_ssid_settings_2_band_steering_enabled = true
    per_ssid_settings_3_band_operation_mode = "dual"
    per_ssid_settings_3_band_steering_enabled = true
    per_ssid_settings_4_band_operation_mode = "dual"
    per_ssid_settings_4_band_steering_enabled = true
    two_four_ghz_settings_ax_enabled = true
    two_four_ghz_settings_min_bitrate = 12
  }]
}
