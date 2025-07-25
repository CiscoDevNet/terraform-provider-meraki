---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_firewalled_service Resource - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This resource can manage the Appliance Firewalled Service configuration.
---

# meraki_appliance_firewalled_service (Resource)

This resource can manage the `Appliance Firewalled Service` configuration.

## Example Usage

```terraform
resource "meraki_appliance_firewalled_service" "example" {
  network_id  = "L_123456"
  service     = "ICMP"
  access      = "restricted"
  allowed_ips = ["123.123.123.1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access` (String) A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are 'blocked' (no remote IPs can access the service), 'restricted' (only allowed IPs can access the service), and 'unrestriced' (any remote IP can access the service). This field is required
  - Choices: `blocked`, `restricted`, `unrestricted`
- `network_id` (String) Network ID
- `service` (String) Service
  - Choices: `ICMP`, `SNMP`, `web`

### Optional

- `allowed_ips` (List of String) An array of allowed IPs that can access the service. This field is required if 'access' is set to 'restricted'. Otherwise this field is ignored

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_appliance_firewalled_service.example "<network_id>,<service>"
```
