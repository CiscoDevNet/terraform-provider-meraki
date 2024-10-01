---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_one_to_one_nat_rules Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance One To One NAT Rules configuration.
---

# meraki_appliance_one_to_one_nat_rules (Resource)

This resource can manage the `Appliance One To One NAT Rules` configuration.

## Example Usage

```terraform
resource "meraki_appliance_one_to_one_nat_rules" "example" {
  network_id = "L_123456"
  rules = [
    {
      lan_ip    = "192.168.128.22"
      name      = "Service behind NAT"
      public_ip = "146.12.3.33"
      uplink    = "internet1"
      allowed_inbound = [
        {
          protocol          = "tcp"
          allowed_ips       = ["10.82.112.0/24"]
          destination_ports = ["80"]
        }
      ]
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `rules` (Attributes List) An array of 1:1 nat rules (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `lan_ip` (String) The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
- `public_ip` (String) The IP address that will be used to access the internal resource from the WAN

Optional:

- `allowed_inbound` (Attributes List) The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource (see [below for nested schema](#nestedatt--rules--allowed_inbound))
- `name` (String) A descriptive name for the rule
- `uplink` (String) The physical WAN interface on which the traffic will arrive (`internet1` or, if available, `internet2`)
  - Choices: `internet1`, `internet2`

<a id="nestedatt--rules--allowed_inbound"></a>
### Nested Schema for `rules.allowed_inbound`

Optional:

- `allowed_ips` (List of String) An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or `any`
- `destination_ports` (List of String) An array of ports or port ranges that will be forwarded to the host on the LAN
- `protocol` (String) Either of the following: `tcp`, `udp`, `icmp-ping` or `any`
  - Choices: `any`, `icmp-ping`, `tcp`, `udp`

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_appliance_one_to_one_nat_rules.example "<network_id>"
```