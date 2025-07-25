---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_splash_settings Resource - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This resource can manage the Wireless SSID Splash Settings configuration.
---

# meraki_wireless_ssid_splash_settings (Resource)

This resource can manage the `Wireless SSID Splash Settings` configuration.

## Example Usage

```terraform
resource "meraki_wireless_ssid_splash_settings" "example" {
  network_id                                    = "L_123456"
  number                                        = "0"
  allow_simultaneous_logins                     = false
  block_all_traffic_before_sign_on              = false
  controller_disconnection_behavior             = "default"
  redirect_url                                  = "https://example.com"
  splash_timeout                                = 1440
  use_redirect_url                              = true
  use_splash_url                                = false
  welcome_message                               = "Welcome!"
  billing_prepaid_access_fast_login_enabled     = false
  billing_reply_to_email_address                = "user@email.com"
  billing_free_access_duration_in_minutes       = 20
  billing_free_access_enabled                   = false
  guest_sponsorship_duration_in_minutes         = 30
  guest_sponsorship_guest_can_request_timeframe = false
  self_registration_authorization_type          = "admin"
  self_registration_enabled                     = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `allow_simultaneous_logins` (Boolean) Whether or not to allow simultaneous logins from different devices.
- `billing_free_access_duration_in_minutes` (Number) How long a device can use a network for free.
- `billing_free_access_enabled` (Boolean) Whether or not free access is enabled.
- `billing_prepaid_access_fast_login_enabled` (Boolean) Whether or not billing uses the fast login prepaid access option.
- `billing_reply_to_email_address` (String) The email address that receives replies from clients.
- `block_all_traffic_before_sign_on` (Boolean) How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
- `controller_disconnection_behavior` (String) How login attempts should be handled when the controller is unreachable. Can be either `open`, `restricted`, or `default`.
  - Choices: `default`, `open`, `restricted`
- `guest_sponsorship_duration_in_minutes` (Number) Duration in minutes of sponsored guest authorization. Must be between 1 and 60480 (6 weeks)
- `guest_sponsorship_guest_can_request_timeframe` (Boolean) Whether or not guests can specify how much time they are requesting.
- `redirect_url` (String) The custom redirect URL where the users will go after the splash page.
- `self_registration_authorization_type` (String) How created user accounts should be authorized. Must be included in: [admin, auto, self_email]
  - Choices: `admin`, `auto`, `self_email`
- `self_registration_enabled` (Boolean) Whether or not to allow users to create their own account on the network.
- `sentry_enrollment_enforced_systems` (Set of String) The system types that the Sentry enforces. Must be included in: `iOS, `Android`, `macOS`, and `Windows`.
- `sentry_enrollment_strength` (String) The strength of the enforcement of selected system types. Must be one of: `focused`, `click-through`, and `strict`.
  - Choices: `click-through`, `focused`, `strict`
- `sentry_enrollment_systems_manager_network_id` (String) The network ID of the Systems Manager network.
- `splash_image_extension` (String) The extension of the image file.
- `splash_image_image_contents` (String) The file contents (a base 64 encoded string) of your new image.
- `splash_image_image_format` (String) The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.
  - Choices: `gif`, `jpg`, `png`
- `splash_image_md5` (String) The MD5 value of the image file. Setting this to null will remove the image from the splash page.
- `splash_logo_extension` (String) The extension of the logo file.
- `splash_logo_image_contents` (String) The file contents (a base 64 encoded string) of your new logo.
- `splash_logo_image_format` (String) The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.
  - Choices: `gif`, `jpg`, `png`
- `splash_logo_md5` (String) The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.
- `splash_prepaid_front_extension` (String) The extension of the prepaid front image file.
- `splash_prepaid_front_image_contents` (String) The file contents (a base 64 encoded string) of your new prepaid front.
- `splash_prepaid_front_image_format` (String) The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.
  - Choices: `gif`, `jpg`, `png`
- `splash_prepaid_front_md5` (String) The MD5 value of the prepaid front image file. Setting this to null will remove the prepaid front from the splash page.
- `splash_timeout` (Number) Splash timeout in minutes. This will determine how often users will see the splash page.
- `splash_url` (String) [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see `useSplashUrl`
- `theme_id` (String) The id of the selected splash theme.
- `use_redirect_url` (Boolean) The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.
- `use_splash_url` (Boolean) [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID`s access control settings, it may not be possible to use the custom splash URL.
- `welcome_message` (String) The welcome message for the users on the splash page.

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_wireless_ssid_splash_settings.example "<network_id>,<number>"
```
