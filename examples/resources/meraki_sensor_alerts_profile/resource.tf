resource "meraki_sensor_alerts_profile" "example" {
  network_id             = "L_123456"
  name                   = "My Sensor Alert Profile"
  recipients_emails      = ["miles@meraki.com"]
  recipients_sms_numbers = ["+15555555555"]
  conditions = [
    {
      direction                     = "above"
      duration                      = 60
      metric                        = "temperature"
      threshold_temperature_celsius = 20.5
    }
  ]
  serials = ["Q234-ABCD-0001"]
}
