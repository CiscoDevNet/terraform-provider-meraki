resource "meraki_appliance_dns_local_record" "example" {
  organization_id = "123456"
  address = "10.1.1.0"
  hostname = "www.test.com"
  profile_id = "1"
}
