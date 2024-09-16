resource "meraki_switch_routing_ospf" "example" {
  network_id                        = "L_123456"
  dead_timer_in_seconds             = 40
  enabled                           = true
  hello_timer_in_seconds            = 10
  md5_authentication_enabled        = true
  md5_authentication_key_id         = 1
  md5_authentication_key_passphrase = "abc1234"
  v3_dead_timer_in_seconds          = 40
  v3_enabled                        = true
  v3_hello_timer_in_seconds         = 10
  v3_areas = [
    {
      area_id   = "0"
      area_name = "V3 Backbone"
      area_type = "normal"
    }
  ]
  areas = [
    {
      area_id   = "0"
      area_name = "Backbone"
      area_type = "normal"
    }
  ]
}
