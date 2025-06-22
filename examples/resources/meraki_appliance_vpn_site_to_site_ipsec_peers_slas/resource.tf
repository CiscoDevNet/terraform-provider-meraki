resource "meraki_appliance_vpn_site_to_site_ipsec_peers_slas" "example" {
  organization_id = "123456"
  items = [
    {
      name = "sla policy"
      uri  = "http://checkthisendpoint.com"
    }
  ]
}
