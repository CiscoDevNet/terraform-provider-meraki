resource "meraki_appliance_dns_split_profile" "example" {
  organization_id       = "123456"
  name                  = "Default profile"
  nameservers_addresses = ["12.1.10.1"]
  hostnames             = ["*.test1.com"]
}
