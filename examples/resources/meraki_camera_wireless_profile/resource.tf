resource "meraki_camera_wireless_profile" "example" {
  network_id = "L_123456"
  name = "wireless profile A"
  identity_password = "password123"
  identity_username = "identityname"
  ssid_auth_mode = "8021x-radius"
  ssid_encryption_mode = "wpa-eap"
  ssid_name = "ssid test"
}
