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

func TestAccDataSourceMerakiWirelessSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "ipv6_bridge_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "led_lights_on", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "location_analytics_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "meshing_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "upgrade_strategy", "minimizeUpgradeTime"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "named_vlans_pool_dhcp_monitoring_duration", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_settings.test", "named_vlans_pool_dhcp_monitoring_enabled", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSettingsPrerequisitesConfig + testAccDataSourceMerakiWirelessSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSettingsPrerequisitesConfig = `
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

func testAccDataSourceMerakiWirelessSettingsConfig() string {
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

	config += `
		data "meraki_wireless_settings" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_wireless_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
