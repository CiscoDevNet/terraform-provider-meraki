resource "meraki_wireless_air_marshal_rule" "example" {
  network_id   = "L_123456"
  type         = "allow"
  match_string = "00:11:22:33:44:55"
  match_type   = "bssid"
}
