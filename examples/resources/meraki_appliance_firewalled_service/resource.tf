resource "meraki_appliance_firewalled_service" "example" {
  network_id  = "L_123456"
  service     = "ICMP"
  access      = "restricted"
  allowed_ips = ["123.123.123.1"]
}
