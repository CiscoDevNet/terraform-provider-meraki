resource "meraki_network_syslog_servers" "example" {
  network_id = "L_123456"
  servers = [
    {
      host  = "1.2.3.4"
      port  = 443
      roles = ["Wireless Event log"]
    }
  ]
}
