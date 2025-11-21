resource "meraki_network_client_policy" "example" {
  network_id = "L_123456"
  client_id = "1.2.3.4"
  device_policy = "Group policy"
  group_policy_id = "101"
}
