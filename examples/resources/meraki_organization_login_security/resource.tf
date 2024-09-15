resource "meraki_organization_login_security" "example" {
  organization_id                                     = "123456"
  account_lockout_attempts                            = 3
  enforce_account_lockout                             = true
  enforce_different_passwords                         = true
  enforce_idle_timeout                                = true
  enforce_login_ip_ranges                             = true
  enforce_password_expiration                         = true
  enforce_strong_passwords                            = true
  enforce_two_factor_auth                             = true
  idle_timeout_minutes                                = 30
  num_different_passwords                             = 3
  password_expiration_days                            = 90
  api_authentication_ip_restrictions_for_keys_enabled = true
  api_authentication_ip_restrictions_for_keys_ranges  = ["192.195.83.1"]
  login_ip_ranges                                     = ["192.195.83.1"]
}
