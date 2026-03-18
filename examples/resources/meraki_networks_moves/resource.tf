resource "meraki_networks_moves" "example" {
  organization_id         = "123456"
  simulate                = false
  network_id              = "N_569142402909112097"
  organizations_target_id = "146308"
}
