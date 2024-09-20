resource "meraki_wireless_ssid_bonjour_forwarding" "example" {
  network_id = "L_123456"
  number     = "0"
  enabled    = true
  rules = [
    {
      description = "A simple bonjour rule"
      vlan_id     = "1"
      services    = ["All Services"]
    }
  ]
}
