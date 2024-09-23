resource "meraki_wireless_ssid_hotspot_20" "example" {
  network_id          = "L_123456"
  number              = "0"
  enabled             = true
  network_access_type = "Private network"
  operator_name       = "Meraki Product Management"
  venue_name          = "SF Branch"
  venue_type          = "Unspecified Assembly"
  domains             = ["meraki.local"]
  mcc_mncs = [
    {
      mcc = "123"
      mnc = "456"
    }
  ]
  nai_realms = [
    {
      format = "1"
      realm  = ""
      methods = [
        {
          id                                                   = "1"
          authentication_types_non_eap_inner_authentication    = ["MSCHAPV2"]
          authentication_types_eap_inner_authentication        = ["EAP-TLS"]
          authentication_types_credentials                     = ["SIM"]
          authentication_types_tunneled_eap_method_credentials = ["USIM"]
        }
      ]
    }
  ]
  roam_consort_ois = ["ABC123"]
}
