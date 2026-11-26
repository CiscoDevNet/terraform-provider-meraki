resource "meraki_appliance_traffic_shaping_vpn_exclusions" "example" {
  network_id = "L_123456"
  applications = [
    {
      id          = "meraki:layer7/application/2"
      name        = "File Transfer Protocol"
      protocol    = "tcp"
      source_cidr = "192.168.1.0/24"
      source_port = "8000"
    }
  ]
  custom = [
    {
      destination = "192.168.3.0/24"
      port        = "8000"
      protocol    = "tcp"
      source_cidr = "192.168.1.0/24"
      source_port = "8000"
    }
  ]
  major_applications = [
    {
      id             = "meraki:vpnExclusions/application/2"
      name           = "Office 365 Sharepoint"
      protocol       = "tcp"
      source_vlan_id = "20"
    }
  ]
}
