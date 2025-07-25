resource "meraki_appliance_ssids" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    number              = "1"
    auth_mode           = "8021x-radius"
    enabled             = true
    name                = "My SSID"
    visible             = true
    wpa_encryption_mode = "WPA2 only"
    radius_servers = [
      {
        host   = "0.0.0.0"
        port   = 1000
        secret = "secret"
      }
    ]
  }]
}
