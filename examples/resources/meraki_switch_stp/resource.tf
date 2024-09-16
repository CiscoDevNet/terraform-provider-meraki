resource "meraki_switch_stp" "example" {
  network_id   = "L_123456"
  rstp_enabled = true
  stp_bridge_priority = [
    {
      stp_priority    = 4096
      stacks          = ["789102"]
      switch_profiles = ["1098"]
      switches        = ["Q234-ABCD-0001"]
    }
  ]
}
