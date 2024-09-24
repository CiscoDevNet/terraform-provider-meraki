---
subcategory: "Guides"
page_title: "Getting Started"
description: |-
    Getting Started
---

# Getting Started

This example demonstrates how the provider can be used to configure a new network with a switch. The full example can be found here: [https://github.com/CiscoDevNet/terraform-provider-meraki/tree/main/examples/basic/getting_started](https://github.com/CiscoDevNet/terraform-provider-meraki/tree/main/examples/basic/getting_started)

First of all we need to add the necessary provider configuration to the Terraform configuration file:

```hcl
terraform {
  required_providers {
    meraki = {
      source = "CiscoDevNet/meraki"
    }
  }
}

provider "meraki" {
  api_key = "YOUR_API_KEY"
}
```

Alternatively to the `api_key` you can also use the `MERAKI_API_KEY` environment variable.

We want to add our network to an existing organization, we therefore need to get the related organization ID first. We can use the `meraki_organization` data source to retrieve the organization ID by name.

```hcl
data "meraki_organization" "org1" {
  name = "Org1"
}
```
Once we have the organization ID we can create a new network in this organization.

```hcl
resource "meraki_network" "tf_1" {
  organization_id = data.meraki_organization.org1.id
  name            = "TF-1"
  notes           = "My first network created by Terraform"
  product_types   = ["switch", "wireless"]
}
```

Next we want to add a switch to the network. We can use the `meraki_network_device_claim` resource to claim a switch to the network.

```hcl
resource "meraki_network_device_claim" "switch" {
  network_id = meraki_network.tf_1.id
  serials    = ["1234-1234-1234"]
}
```

Using the `meraki_device` resource we can configure some additional device information.

```hcl
resource "meraki_device" "switch" {
  serial  = tolist(meraki_network_device_claim.switch.serials)[0]
  address = "1600 Pennsylvania Ave"
  name    = "TF-Switch-1"
  notes   = "My first switch configured by Terraform"
}
```

Additional configuration can be applied to the switch by making use of the `meraki_switch_*` resources.

```hcl
resource "meraki_switch_mtu" "mtu" {
  network_id       = meraki_network.tf_1.id
  default_mtu_size = 9100
}

resource "meraki_switch_port" "port_5_10" {
  count   = 6
  serial  = tolist(meraki_network_device_claim.switch.serials)[0]
  port_id = count.index + 5
  enabled = true
  name    = "My first switch ports configured by Terraform"
  type    = "access"
  vlan    = 10
}
```

Once the initial configuration is complete we can run `terraform apply` to create the network and configure the switch accordingly. All the desired configuration can then be compared against the current state by running `terraform plan` to detect potential configuration drift. To simulate an out-of-band change we can log into Meraki Dashboard and change the name of a switch port. Running `terraform plan` again will show the drift and allow us to correct it by running `terraform apply`.

```shell
$ terraform apply
...
Terraform will perform the following actions:

  # meraki_switch_port.port_5_10[4] will be updated in-place
  ~ resource "meraki_switch_port" "port_5_10" {
        id      = "9"
      ~ name    = "Manually changed" -> "My first switch ports configured by Terraform"
        # (5 unchanged attributes hidden)
    }

Plan: 0 to add, 1 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

meraki_switch_port.port_5_10[4]: Modifying... [id=9]
meraki_switch_port.port_5_10[4]: Modifications complete after 0s [id=9]

Apply complete! Resources: 0 added, 1 changed, 0 destroyed.
```

In typical Terraform fashion we can also remove all the resources managed by Terraform with a single command by running `terraform destroy`.

```shell
$ terraform destroy
...
Plan: 0 to add, 0 to change, 10 to destroy.

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

meraki_switch_mtu.mtu: Destroying... [id=L_709316941310864274]
meraki_device.switch: Destroying...
meraki_switch_port.port_5_10[2]: Destroying... [id=7]
meraki_switch_port.port_5_10[1]: Destroying... [id=6]
meraki_switch_port.port_5_10[4]: Destroying... [id=9]
meraki_switch_port.port_5_10[0]: Destroying... [id=5]
meraki_switch_port.port_5_10[5]: Destroying... [id=10]
meraki_switch_port.port_5_10[3]: Destroying... [id=8]
meraki_device.switch: Destruction complete after 0s
meraki_switch_mtu.mtu: Destruction complete after 0s
meraki_switch_port.port_5_10[2]: Destruction complete after 0s
meraki_switch_port.port_5_10[1]: Destruction complete after 0s
meraki_switch_port.port_5_10[4]: Destruction complete after 0s
meraki_switch_port.port_5_10[3]: Destruction complete after 0s
meraki_switch_port.port_5_10[0]: Destruction complete after 0s
meraki_switch_port.port_5_10[5]: Destruction complete after 0s
meraki_network_device_claim.switch: Destroying... [id=L_709316941310864274]
meraki_network_device_claim.switch: Destruction complete after 3s
meraki_network.tf_1: Destroying... [id=L_709316941310864274]
meraki_network.tf_1: Destruction complete after 1s

Destroy complete! Resources: 10 destroyed.
