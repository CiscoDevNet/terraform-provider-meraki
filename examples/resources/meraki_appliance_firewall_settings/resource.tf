resource "meraki_appliance_firewall_settings" "example" {
  network_id                               = "L_123456"
  spoofing_protection_ip_source_guard_mode = "block"
}
