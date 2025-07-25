---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organization_saml Resource - terraform-provider-meraki"
subcategory: "Organizations"
description: |-
  This resource can manage the Organization SAML configuration.
---

# meraki_organization_saml (Resource)

This resource can manage the `Organization SAML` configuration.

## Example Usage

```terraform
resource "meraki_organization_saml" "example" {
  organization_id = "123456"
  enabled         = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) Organization ID

### Optional

- `enabled` (Boolean) Boolean for updating SAML SSO enabled settings.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_organization_saml.example "<organization_id>"
```
