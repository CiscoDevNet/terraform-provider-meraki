---
subcategory: "Guides"
page_title: "Bulk Resources"
description: |-
    Bulk Resources
---

# Bulk Resources

Bulk resources are designed to manage multiple objects of the same type within a single resource block. Instead of creating and managing each object individually, bulk resources allow you to define a list of objects, which are then managed collectively. This approach is particularly useful for scenarios where you need to create, update, or delete many similar objects at once, such as switch ports, policy objects, etc. Bulk resources require significantly fewer API calls, making them more efficient for large-scale operations.

->Using bulk resources helps avoid hitting Meraki API rate limits, as fewer API calls are required for large-scale changes.

## How Bulk Resources Work

Bulk resources leverage "Action Batches", which is a Meraki Dashboard API mechanism for submitting batched configuration requests in a single synchronous or asynchronous transaction. These endpoints allow multiple create, update, or delete actions to be performed in a single API call. By grouping actions together, bulk resources require significantly fewer API calls compared to managing each object separately.

For more information, see the [Meraki Action Batches documentation](https://developer.cisco.com/meraki/api-v1/action-batches-overview/).

## Example Usage

Assuming you have a 48 port switch and you want to configure all ports, you can use a bulk resource to define the desired state for all ports in a single resource block. Here's an example of how you might define such a resource in Terraform:

```hcl
resource "meraki_switch_ports" "switch_1" {
  serial          = "1234-ABCD-1234"
  organization_id = "123456"
  items = [
    {
      port_id  = "1"
      name     = "Workstation 1"
      type     = "access"
      vlan     = 10
    },
    {
      port_id  = "2"
      name     = "Workstation 2"
      type     = "access"
      vlan     = 10
    },
    {
      port_id  = "3"
      name     = "Workstation 3"
      type     = "access"
      vlan     = 10
    },
    {
      port_id  = "4"
      name     = "Workstation 4"
      type     = "access"
      vlan     = 10
    }
  ]
}
```

## Changing Objects

The collection of objects (`items`) is managed as a `Set` which is a data structure that allows for the storage of unique elements without ordering. This means that when you modify the `items` collection, you are actually adding or removing elements from a set, rather than changing the order of elements. As a result, the Meraki Terraform provider can efficiently determine which elements have been added, removed, or changed, and only send the necessary updates to the Meraki API.

Please note, however, that the use of a `Set` might show unexpected plan outputs as elements within a set are immutable, and therefore any changes to an element will show as a new element being added and the old element being removed. In this example, the vlan of a single port is being changed from 10 to 20:

```
Terraform will perform the following actions:

  # meraki_switch_ports.test will be updated in-place
  ~ resource "meraki_switch_ports" "test" {
        id              = "1576869"
      ~ items           = [
          - {
              - name    = "Workstation 3" -> null
              - port_id = "3" -> null
              - type    = "access" -> null
              - vlan    = 10 -> null
            },
          + {
              + name    = "Workstation 3"
              + port_id = "3"
              + type    = "access"
              + vlan    = 20
            },
            # (3 unchanged elements hidden)
        ]
        # (2 unchanged attributes hidden)
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

->The Terraform plan output may look more dramatic than the actual changes performed by the provider. Even if the plan shows an element being "removed" and "added", the provider will only update the existing object in place via the Meraki API, whenever possible.

To understand which elements have changed, been added or deleted, we need to uniquely identify each element in the collection. This is typically done using a combination of the element's attributes. In the above example (`meraki_switch_ports` resource) the `port_id` attribute is used to uniquely identify each port. Be aware, that whenever one of the attributes that contributes to the uniqueness of an element is changed, the provider will delete and recreate the element. The attributes that uniquely identify each element are listed in the documentation.
