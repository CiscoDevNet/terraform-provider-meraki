resource "meraki_appliance_network_security_intrusion" "example" {
  network_id                       = "L_123456"
  ids_rulesets                     = "balanced"
  mode                             = "prevention"
  protected_networks_use_default   = false
  protected_networks_excluded_cidr = ["10.0.0.0/8"]
  protected_networks_included_cidr = ["10.0.0.0/8"]
}
