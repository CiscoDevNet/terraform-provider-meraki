resource "meraki_organization_config_template" "example" {
  organization_id = "123456"
  copy_from_network_id = "N_24329156"
  name = "My config template"
  time_zone = "America/Los_Angeles"
}
