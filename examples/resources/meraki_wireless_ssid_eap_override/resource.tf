resource "meraki_wireless_ssid_eap_override" "example" {
  network_id = "L_123456"
  number = "0"
  max_retries = 5
  timeout = 5
  eapol_key_retries = 4
  eapol_key_timeout_in_ms = 5000
  identity_retries = 5
  identity_timeout = 5
}
