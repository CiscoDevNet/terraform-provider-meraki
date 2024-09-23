---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_l3_firewall_rules Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless SSID L3 Firewall Rules configuration.
---

# meraki_wireless_ssid_l3_firewall_rules (Resource)

This resource can manage the `Wireless SSID L3 Firewall Rules` configuration.

## Example Usage

```terraform
resource "meraki_wireless_ssid_l3_firewall_rules" "example" {
  network_id       = "L_123456"
  number           = "0"
  allow_lan_access = true
  rules = [
    {
      comment   = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr = "Any"
      dest_port = "443"
      policy    = "allow"
      protocol  = "tcp"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `allow_lan_access` (Boolean) Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
- `rules` (Attributes List) An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `dest_cidr` (String) Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or `Any`
- `policy` (String) `allow` or `deny` traffic specified by this rule
  - Choices: `allow`, `deny`
- `protocol` (String) The type of protocol (must be `tcp`, `udp`, `icmp`, `icmp6` or `any`)
  - Choices: `any`, `icmp`, `icmp6`, `tcp`, `udp`

Optional:

- `comment` (String) Description of the rule (optional)
- `dest_port` (String) Comma-separated list of destination port(s) (integer in the range 1-65535), or `any`

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_wireless_ssid_l3_firewall_rules.example "<network_id>,<number>"
```