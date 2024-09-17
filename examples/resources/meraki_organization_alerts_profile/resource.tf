resource "meraki_organization_alerts_profile" "example" {
  organization_id              = "123456"
  description                  = "WAN 1 high utilization"
  type                         = "wanUtilization"
  alert_condition_bit_rate_bps = 10000
  alert_condition_duration     = 60
  alert_condition_interface    = "wan1"
  alert_condition_jitter_ms    = 100
  alert_condition_latency_ms   = 100
  alert_condition_loss_ratio   = 0.1
  alert_condition_mos          = 3.5
  alert_condition_window       = 600
  recipients_emails            = ["admin@example.org"]
  recipients_http_server_ids   = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vcGF0aA=="]
  network_tags                 = ["tag1"]
}
