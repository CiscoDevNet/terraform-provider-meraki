resource "meraki_organization_sase_site_attachment" "example" {
  organization_id = "123456"
  network_id      = "N_1234567890"
  region_slug     = "us-west-1"
}
