resource "meraki_network_snmp" "example" {
  network_id = "L_123456"
  access     = "users"
  users = [
    {
      passphrase = "N7!qW4#cT3@p"
      username   = "AzureDiamond"
    }
  ]
}
