resource "meraki_appliance_warm_spare" "example" {
  network_id = "L_123456"
  enabled = true
  spare_serial = "Q234-ABCD-5678"
  uplink_mode = "virtual"
  virtual_ip1 = "1.2.3.4"
  virtual_ip2 = "2.3.4.5"
}
