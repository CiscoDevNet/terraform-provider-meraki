resource "meraki_sensor_mqtt_brokers" "example" {
  network_id = "L_123456"
  organization_id = "123456"
  items = [{
    mqtt_broker_id = "123456"
    enabled = true
  }]
}
