resource "meraki_camera_sense" "example" {
  serial                  = "1234-ABCD-1234"
  sense_enabled           = true
  audio_detection_enabled = false
}
