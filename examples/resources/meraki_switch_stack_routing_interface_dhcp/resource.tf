resource "meraki_switch_stack_routing_interface_dhcp" "example" {
  network_id = "L_123456"
  switch_stack_id = "1234"
  interface_id = "1234"
  boot_file_name = "home_boot_file"
  boot_next_server = "1.2.3.4"
  boot_options_enabled = true
  dhcp_lease_time = "1 day"
  dhcp_mode = "dhcpServer"
  dns_nameservers_option = "custom"
  dhcp_options = [
    {
      code = "5"
      type = "text"
      value = "five"
    }
  ]
  dns_custom_nameservers = ["8.8.8.8"]
  fixed_ip_assignments = [
    {
      ip = "192.168.1.12"
      mac = "22:33:44:55:66:77"
      name = "Cisco Meraki valued client"
    }
  ]
  reserved_ip_ranges = [
    {
      comment = "A reserved IP range"
      end = "192.168.1.10"
      start = "192.168.1.1"
    }
  ]
}
