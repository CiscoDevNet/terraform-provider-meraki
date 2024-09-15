resource "meraki_switch_dscp_to_cos_mappings" "example" {
  network_id = "L_123456"
  mappings = [
    {
      cos   = 1
      dscp  = 1
      title = "Video"
    }
  ]
}
