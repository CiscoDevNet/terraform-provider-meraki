resource "meraki_network_alerts_settings" "example" {
  network_id = "L_123456"
  default_destinations_all_admins = true
  default_destinations_snmp = true
  default_destinations_emails = ["miles@meraki.com"]
  muting_by_port_schedules_enabled = true
  alerts = [
    {
      enabled = true
      type = "gatewayDown"
      alert_destinations_all_admins = false
      alert_destinations_snmp = false
      alert_destinations_emails = ["miles@meraki.com"]
    }
  ]
}
