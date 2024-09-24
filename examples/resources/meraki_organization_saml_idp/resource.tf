resource "meraki_organization_saml_idp" "example" {
  organization_id           = "123456"
  slo_logout_url            = "https://somewhere.com"
  x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:AA"
}