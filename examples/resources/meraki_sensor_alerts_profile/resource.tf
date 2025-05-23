resource "meraki_sensor_alerts_profile" "example" {
  network_id         = "L_123456"
  include_sensor_url = true
  message            = "Check with Miles on what to do."
  name               = "My Sensor Alert Profile"
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
