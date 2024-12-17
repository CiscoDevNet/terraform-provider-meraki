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

func TestAccDataSourceMerakiWirelessDeviceBluetoothSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_ap_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_ap_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_device_bluetooth_settings.test", "major", "13"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_device_bluetooth_settings.test", "minor", "125"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_device_bluetooth_settings.test", "uuid", "00000000-0000-0000-0000-000000000000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessDeviceBluetoothSettingsPrerequisitesConfig + testAccDataSourceMerakiWirelessDeviceBluetoothSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessDeviceBluetoothSettingsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_ap_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_wireless_network_bluetooth_settings" "test" {
  network_id          = meraki_network.test.id
  advertising_enabled = true
  uuid                = "00000000-0000-0000-0000-000000000000"
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_wireless_network_bluetooth_settings.test.network_id
  serials    = [var.test_ap_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessDeviceBluetoothSettingsConfig() string {
	config := `resource "meraki_wireless_device_bluetooth_settings" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  major = 13` + "\n"
	config += `  minor = 125` + "\n"
	config += `  uuid = "00000000-0000-0000-0000-000000000000"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_device_bluetooth_settings" "test" {
			serial = tolist(meraki_network_device_claim.test.serials)[0]
			depends_on = [meraki_wireless_device_bluetooth_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
