---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_network_client_policy Resource - terraform-provider-meraki"
subcategory: "Networks"
description: |-
  This resource can manage the Network Client Policy configuration.
---

# meraki_network_client_policy (Resource)

This resource can manage the `Network Client Policy` configuration.

## Example Usage

```terraform
resource "meraki_network_client_policy" "example" {
  network_id      = "L_123456"
  client_id       = "1.2.3.4"
  device_policy   = "Group policy"
  group_policy_id = "101"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_id` (String) Client ID. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
- `device_policy` (String) The policy to assign. Can be `Whitelisted`, `Blocked`, `Normal` or `Group policy`. Required.
- `network_id` (String) Network ID

### Optional

- `group_policy_id` (String) [Optional] If `devicePolicy` is set to `Group policy` this param is used to specify the group policy ID.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_network_client_policy.example "<network_id>,<client_id>"
```