resource "meraki_wireless_ssid_firewall_isolation_allowlist_entry" "example" {
  organization_id = "123456"
  description     = "Example mac address"
  client_mac      = "A1:B2:C3:D4:E5:F6"
  network_id      = "N_123"
  ssid_number     = 2
}
