resource "meraki_network" "example" {
  organization_id = "123456"
  name            = "Main Office"
  notes           = "Additional description of the network"
  time_zone       = "America/Los_Angeles"
  product_types   = ["switch"]
  tags            = ["tag1"]
}
