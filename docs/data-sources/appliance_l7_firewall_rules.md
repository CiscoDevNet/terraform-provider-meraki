---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_l7_firewall_rules Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance L7 Firewall Rules configuration.
---

# meraki_appliance_l7_firewall_rules (Data Source)

This data source can read the `Appliance L7 Firewall Rules` configuration.

## Example Usage

```terraform
data "meraki_appliance_l7_firewall_rules" "example" {
  network_id = "L_123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object
- `rules` (Attributes List) An ordered array of the MX L7 firewall rules (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `policy` (String) `Deny` traffic specified by this rule
- `type` (String) Type of the L7 rule. One of: `application`, `applicationCategory`, `host`, `port`, `ipRange`, `blockedCountries`, `allowedCountries`
- `value` (String) The `value` of what you want to block. The application categories and application ids can be retrieved from the the `MX L7 application categories` endpoint.
- `value_countries` (List of String) If type is `blockedCountries` or `allowedCountries` this attribute should be used instead of `value`. A list of countries, that follow the two-letter ISO 3166-1 alpha-2 format.
