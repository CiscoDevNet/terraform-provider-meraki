resource "meraki_sm_admin_role" "example" {
  organization_id = "123456"
  name = "sample name"
  scope = "all_tags"
  tags = ["tag"]
}
