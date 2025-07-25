---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization_licenses Data Source - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This data source can read the Organization Licenses configuration in bulk.
---

# meraki_organization_licenses (Data Source)

This data source can read the `Organization Licenses` configuration in bulk.

## Example Usage

```terraform
data "meraki_organization_licenses" "example" {
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

- `activation_date` (String) The date the license started burning
- `claim_date` (String) The date the license was claimed into the organization
- `device_serial` (String) Serial number of the device the license is assigned to
- `duration_in_days` (Number) The duration of the individual license
- `expiration_date` (String) The date the license will expire
- `head_license_id` (String) The id of the head license this license is queued behind. If there is no head license, it returns nil.
- `id` (String) The id of the object
- `license_key` (String) License key
- `license_type` (String) License type
- `network_id` (String) ID of the network the license is assigned to
- `order_number` (String) Order number
- `permanently_queued_licenses` (Attributes List) DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device. (see [below for nested schema](#nestedatt--items--permanently_queued_licenses))
- `seat_count` (Number) The number of seats of the license. Only applicable to SM licenses.
- `state` (String) The state of the license. All queued licenses have a status of `recentlyQueued`.
- `total_duration_in_days` (Number) The duration of the license plus all permanently queued licenses associated with it

<a id="nestedatt--items--permanently_queued_licenses"></a>
### Nested Schema for `items.permanently_queued_licenses`

Read-Only:

- `duration_in_days` (Number) The duration of the individual license
- `id` (String) Permanently queued license ID
- `license_key` (String) License key
- `license_type` (String) License type
- `order_number` (String) Order number
