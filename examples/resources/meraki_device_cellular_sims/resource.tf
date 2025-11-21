resource "meraki_device_cellular_sims" "example" {
  serial = "1234-1234-1234"
  sim_failover_enabled = true
  sim_failover_timeout = 300
  sim_ordering = ["sim1"]
  sims = [
    {
      is_primary = false
      sim_order = 3
      slot = "sim1"
        apns = [
          {
          name = "internet"
          authentication_password = "secret"
          authentication_type = "pap"
          authentication_username = "milesmeraki"
          allowed_ip_types = ["ipv4"]
          }
        ]
    }
  ]
}
