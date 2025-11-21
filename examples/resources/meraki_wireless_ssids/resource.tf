resource "meraki_wireless_ssids" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    number = "0"
    auth_mode = "psk"
    available_on_all_aps = false
    band_selection = "5 GHz band only"
    enabled = false
    encryption_mode = "wpa"
    ip_assignment_mode = "Bridge mode"
    lan_isolation_enabled = false
    mandatory_dhcp_enabled = false
    min_bitrate = 5.5
    name = "My SSID"
    per_client_bandwidth_limit_down = 0
    per_client_bandwidth_limit_up = 0
    per_ssid_bandwidth_limit_down = 0
    per_ssid_bandwidth_limit_up = 0
    psk = "deadbeef"
    splash_page = "Click-through splash page"
    use_vlan_tagging = false
    visible = false
    walled_garden_enabled = false
    wpa_encryption_mode = "WPA2 only"
    dot11r_adaptive = false
    dot11r_enabled = false
    dot11w_enabled = false
    dot11w_required = false
    speed_burst_enabled = false
    availability_tags = ["tag1"]
  }]
}
