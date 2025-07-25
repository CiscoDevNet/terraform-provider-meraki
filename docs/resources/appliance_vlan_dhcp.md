---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_vlan_dhcp Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource is meant to be used in addition to the meraki_appliance_vlan resource to configure DHCP settings for a VLAN. It requires the VLAN to be already configured on the appliance.
---

# meraki_appliance_vlan_dhcp (Resource)

This resource is meant to be used in addition to the `meraki_appliance_vlan` resource to configure DHCP settings for a VLAN. It requires the VLAN to be already configured on the appliance.

## Example Usage

```terraform
resource "meraki_appliance_vlan_dhcp" "example" {
  network_id                = "L_123456"
  vlan_id                   = "1234"
  dhcp_boot_options_enabled = false
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "1 day"
  dns_nameservers           = "upstream_dns"
  mandatory_dhcp_enabled    = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `vlan_id` (String) The VLAN ID of the new VLAN (must be between 1 and 4094)

### Optional

- `dhcp_boot_filename` (String) DHCP boot option for boot filename
- `dhcp_boot_next_server` (String) DHCP boot option to direct boot clients to the server to load the boot file from
- `dhcp_boot_options_enabled` (Boolean) Use DHCP boot options specified in other properties
- `dhcp_handling` (String) The appliance`s handling of DHCP requests on this VLAN. One of: `Run a DHCP server`, `Relay DHCP to another server` or `Do not respond to DHCP requests`
  - Choices: `Do not respond to DHCP requests`, `Relay DHCP to another server`, `Run a DHCP server`
- `dhcp_lease_time` (String) The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: `30 minutes`, `1 hour`, `4 hours`, `12 hours`, `1 day` or `1 week`
  - Choices: `1 day`, `1 hour`, `1 week`, `12 hours`, `30 minutes`, `4 hours`
- `dhcp_options` (Attributes List) The list of DHCP options that will be included in DHCP responses. Each object in the list should have 'code', 'type', and 'value' properties. (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcp_relay_server_ips` (List of String) The IPs of the DHCP servers that DHCP requests should be relayed to
- `dns_nameservers` (String) The DNS nameservers used for DHCP responses, either 'upstream_dns', 'google_dns', 'opendns', or a newline seperated string of IP addresses or domain names
- `mandatory_dhcp_enabled` (Boolean) Enable Mandatory DHCP on VLAN.
- `reserved_ip_ranges` (Attributes List) The DHCP reserved IP ranges on the VLAN (see [below for nested schema](#nestedatt--reserved_ip_ranges))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Required:

- `code` (String) The code for the DHCP option. This should be an integer between 2 and 254.
- `type` (String) The type for the DHCP option. One of: `text`, `ip`, `hex` or `integer`
  - Choices: `hex`, `integer`, `ip`, `text`
- `value` (String) The value for the DHCP option


<a id="nestedatt--reserved_ip_ranges"></a>
### Nested Schema for `reserved_ip_ranges`

Required:

- `comment` (String) A text comment for the reserved range
- `end` (String) The last IP in the reserved range
- `start` (String) The first IP in the reserved range

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_vlan_dhcp.example "<network_id>,<id>"
```
