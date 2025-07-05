resource "meraki_sm_admin_roles" "example" {
  organization_id = "123456"
  items = [{
    name  = "sample name"
    scope = "all_tags"
    tags  = ["tag"]
  }]
}
