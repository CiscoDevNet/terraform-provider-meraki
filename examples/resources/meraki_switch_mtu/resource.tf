resource "meraki_switch_mtu" "example" {
  network_id = "L_123456"
  default_mtu_size = 9198
  overrides = [
    {
      mtu_size = 1500
      switches = ["Q234-ABCD-0001"]
    }
  ]
}
