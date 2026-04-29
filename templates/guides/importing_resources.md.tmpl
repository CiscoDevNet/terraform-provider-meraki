---
subcategory: "Guides"
page_title: "Importing Resources"
description: |-
    Importing Resources
---

# Importing Resources

Existing Meraki resources can be imported into Terraform state using the `terraform import` command or by using an `import` block in the configuration.

## String-Based Import

Resources can be imported using a comma-separated string of the resource's identity attributes. For example, to import an organization admin:

```terraform
import {
  to = meraki_organization_admin.example
  id = "<organization_id>,<id>"
}
```

Or using the CLI:

```shell
terraform import meraki_organization_admin.example "<organization_id>,<id>"
```

For resources scoped to a network with additional path parameters, all identity attributes must be included in order:

```terraform
import {
  to = meraki_wireless_ssid_identity_psk.example
  id = "<network_id>,<number>,<id>"
}
```

For resources with a single identity attribute, only that value is needed:

```terraform
import {
  to = meraki_organization.example
  id = "<id>"
}
```

## Identity-Based Import

Starting with Terraform 1.12, resources also support identity-based imports using an `identity` block. This approach is more readable and less error-prone than comma-separated strings, since each attribute is explicitly named.

```terraform
import {
  to = meraki_organization_admin.example
  identity = {
    "organization_id" = "<organization_id>"
    "id"              = "<id>"
  }
}
```

A network-scoped resource with multiple identity attributes:

```terraform
import {
  to = meraki_wireless_ssid_identity_psk.example
  identity = {
    "network_id" = "<network_id>"
    "number"     = "<number>"
    "id"         = "<id>"
  }
}
```

Each resource's documentation page includes an `Import` section with the exact identity attributes required for that resource.

## Compatibility

The provider supports both import methods. Identity-based imports require Terraform 1.12 or later. On older Terraform versions, use string-based imports.
