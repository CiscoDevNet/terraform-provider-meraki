resource "meraki_organization_extensions_thousand_eyes_network" "example" {
  organization_id = "123456"
  enabled = true
  network_id = "N_123456"
  tests = [
    {
      network_id = "N_123456"
      template_id = "123abc"
      template_user_inputs_tenant = "cisco"
    }
  ]
}
