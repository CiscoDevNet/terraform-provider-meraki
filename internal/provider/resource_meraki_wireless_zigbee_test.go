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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiWirelessZigbee(t *testing.T) {
	if os.Getenv("WIRELESS_ZIGBEE") == "" {
        t.Skip("skipping test, set environment variable WIRELESS_ZIGBEE")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_ap_1_serial") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_ap_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "defaults_channel", "25"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "defaults_transmit_power_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "lock_management_address", "10.100.100.200"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "lock_management_password", "password"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_zigbee.test", "lock_management_username", "user"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessZigbeePrerequisitesConfig+testAccMerakiWirelessZigbeeConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessZigbeePrerequisitesConfig+testAccMerakiWirelessZigbeeConfig_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_wireless_zigbee.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: merakiWirelessZigbeeImportStateIdFunc("meraki_wireless_zigbee.test"),
		ImportStateVerifyIgnore: []string{  },
		Check: resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessZigbeeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessZigbeePrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessZigbeeConfig_minimum() string {
	config := `resource "meraki_wireless_zigbee" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessZigbeeConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional



// End of section. //template:end testAccConfigAdditional
