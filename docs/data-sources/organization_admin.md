---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization_admin Data Source - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This data source can read the Organization Admin configuration.
---

# meraki_organization_admin (Data Source)

This data source can read the `Organization Admin` configuration.

## Example Usage

```terraform
data "meraki_organization_admin" "example" {
  id              = "12345678"
  organization_id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) Organization ID

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the dashboard administrator

### Read-Only

- `authentication_method` (String) No longer used as of Cisco SecureX end-of-life. Can be one of `Email`. The default is Email authentication.
- `email` (String) The email of the dashboard administrator. This attribute can not be updated.
- `networks` (Attributes List) The list of networks that the dashboard administrator has privileges on (see [below for nested schema](#nestedatt--networks))
- `org_access` (String) The privilege of the dashboard administrator on the organization. Can be one of `full`, `read-only`, `enterprise` or `none`
- `tags` (Attributes List) The list of tags that the dashboard administrator has privileges on (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Read-Only:

- `access` (String) The privilege of the dashboard administrator on the network. Can be one of `full`, `read-only`, `guest-ambassador` or `monitor-only`
- `id` (String) The network ID


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `access` (String) The privilege of the dashboard administrator on the tag. Can be one of `full`, `read-only`, `guest-ambassador` or `monitor-only`
- `tag` (String) The name of the tag