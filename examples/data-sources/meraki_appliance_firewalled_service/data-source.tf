data "meraki_appliance_firewalled_service" "example" {
  network_id = "L_123456"
  service = "ICMP"
}
