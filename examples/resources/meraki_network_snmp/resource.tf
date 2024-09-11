resource "meraki_network_snmp" "example" {
  network_id       = "L_123456"
  access           = "users"
  community_string = "MerakiCommunity"
  users = [
    {
      username   = "User1"
      passphrase = "Password123"
    }
  ]
}
