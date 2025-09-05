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

func TestAccDataSourceMerakiWirelessZigbee(t *testing.T) {
	if os.Getenv("WIRELESS_ZIGBEE") == "" {
		t.Skip("skipping test, set environment variable WIRELESS_ZIGBEE")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_ap_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_ap_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "defaults_channel", "25"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "defaults_transmit_power_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "lock_management_address", "10.100.100.200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "lock_management_password", "password"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_zigbee.test", "lock_management_username", "user"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessZigbeePrerequisitesConfig + testAccDataSourceMerakiWirelessZigbeeConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessZigbeePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_ap_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_ap_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessZigbeeConfig() string {
	config := `resource "meraki_wireless_zigbee" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  enabled = false` + "\n"
	config += `  defaults_channel = "25"` + "\n"
	config += `  defaults_transmit_power_level = 10` + "\n"
	config += `  iot_controller_serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  lock_management_address = "10.100.100.200"` + "\n"
	config += `  lock_management_password = "password"` + "\n"
	config += `  lock_management_username = "user"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_zigbee" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_wireless_zigbee.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
