resource "meraki_network_snmp" "example" {
  network_id       = "L_123456"
  access           = "users"
  community_string = "sample"
  users = [
    {
      passphrase = "hunter2"
      username   = "AzureDiamond"
    }
  ]
}
