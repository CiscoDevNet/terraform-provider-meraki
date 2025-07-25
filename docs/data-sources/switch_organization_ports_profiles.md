---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_organization_ports_profiles Data Source - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This data source can read the Switch Organization Ports Profile configuration in bulk.
  ~>Warning: This resource or data source depends on an Early Access API endpoint. These API endpoints are subject to breaking changes without prior notice.
---

# meraki_switch_organization_ports_profiles (Data Source)

This data source can read the `Switch Organization Ports Profile` configuration in bulk.

~>Warning: This resource or data source depends on an Early Access API endpoint. These API endpoints are subject to breaking changes without prior notice.

## Example Usage

```terraform
data "meraki_switch_organization_ports_profiles" "example" {
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) Organization ID

### Read-Only

- `items` (Attributes Set) The list of items (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `description` (String) Text describing the profile.
- `id` (String) The id of the object
- `is_organization_wide` (Boolean) The scope of the profile whether it is organization level or network level
- `name` (String) The name of the profile.
- `network_id` (String) The network identifier
- `networks_type` (String) Flag to identify if the networks networks are excluded or included
- `networks_values` (Attributes List) The network object containing name and id (see [below for nested schema](#nestedatt--items--networks_values))
- `port_access_policy_number` (Number) The number of a custom access policy to configure on the port profile. Only applicable when `accessPolicyType` is `Custom access policy`.
- `port_access_policy_type` (String) The type of the access policy of the port profile. Only applicable to access ports.
- `port_adaptive_policy_group_id` (String) The adaptive policy group ID that will be used to tag traffic through this port profile. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API.
- `port_allowed_vlans` (String) The VLANs allowed on the port profile. Only applicable to trunk ports.
- `port_dai_trusted` (Boolean) If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
- `port_isolation_enabled` (Boolean) The isolation status of the port profile.
- `port_mac_allow_list` (List of String) Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.
- `port_peer_sgt_capable` (Boolean) If true, Peer SGT is enabled for traffic through this port profile. Applicable to trunk ports only.
- `port_poe_enabled` (Boolean) The PoE status of the port profile.
- `port_rstp_enabled` (Boolean) The rapid spanning tree protocol status.
- `port_sticky_mac_allow_list` (List of String) The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.
- `port_sticky_mac_allow_list_limit` (Number) The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.
- `port_storm_control_enabled` (Boolean) The storm control status of the port profile.
- `port_stp_guard` (String) The state of the STP guard.
- `port_type` (String) The type of the port profile.
- `port_udld` (String) The action to take when Unidirectional Link is detected. LinkDefault configuration is Alert only.
- `port_vlan` (Number) The VLAN of the port profile. A null value will clear the value set for trunk ports.
- `port_voice_vlan` (Number) The voice VLAN of the port profile. Only applicable to access ports.
- `tags` (List of String) Space-seperated list of tags

<a id="nestedatt--items--networks_values"></a>
### Nested Schema for `items.networks_values`

Read-Only:

- `id` (String) The network identifier
- `name` (String) The network name
