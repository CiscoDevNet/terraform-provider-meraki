resource "meraki_camera_custom_analytics" "example" {
  serial      = "1234-ABCD-1234"
  artifact_id = "1"
  enabled     = true
  parameters = [
    {
      name  = "detection_threshold"
      value = "0.5"
    }
  ]
}
