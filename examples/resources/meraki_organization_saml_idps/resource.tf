resource "meraki_organization_saml_idps" "example" {
  organization_id = "123456"
  items = [{
    slo_logout_url            = "https://somewhere.com"
    sso_login_url             = "https://onelogin.com/trust/saml2/http-post/sso/3de5f942-e7b8-4cb9-94e3-85828111158b"
    x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:CA"
  }]
}
