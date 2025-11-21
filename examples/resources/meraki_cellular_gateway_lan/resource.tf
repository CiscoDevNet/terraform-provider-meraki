resource "meraki_cellular_gateway_lan" "example" {
  serial = "1234-ABCD-1234"
  fixed_ip_assignments = [
    {
      ip = "172.31.128.10"
      mac = "0b:00:00:00:00:ac"
      name = "server 1"
    }
  ]
  reserved_ip_ranges = [
    {
      comment = "A reserved IP range"
      end = "172.31.128.1"
      start = "172.31.128.0"
    }
  ]
}
