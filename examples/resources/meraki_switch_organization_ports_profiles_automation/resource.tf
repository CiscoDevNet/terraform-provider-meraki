resource "meraki_switch_organization_ports_profiles_automation" "example" {
  organization_id = "123456"
  description     = "A full length description of the automation."
  name            = "Automation 1"
  rules = [
    {
      priority     = 1
      profile_id   = "32"
      profile_name = "Profile 2"
      conditions = [
        {
          attribute = "LLDP system description"
          values    = ["Meraki MR*"]
        }
      ]
    }
  ]
}
