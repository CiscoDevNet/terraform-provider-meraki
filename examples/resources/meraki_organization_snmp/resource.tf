resource "meraki_organization_snmp" "example" {
  organization_id = "123"
  v2c_enabled     = false
  v3_auth_mode    = "SHA"
  v3_auth_pass    = "password"
  v3_enabled      = true
  v3_priv_mode    = "AES128"
  v3_priv_pass    = "password"
  peer_ips        = ["123.123.123.1"]
}
