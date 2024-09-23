resource "meraki_wireless_ssid_vpn" "example" {
  network_id                  = "L_123456"
  number                      = "0"
  concentrator_network_id     = "N_123"
  concentrator_vlan_id        = 44
  failover_heartbeat_interval = 10
  failover_idle_timeout       = 30
  failover_request_ip         = "1.1.1.1"
  split_tunnel_enabled        = true
  split_tunnel_rules = [
    {
      comment   = "split tunnel rule 1"
      dest_cidr = "1.1.1.1/32"
      dest_port = "any"
      policy    = "allow"
      protocol  = "Any"
    }
  ]
}
