resource "meraki_switch_stack" "example" {
  network_id = "L_123456"
  name = "A cool stack"
  serials = ["QBZY-XWVU-TSRQ"]
}
