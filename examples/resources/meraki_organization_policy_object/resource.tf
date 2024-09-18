resource "meraki_organization_policy_object" "example" {
  organization_id = "123456"
  category        = "network"
  cidr            = "10.0.0.0/24"
  name            = "Web Servers - Datacenter 10"
  type            = "cidr"
}
