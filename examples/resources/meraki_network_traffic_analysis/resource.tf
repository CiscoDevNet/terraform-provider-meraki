resource "meraki_network_traffic_analysis" "example" {
  network_id = "L_123456"
  mode       = "basic"
  custom_pie_chart_items = [
    {
      name  = "Item from hostname"
      type  = "host"
      value = "example.com"
    }
  ]
}
