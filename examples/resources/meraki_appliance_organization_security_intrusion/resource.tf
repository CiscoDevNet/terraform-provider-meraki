resource "meraki_appliance_organization_security_intrusion" "example" {
  organization_id = "123456"
  allowed_rules = [
    {
      message = "SQL sa login failed"
      rule_id = "meraki:intrusion/snort/GID/01/SID/688"
    }
  ]
}
