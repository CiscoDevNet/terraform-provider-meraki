---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_cellular_gateway_subnet_pool Resource - terraform-provider-meraki"
subcategory: "Cellular Gateways"
description: |-
  This resource can manage the Cellular Gateway Subnet Pool configuration.
---

# meraki_cellular_gateway_subnet_pool (Resource)

This resource can manage the `Cellular Gateway Subnet Pool` configuration.

## Example Usage

```terraform
resource "meraki_cellular_gateway_subnet_pool" "example" {
  network_id = "L_123456"
  cidr       = "192.168.0.0/24"
  mask       = 26
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cidr` (String) CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
- `mask` (Number) Mask used for the subnet of all MGs in this network.
- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_cellular_gateway_subnet_pool.example "<network_id>"
```
