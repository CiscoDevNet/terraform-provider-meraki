resource "meraki_network_webhook_payload_templates" "example" {
  network_id      = "L_123456"
  organization_id = "123456"
  items = [{
    body = "{\"event_type\":\"{{alertTypeId}}\",\"client_payload\":{\"text\":\"{{alertData}}\"}}"
    name = "Custom Template New"
    headers = [
      {
        name     = "Authorization"
        template = "Bearer {{sharedSecret}}"
      }
    ]
  }]
}
