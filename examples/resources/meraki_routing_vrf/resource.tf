resource "meraki_routing_vrf" "example" {
  organization_id     = "123456"
  description         = "Virtual network for Engineering"
  name                = "Engineering"
  route_distinguisher = "65000:42"
  route_target        = "65000:42"
  appliance_enabled   = true
  appliance_index     = "1"
}
