resource "meraki_wireless_ethernet_port_profile" "example" {
  network_id = "L_123456"
  name       = "name"
  ports = [
    {
      enabled = true
      name    = "port"
      ssid    = 1
    }
  ]
  usb_ports = [
    {
      enabled = true
      name    = "usb port"
      ssid    = 2
    }
  ]
}
