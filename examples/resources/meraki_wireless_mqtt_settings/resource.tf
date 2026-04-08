resource "meraki_wireless_mqtt_settings" "example" {
  organization_id           = "123456"
  ble_enabled               = false
  ble_type                  = "iBeacon"
  ble_flush_frequency       = 60
  ble_hysteresis_enabled    = true
  ble_hysteresis_threshold  = 1
  mqtt_enabled              = true
  mqtt_topic                = "Test Topic"
  mqtt_broker_name          = "My Broker"
  mqtt_publishing_frequency = 1
  mqtt_publishing_qos       = 1
  mqtt_message_fields       = ["RSSI"]
  network_id                = "L_1234"
  wifi_enabled              = false
  wifi_type                 = "Associated"
  wifi_flush_frequency      = 60
  wifi_hysteresis_enabled   = false
  wifi_hysteresis_threshold = 1
}
