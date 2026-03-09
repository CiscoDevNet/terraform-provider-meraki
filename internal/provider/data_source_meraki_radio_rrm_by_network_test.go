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

func TestAccDataSourceMerakiRadioRRMByNetwork(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "meta_counts_items_remaining", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "meta_counts_items_total", "42"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.name", "My Network"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.network_id", "L_12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.time_zone", "America/Los_Angeles"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.ai_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.ai_last_enabled_at", "2026-02-05T14:49:07Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_minimize_changes_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_schedule_mode", "automatic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_schedule_automatic_end", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_schedule_automatic_start", "08:00"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_schedule_manual_end", "15:00"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.busy_hour_schedule_manual_start", "10:00"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.channel_avoidance_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_radio_rrm_by_network.test", "items.0.fra_enabled", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiRadioRRMByNetworkPrerequisitesConfig + testAccDataSourceMerakiRadioRRMByNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiRadioRRMByNetworkPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiRadioRRMByNetworkConfig() string {
	config := `resource "meraki_radio_rrm_by_network" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  meta_counts_items_remaining = 5` + "\n"
	config += `  meta_counts_items_total = 42` + "\n"
	config += `  items = [{` + "\n"
	config += `    name = "My Network"` + "\n"
	config += `    network_id = "L_12345"` + "\n"
	config += `    time_zone = "America/Los_Angeles"` + "\n"
	config += `    ai_enabled = true` + "\n"
	config += `    ai_last_enabled_at = "2026-02-05T14:49:07Z"` + "\n"
	config += `    busy_hour_minimize_changes_enabled = true` + "\n"
	config += `    busy_hour_schedule_mode = "automatic"` + "\n"
	config += `    busy_hour_schedule_automatic_end = "17:00"` + "\n"
	config += `    busy_hour_schedule_automatic_start = "08:00"` + "\n"
	config += `    busy_hour_schedule_manual_end = "15:00"` + "\n"
	config += `    busy_hour_schedule_manual_start = "10:00"` + "\n"
	config += `    channel_avoidance_enabled = true` + "\n"
	config += `    fra_enabled = false` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_radio_rrm_by_network" "test" {
			id = meraki_radio_rrm_by_network.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_radio_rrm_by_network.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
