resource "meraki_organization_integrations_xdr_networks" "example" {
  organization_id = "123456"
  networks = [
    {
      network_id    = "N_1234567"
      product_types = ["appliance"]
    }
  ]
}
