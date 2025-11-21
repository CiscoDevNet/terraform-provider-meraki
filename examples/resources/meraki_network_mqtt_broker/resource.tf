resource "meraki_network_mqtt_broker" "example" {
  network_id = "L_123456"
  host = "1.2.3.4"
  name = "MQTT_Broker_1"
  port = 443
  authentication_password = "*****"
  authentication_username = "milesmeraki"
  security_mode = "tls"
  security_tls_ca_certificate = "*****"
  security_tls_verify_hostnames = true
}
