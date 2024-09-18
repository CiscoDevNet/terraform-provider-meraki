resource "meraki_organization_adaptive_policy_group" "example" {
  organization_id = "123456"
  description     = "Group of XYZ Corp Employees"
  name            = "Employee Group"
  sgt             = 1000
  policy_objects = [
    {
      id   = "2345"
      name = "Example Policy Object"
    }
  ]
}
