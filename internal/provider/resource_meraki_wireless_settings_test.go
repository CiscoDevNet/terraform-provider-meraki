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

func TestAccMerakiWirelessSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "ipv6_bridge_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "led_lights_on", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "location_analytics_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "meshing_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "upgrade_strategy", "minimizeUpgradeTime"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "named_vlans_pool_dhcp_monitoring_duration", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_settings.test", "named_vlans_pool_dhcp_monitoring_enabled", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessSettingsPrerequisitesConfig + testAccMerakiWirelessSettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSettingsPrerequisitesConfig + testAccMerakiWirelessSettingsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_settings.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessSettingsImportStateIdFunc("meraki_wireless_settings.test"),
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

func merakiWirelessSettingsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessSettingsPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessSettingsConfig_minimum() string {
	config := `resource "meraki_wireless_settings" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  ipv6_bridge_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessSettingsConfig_all() string {
	config := `resource "meraki_wireless_settings" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  ipv6_bridge_enabled = true` + "\n"
	config += `  led_lights_on = true` + "\n"
	config += `  location_analytics_enabled = false` + "\n"
	config += `  meshing_enabled = true` + "\n"
	config += `  upgrade_strategy = "minimizeUpgradeTime"` + "\n"
	config += `  named_vlans_pool_dhcp_monitoring_duration = 5` + "\n"
	config += `  named_vlans_pool_dhcp_monitoring_enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
