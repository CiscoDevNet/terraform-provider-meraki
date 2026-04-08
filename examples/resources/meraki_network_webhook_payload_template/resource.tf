resource "meraki_network_webhook_payload_template" "example" {
  network_id = "L_123456"
  body       = "{\"event_type\":\"{{alertTypeId}}\",\"client_payload\":{\"text\":\"{{alertData}}\"}}"
  name       = "Custom 1"
  headers = [
    {
      name     = "Authorization"
      template = "Bearer {{sharedSecret}}"
    }
  ]
}
