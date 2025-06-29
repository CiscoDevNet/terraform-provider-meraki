resource "meraki_organization_policy_objects" "example" {
  organization_id = "123456"
  items = [{
    category = "network"
    cidr     = "10.0.0.0/24"
    name     = "Web Servers - Datacenter 10"
    type     = "cidr"
  }]
}
