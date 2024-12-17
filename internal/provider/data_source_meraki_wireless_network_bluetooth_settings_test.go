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

func TestAccDataSourceMerakiWirelessNetworkBluetoothSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "advertising_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "major", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "major_minor_assignment_mode", "Non-unique"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "minor", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "scanning_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_network_bluetooth_settings.test", "uuid", "00000000-0000-0000-0000-000000000000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessNetworkBluetoothSettingsPrerequisitesConfig + testAccDataSourceMerakiWirelessNetworkBluetoothSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessNetworkBluetoothSettingsPrerequisitesConfig = `
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessNetworkBluetoothSettingsConfig() string {
	config := `resource "meraki_wireless_network_bluetooth_settings" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  advertising_enabled = true` + "\n"
	config += `  major = 1` + "\n"
	config += `  major_minor_assignment_mode = "Non-unique"` + "\n"
	config += `  minor = 1` + "\n"
	config += `  scanning_enabled = true` + "\n"
	config += `  uuid = "00000000-0000-0000-0000-000000000000"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_network_bluetooth_settings" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_wireless_network_bluetooth_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
