resource "meraki_organization_alerts_profile" "example" {
  organization_id              = "123456"
  description                  = "WAN 1 high utilization"
  type                         = "wanUtilization"
  alert_condition_bit_rate_bps = 10000
  alert_condition_duration     = 60
  alert_condition_interface    = "wan1"
  alert_condition_window       = 600
  recipients_emails            = ["admin@example.org"]
  network_tags                 = ["tag1"]
}
