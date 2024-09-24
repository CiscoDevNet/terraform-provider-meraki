resource "meraki_appliance_content_filtering" "example" {
  network_id             = "L_123456"
  url_category_list_size = "topSites"
  allowed_url_patterns   = ["http://www.example.org"]
  blocked_url_categories = ["meraki:contentFiltering/category/C1"]
  blocked_url_patterns   = ["http://www.example.com"]
}
