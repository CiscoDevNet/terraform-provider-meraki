---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_l7_firewall_rules Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance L7 Firewall Rules configuration.
---

# meraki_appliance_l7_firewall_rules (Resource)

This resource can manage the `Appliance L7 Firewall Rules` configuration.

## Example Usage

```terraform
resource "meraki_appliance_l7_firewall_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      policy = "deny"
      type   = "host"
      value  = "google.com"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `rules` (Attributes List) An ordered array of the MX L7 firewall rules (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `policy` (String) `Deny` traffic specified by this rule
  - Choices: `deny`
- `type` (String) Type of the L7 rule. One of: `application`, `applicationCategory`, `host`, `port`, `ipRange`, `blockedCountries`, `allowedCountries`
  - Choices: `application`, `applicationCategory`, `host`, `ipRange`, `port`, `blockedCountries`, `allowedCountries`
- `value` (String) The `value` of what you want to block. The application categories and application ids can be retrieved from the the `MX L7 application categories` endpoint.
- `value_countries` (List of String) If type is `blockedCountries` or `allowedCountries` this attribute should be used instead of `value`. A list of countries, that follow the two-letter ISO 3166-1 alpha-2 format.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_l7_firewall_rules.example "<network_id>"
```
