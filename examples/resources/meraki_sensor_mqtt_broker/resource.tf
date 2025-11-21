resource "meraki_sensor_mqtt_broker" "example" {
  network_id = "L_123456"
  mqtt_broker_id = "123456"
  enabled = true
}
