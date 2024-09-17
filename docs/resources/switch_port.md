---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_switch_port Resource - terraform-provider-meraki"
subcategory: "Switches"
description: |-
  This resource can manage the Switch Port configuration.
---

# meraki_switch_port (Resource)

This resource can manage the `Switch Port` configuration.

## Example Usage

```terraform
resource "meraki_switch_port" "example" {
  serial                      = "1234-ABCD-1234"
  port_id                     = "1"
  access_policy_number        = 2
  access_policy_type          = "Sticky MAC allow list"
  adaptive_policy_group_id    = "123"
  allowed_vlans               = "1,3,5-10"
  dai_trusted                 = false
  enabled                     = true
  flexible_stacking_enabled   = true
  isolation_enabled           = false
  link_negotiation            = "Auto negotiate"
  name                        = "My switch port"
  peer_sgt_capable            = false
  poe_enabled                 = true
  port_schedule_id            = "1234"
  rstp_enabled                = true
  sticky_mac_allow_list_limit = 5
  storm_control_enabled       = true
  stp_guard                   = "disabled"
  type                        = "access"
  udld                        = "Alert only"
  vlan                        = 10
  voice_vlan                  = 20
  dot3az_enabled              = false
  profile_enabled             = false
  profile_id                  = "1284392014819"
  profile_iname               = "iname"
  mac_allow_list              = ["34:56:fe:ce:8e:a0"]
  sticky_mac_allow_list       = ["34:56:fe:ce:8e:b0"]
  tags                        = ["tag1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port_id` (String) Port ID
- `serial` (String) Switch serial

### Optional

- `access_policy_number` (Number) The number of a custom access policy to configure on the switch port. Only applicable when `accessPolicyType` is `Custom access policy`.
- `access_policy_type` (String) The type of the access policy of the switch port. Only applicable to access ports. Can be one of `Open`, `Custom access policy`, `MAC allow list` or `Sticky MAC allow list`.
  - Choices: `Custom access policy`, `MAC allow list`, `Open`, `Sticky MAC allow list`
- `adaptive_policy_group_id` (String) The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.
- `allowed_vlans` (String) The VLANs allowed on the switch port. Only applicable to trunk ports.
- `dai_trusted` (Boolean) If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
- `dot3az_enabled` (Boolean) The Energy Efficient Ethernet status of the switch port.
- `enabled` (Boolean) The status of the switch port.
- `flexible_stacking_enabled` (Boolean) For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
- `isolation_enabled` (Boolean) The isolation status of the switch port.
- `link_negotiation` (String) The link speed for the switch port.
- `mac_allow_list` (List of String) Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.
- `name` (String) The name of the switch port.
- `peer_sgt_capable` (Boolean) If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.
- `poe_enabled` (Boolean) The PoE status of the switch port.
- `port_schedule_id` (String) The ID of the port schedule. A value of null will clear the port schedule.
- `profile_enabled` (Boolean) When enabled, override this port`s configuration with a port profile.
- `profile_id` (String) When enabled, the ID of the port profile used to override the port`s configuration.
- `profile_iname` (String) When enabled, the IName of the profile.
- `rstp_enabled` (Boolean) The rapid spanning tree protocol status.
- `sticky_mac_allow_list` (List of String) The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.
- `sticky_mac_allow_list_limit` (Number) The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.
- `storm_control_enabled` (Boolean) The storm control status of the switch port.
- `stp_guard` (String) The state of the STP guard (`disabled`, `root guard`, `bpdu guard` or `loop guard`).
  - Choices: `bpdu guard`, `disabled`, `loop guard`, `root guard`
- `tags` (List of String) The list of tags of the switch port.
- `type` (String) The type of the switch port (`trunk`, `access` or `stack`).
  - Choices: `access`, `stack`, `trunk`
- `udld` (String) The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
  - Choices: `Alert only`, `Enforce`
- `vlan` (Number) The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
- `voice_vlan` (Number) The voice VLAN of the switch port. Only applicable to access ports.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_switch_port.example "<serial>,<port_id>"
```