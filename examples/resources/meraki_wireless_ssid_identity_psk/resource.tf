resource "meraki_wireless_ssid_identity_psk" "example" {
  network_id = "L_123456"
  number = "0"
  expires_at = "2018-02-11T00:00:00.090209Z"
  group_policy_id = "101"
  name = "Sample Identity PSK"
  passphrase = "Cisco123"
}
