resource "meraki_sm_target_group" "example" {
  network_id = "L_123456"
  name       = "Target group name"
  scope      = "withAny, tag1, tag2"
}
