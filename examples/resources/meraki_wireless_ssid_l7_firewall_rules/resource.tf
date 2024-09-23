resource "meraki_wireless_ssid_l7_firewall_rules" "example" {
  network_id = "L_123456"
  number     = "0"
  rules = [
    {
      policy = "deny"
      type   = "host"
      value  = "google.com"
    }
  ]
}
