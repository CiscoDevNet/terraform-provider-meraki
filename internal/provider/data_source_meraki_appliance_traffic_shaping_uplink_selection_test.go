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

func TestAccDataSourceMerakiApplianceTrafficShapingUplinkSelection(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "default_uplink", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "load_balancing_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "failover_and_failback_immediate_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.preferred_uplink", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.type", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_cidr", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_port", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_cidr", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_traffic_shaping_uplink_selection.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_port", "1-1024"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceTrafficShapingUplinkSelectionPrerequisitesConfig + testAccDataSourceMerakiApplianceTrafficShapingUplinkSelectionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceTrafficShapingUplinkSelectionPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_appliance_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_appliance_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceTrafficShapingUplinkSelectionConfig() string {
	config := `resource "meraki_appliance_traffic_shaping_uplink_selection" "test" {` + "\n"
	config += `  network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `  default_uplink = "wan1"` + "\n"
	config += `  load_balancing_enabled = true` + "\n"
	config += `  failover_and_failback_immediate_enabled = true` + "\n"
	config += `  wan_traffic_uplink_preferences = [{` + "\n"
	config += `    preferred_uplink = "wan1"` + "\n"
	config += `    traffic_filters = [{` + "\n"
	config += `      type = "custom"` + "\n"
	config += `      protocol = "tcp"` + "\n"
	config += `      destination_cidr = "any"` + "\n"
	config += `      destination_port = "any"` + "\n"
	config += `      source_cidr = "any"` + "\n"
	config += `      source_port = "1-1024"` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_traffic_shaping_uplink_selection" "test" {
			network_id = meraki_network_device_claim.test.network_id
			depends_on = [meraki_appliance_traffic_shaping_uplink_selection.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
