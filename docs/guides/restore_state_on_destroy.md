---
subcategory: "Guides"
page_title: "Restore State on Destroy"
description: |-
    Restore State on Destroy
---

# Restore State on Destroy (Experimental)

~>**This feature is experimental and may change in future releases. Use with caution and review the [Limitations](#limitations) section carefully.**

Many Meraki resources represent API-managed settings that always exist on the dashboard and cannot be truly deleted — only updated. These are referred to as **singleton resources**. By default, when you destroy a singleton resource in Terraform, it is simply removed from the Terraform state with no API call, or a set of hardcoded default values is sent to reset the configuration.

With the `restore_original_state_on_destroy` provider flag, you can opt in to a different behavior: the provider captures the original API state when the resource is first created and restores it when the resource is destroyed.

## Background: Singleton Resources

Singleton resources correspond to Meraki Dashboard API endpoints that represent settings or configurations which always exist for a given scope (network, device, or organization). They cannot be created or deleted via the API — only read and updated via PUT.

Examples of singleton resources include:

- `meraki_network_settings` — General network settings
- `meraki_switch_stp` — Spanning Tree Protocol settings
- `meraki_appliance_firewall_settings` — Appliance firewall settings
- `meraki_wireless_ssid` — SSID configuration (SSIDs always exist, they can only be enabled/disabled)

In contrast, standard resources like `meraki_network` or `meraki_organization_admin` support full create and delete operations via the API.

## Enabling the Feature

Add the `restore_original_state_on_destroy` attribute to your provider configuration:

```hcl
provider "meraki" {
  api_key                           = "YOUR_API_KEY"
  restore_original_state_on_destroy = true
}
```

Alternatively, use the environment variable:

```shell
export MERAKI_RESTORE_ORIGINAL_STATE_ON_DESTROY=true
```

The flag defaults to `false`, preserving the existing behavior from previous provider versions.

## How It Works

When `restore_original_state_on_destroy = true`:

- **On Create**: Before applying your desired configuration, the provider reads the current API state and saves it in Terraform's private state storage. It then applies your configuration as usual.
- **On Update**: No change in behavior — updates work the same regardless of this flag.
- **On Destroy**: The provider sends the saved initial state back to the API via PUT, restoring the original settings. The resource is then removed from the Terraform state.

## Default Behavior (Flag Disabled)

When `restore_original_state_on_destroy` is not set or set to `false`, the provider retains its existing destroy behavior:

- **Singleton resources with hardcoded defaults**: Some resources send a predefined set of default values on destroy. For example, firewall rule resources clear their rules list, and SSID resources reset to an unconfigured state.
- **Singleton resources without defaults**: The resource is simply removed from the Terraform state with no API call. The configuration on the Meraki Dashboard remains unchanged.

This behavior is fully backward-compatible with all previous provider versions.

## Limitations

~>**Read this section carefully before enabling the feature.**

### Imported Resources

Resources added to Terraform via `terraform import` will **not** have their initial state captured, because the import process does not trigger the Create lifecycle. If you later destroy an imported resource, it will fall back to the default behavior (hardcoded defaults or no-op).

### Resources Created Before Enabling the Flag

If you enable the flag on an existing Terraform deployment, resources that were already created (with the flag disabled or with an older provider version) will not have initial state saved in their private state. Destroying these resources will also fall back to the default behavior.

To benefit from state restoration, you would need to remove the resource from the Terraform state (`terraform state rm`) and re-create it with the flag enabled.

### Provider-Wide Setting

The flag applies globally to **all** singleton resources managed by this provider instance. It cannot be enabled or disabled on a per-resource basis.

### API State Drift

The saved initial state is a point-in-time snapshot captured when the resource was first created by Terraform. If the API state was modified outside of Terraform (e.g., via the Meraki Dashboard or another API client) between creation and destruction, the restored state will reflect the original snapshot, not any interim changes.

### New API Attributes

If the Meraki API introduces new attributes after the initial state was captured, those attributes will not be included in the saved snapshot. On destroy, the restored state will only contain the attributes that existed at the time of creation. Any new attributes added by later API versions will not be reset.

### Warnings on Failure

If the restore API call fails during destroy, the provider emits a **warning** rather than an error. The resource is always removed from the Terraform state regardless of whether the restore succeeded. This is intentional — a failed restore should not block infrastructure teardown.

## Behavioral Summary

| Scenario | Flag Disabled (default) | Flag Enabled |
|---|---|---|
| Create | Normal PUT | GET + save initial state, then PUT |
| Destroy (resource has hardcoded defaults) | PUT hardcoded defaults | PUT saved initial state |
| Destroy (resource has no defaults) | No-op (remove from state only) | PUT saved initial state |
| Destroy (restore API call fails) | N/A | Warning emitted, resource removed from state |
| Destroy after `terraform import` | Default behavior | Falls back to default behavior (no saved state) |
