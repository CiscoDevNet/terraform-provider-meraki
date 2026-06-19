resource "meraki_network_firmware_upgrades_rollback" "example" {
  network_id    = "L_123456"
  product       = "switch"
  time          = "2020-10-21T02:00:00Z"
  to_version_id = "7857"
  reasons = [
    {
      category = "performance"
      comment  = "Network was slower with the upgrade"
    }
  ]
}
