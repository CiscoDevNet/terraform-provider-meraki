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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiNetworkFirmwareUpgrades(t *testing.T) {
	if os.Getenv("NETWORK_FIRMWARE_UPGRADES") == "" {
		t.Skip("skipping test, set environment variable NETWORK_FIRMWARE_UPGRADES")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "timezone", "America/Los_Angeles"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_appliance_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_appliance_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_appliance_next_upgrade_to_version_id", "1001"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_camera_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_camera_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_camera_next_upgrade_to_version_id", "1003"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_cellular_gateway_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_cellular_gateway_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_cellular_gateway_next_upgrade_to_version_id", "1004"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_secure_connect_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_secure_connect_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_secure_connect_next_upgrade_to_version_id", "1007"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_sensor_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_sensor_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_sensor_next_upgrade_to_version_id", "1005"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_next_upgrade_to_version_id", "1002"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_catalyst_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_catalyst_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_switch_catalyst_next_upgrade_to_version_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_next_upgrade_to_version_id", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_controller_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_controller_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "products_wireless_controller_next_upgrade_to_version_id", "1006"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "upgrade_window_day_of_week", "sun"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_firmware_upgrades.test", "upgrade_window_hour_of_day", "4:00"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkFirmwareUpgradesPrerequisitesConfig + testAccMerakiNetworkFirmwareUpgradesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkFirmwareUpgradesPrerequisitesConfig + testAccMerakiNetworkFirmwareUpgradesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_firmware_upgrades.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkFirmwareUpgradesImportStateIdFunc("meraki_network_firmware_upgrades.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiNetworkFirmwareUpgradesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkFirmwareUpgradesPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkFirmwareUpgradesConfig_minimum() string {
	config := `resource "meraki_network_firmware_upgrades" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  timezone = "America/Los_Angeles"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkFirmwareUpgradesConfig_all() string {
	config := `resource "meraki_network_firmware_upgrades" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  timezone = "America/Los_Angeles"` + "\n"
	config += `  products_appliance_participate_in_next_beta_release = false` + "\n"
	config += `  products_appliance_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_appliance_next_upgrade_to_version_id = "1001"` + "\n"
	config += `  products_camera_participate_in_next_beta_release = false` + "\n"
	config += `  products_camera_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_camera_next_upgrade_to_version_id = "1003"` + "\n"
	config += `  products_cellular_gateway_participate_in_next_beta_release = false` + "\n"
	config += `  products_cellular_gateway_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_id = "1004"` + "\n"
	config += `  products_secure_connect_participate_in_next_beta_release = false` + "\n"
	config += `  products_secure_connect_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_id = "1007"` + "\n"
	config += `  products_sensor_participate_in_next_beta_release = false` + "\n"
	config += `  products_sensor_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_sensor_next_upgrade_to_version_id = "1005"` + "\n"
	config += `  products_switch_participate_in_next_beta_release = false` + "\n"
	config += `  products_switch_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_switch_next_upgrade_to_version_id = "1002"` + "\n"
	config += `  products_switch_catalyst_participate_in_next_beta_release = false` + "\n"
	config += `  products_switch_catalyst_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_switch_catalyst_next_upgrade_to_version_id = "1234"` + "\n"
	config += `  products_wireless_participate_in_next_beta_release = false` + "\n"
	config += `  products_wireless_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_wireless_next_upgrade_to_version_id = "1000"` + "\n"
	config += `  products_wireless_controller_participate_in_next_beta_release = false` + "\n"
	config += `  products_wireless_controller_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_id = "1006"` + "\n"
	config += `  upgrade_window_day_of_week = "sun"` + "\n"
	config += `  upgrade_window_hour_of_day = "4:00"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
