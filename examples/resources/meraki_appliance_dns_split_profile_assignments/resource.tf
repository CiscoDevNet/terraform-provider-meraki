resource "meraki_appliance_dns_split_profile_assignments" "example" {
  organization_id = "123456"
  items = [
    {
      network_id = "N_123456"
      profile_id = "1234"
    }
  ]
}
