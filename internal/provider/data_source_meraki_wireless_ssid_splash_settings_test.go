// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiWirelessSSIDSplashSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "allow_simultaneous_logins", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "block_all_traffic_before_sign_on", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "controller_disconnection_behavior", "default"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "redirect_url", "https://example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "splash_timeout", "1440"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "use_redirect_url", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "use_splash_url", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "welcome_message", "Welcome!"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "billing_prepaid_access_fast_login_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "billing_reply_to_email_address", "user@email.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "billing_free_access_duration_in_minutes", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "billing_free_access_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "guest_sponsorship_duration_in_minutes", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_splash_settings.test", "guest_sponsorship_guest_can_request_timeframe", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDSplashSettingsPrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDSplashSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDSplashSettingsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_wireless_ssid" "test" {
  network_id = meraki_network.test.id
  number     = "0"
  name       = "My SSID"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDSplashSettingsConfig() string {
	config := `resource "meraki_wireless_ssid_splash_settings" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	allow_simultaneous_logins = false` + "\n"
	config += `	block_all_traffic_before_sign_on = false` + "\n"
	config += `	controller_disconnection_behavior = "default"` + "\n"
	config += `	redirect_url = "https://example.com"` + "\n"
	config += `	splash_timeout = 1440` + "\n"
	config += `	use_redirect_url = true` + "\n"
	config += `	use_splash_url = false` + "\n"
	config += `	welcome_message = "Welcome!"` + "\n"
	config += `	billing_prepaid_access_fast_login_enabled = false` + "\n"
	config += `	billing_reply_to_email_address = "user@email.com"` + "\n"
	config += `	billing_free_access_duration_in_minutes = 20` + "\n"
	config += `	billing_free_access_enabled = false` + "\n"
	config += `	guest_sponsorship_duration_in_minutes = 30` + "\n"
	config += `	guest_sponsorship_guest_can_request_timeframe = false` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_splash_settings" "test" {
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_splash_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
