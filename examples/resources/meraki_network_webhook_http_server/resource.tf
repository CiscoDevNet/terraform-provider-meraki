resource "meraki_network_webhook_http_server" "example" {
  network_id                           = "L_123456"
  name                                 = "Example Webhook Server"
  shared_secret                        = "shhh"
  url                                  = "https://example.com"
  payload_template_name                = "Meraki (included)"
  payload_template_payload_template_id = "wpt_00001"
}
