resource "meraki_switch_ports" "example" {
  serial          = "1234-ABCD-1234"
  organization_id = "123456"
  items = [{
    port_id                     = "1"
    access_policy_type          = "Sticky MAC allow list"
    allowed_vlans               = "1,3,5-10"
    dai_trusted                 = false
    enabled                     = true
    isolation_enabled           = false
    link_negotiation            = "Auto negotiate"
    name                        = "My switch port"
    poe_enabled                 = false
    rstp_enabled                = true
    sticky_mac_allow_list_limit = 5
    stp_guard                   = "disabled"
    type                        = "access"
    udld                        = "Alert only"
    vlan                        = 10
    voice_vlan                  = 20
    profile_enabled             = false
    sticky_mac_allow_list       = ["34:56:fe:ce:8e:b0"]
    tags                        = ["tag1"]
  }]
}
