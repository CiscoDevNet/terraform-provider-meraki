---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_appliance_dns_local_record Data Source - terraform-provider-meraki"
subcategory: "Appliances"
description: |-
  This data source can read the Appliance DNS Local Record configuration.
---

# meraki_appliance_dns_local_record (Data Source)

This data source can read the `Appliance DNS Local Record` configuration.

## Example Usage

```terraform
data "meraki_appliance_dns_local_record" "example" {
  id              = "12345678"
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object
- `organization_id` (String) Organization ID

### Read-Only

- `address` (String) IP for the DNS record
- `hostname` (String) Hostname for the DNS record
- `profile_id` (String) Profile ID
