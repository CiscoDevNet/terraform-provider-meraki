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

func TestAccDataSourceMerakiWirelessSSIDIdentityPSK(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_identity_psk.test", "expires_at", "2018-02-11T00:00:00.090209Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_identity_psk.test", "name", "Sample Identity PSK"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_identity_psk.test", "passphrase", "Cisco123"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDIdentityPSKPrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDIdentityPSKConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDIdentityPSKPrerequisitesConfig = `
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
resource "meraki_network_group_policy" "test" {
  network_id = meraki_network.test.id
  name       = "No video streaming"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDIdentityPSKConfig() string {
	config := `resource "meraki_wireless_ssid_identity_psk" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	expires_at = "2018-02-11T00:00:00.090209Z"` + "\n"
	config += `	group_policy_id = meraki_network_group_policy.test.id` + "\n"
	config += `	name = "Sample Identity PSK"` + "\n"
	config += `	passphrase = "Cisco123"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_identity_psk" "test" {
			id = meraki_wireless_ssid_identity_psk.test.id
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_identity_psk.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiWirelessSSIDIdentityPSKConfig() string {
	config := `resource "meraki_wireless_ssid_identity_psk" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	expires_at = "2018-02-11T00:00:00.090209Z"` + "\n"
	config += `	group_policy_id = meraki_network_group_policy.test.id` + "\n"
	config += `	name = "Sample Identity PSK"` + "\n"
	config += `	passphrase = "Cisco123"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_identity_psk" "test" {
			name = meraki_wireless_ssid_identity_psk.test.name
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_identity_psk.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
