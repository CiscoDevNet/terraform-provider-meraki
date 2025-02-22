resource "meraki_camera_quality_and_retention" "example" {
  serial                            = "1234-ABCD-1234"
  audio_recording_enabled           = false
  motion_based_retention_enabled    = false
  motion_detector_version           = 2
  quality                           = "Standard"
  resolution                        = "1280x720"
  restricted_bandwidth_mode_enabled = false
}
