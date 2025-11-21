resource "meraki_insight_monitored_media_server" "example" {
  organization_id = "123456"
  address = "123.123.123.1"
  best_effort_monitoring_enabled = true
  name = "Sample VoIP Provider"
}
