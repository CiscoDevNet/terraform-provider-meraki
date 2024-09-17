resource "meraki_switch_warm_spare" "example" {
  serial       = "1234-ABCD-1234"
  enabled      = false
  spare_serial = "Q234-ABCD-0002"
}
