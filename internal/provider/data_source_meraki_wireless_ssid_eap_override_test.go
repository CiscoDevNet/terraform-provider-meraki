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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiWirelessSSIDEAPOverride(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "max_retries", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "eapol_key_retries", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "eapol_key_timeout_in_ms", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "identity_retries", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_eap_override.test", "identity_timeout", "5"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDEAPOverridePrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDEAPOverrideConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDEAPOverridePrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}
resource "meraki_wireless_ssid" "test" {
  number = "0"
  name = "My SSID"
  network_id = meraki_network.test.id
  ip_assignment_mode = "Bridge mode"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDEAPOverrideConfig() string {
	config := `resource "meraki_wireless_ssid_eap_override" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	max_retries = 5` + "\n"
	config += `	timeout = 5` + "\n"
	config += `	eapol_key_retries = 4` + "\n"
	config += `	eapol_key_timeout_in_ms = 5000` + "\n"
	config += `	identity_retries = 5` + "\n"
	config += `	identity_timeout = 5` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_eap_override" "test" {
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_eap_override.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
