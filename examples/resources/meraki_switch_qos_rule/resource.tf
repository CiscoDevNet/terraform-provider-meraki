resource "meraki_switch_qos_rule" "example" {
  network_id     = "L_123456"
  dscp           = 0
  dst_port_range = "3000-3100"
  protocol       = "TCP"
  src_port       = 2000
  vlan           = 100
}
