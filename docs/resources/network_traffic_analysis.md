---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_traffic_analysis Resource - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This resource can manage the Network Traffic Analysis configuration.
---

# meraki_network_traffic_analysis (Resource)

This resource can manage the `Network Traffic Analysis` configuration.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Optional

- `custom_pie_chart_items` (Attributes List) The list of items that make up the custom pie chart for traffic reporting. (see [below for nested schema](#nestedatt--custom_pie_chart_items))
- `mode` (String) The traffic analysis mode for the network. Can be one of `disabled` (do not collect traffic types), `basic` (collect generic traffic categories), or `detailed` (collect destination hostnames).
  - Choices: `basic`, `detailed`, `disabled`

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--custom_pie_chart_items"></a>
### Nested Schema for `custom_pie_chart_items`

Required:

- `name` (String) The name of the custom pie chart item.
- `type` (String) The signature type for the custom pie chart item. Can be one of `host`, `port` or `ipRange`.
  - Choices: `host`, `ipRange`, `port`
- `value` (String) The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item (see sample request/response for more details).

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_network_traffic_analysis.example "<network_id>"
```