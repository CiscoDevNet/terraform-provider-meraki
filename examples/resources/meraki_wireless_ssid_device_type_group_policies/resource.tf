resource "meraki_wireless_ssid_device_type_group_policies" "example" {
  network_id = "L_123456"
  number = "0"
  enabled = true
  device_type_policies = [
    {
      device_policy = "Allowed"
      device_type = "Android"
    }
  ]
}
