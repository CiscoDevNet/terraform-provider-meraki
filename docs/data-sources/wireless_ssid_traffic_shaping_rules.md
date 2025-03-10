---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_traffic_shaping_rules Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless SSID Traffic Shaping Rules configuration.
---

# meraki_wireless_ssid_traffic_shaping_rules (Data Source)

This data source can read the `Wireless SSID Traffic Shaping Rules` configuration.

## Example Usage

```terraform
data "meraki_wireless_ssid_traffic_shaping_rules" "example" {
  network_id = "L_123456"
  number     = "0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Read-Only

- `default_rules_enabled` (Boolean) Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network`s traffic shaping page. Note that default rules count against the rule limit of 8.
- `id` (String) The id of the object
- `rules` (Attributes List) An array of traffic shaping rules. Rules are applied in the order that they are specified in. An empty list (or null) means no rules. Note that you are allowed a maximum of 8 rules. (see [below for nested schema](#nestedatt--rules))
- `traffic_shaping_enabled` (Boolean) Whether traffic shaping rules are applied to clients on your SSID.

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `definitions` (Attributes List) A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. (see [below for nested schema](#nestedatt--rules--definitions))
- `dscp_tag_value` (Number) The DSCP tag applied by your rule. null means `Do not change DSCP tag`. For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
- `pcp_tag_value` (Number) The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority). null means `Do not set PCP tag`.
- `per_client_bandwidth_limits_bandwidth_limits_limit_down` (Number) The maximum download limit (integer, in Kbps).
- `per_client_bandwidth_limits_bandwidth_limits_limit_up` (Number) The maximum upload limit (integer, in Kbps).
- `per_client_bandwidth_limits_settings` (String) How bandwidth limits are applied by your rule. Can be one of `network default`, `ignore` or `custom`.

<a id="nestedatt--rules--definitions"></a>
### Nested Schema for `rules.definitions`

Read-Only:

- `type` (String) The type of definition. Can be one of `application`, `applicationCategory`, `host`, `port`, `ipRange` or `localNet`.
- `value` (String) If 'type' is `host`, `port`, `ipRange` or `localNet`, then 'value' must be a string, matching either a hostname (e.g. 'somesite.com'), a port (e.g. 8080), or an IP range ('192.1.0.0', '192.1.0.0/16', or '10.1.0.0/16:80'). `localNet` also supports CIDR notation, excluding custom ports. If 'type' is `application` or `applicationCategory`, then 'value' must be an application category or application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories endpoint).
