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

func TestAccDataSourceMerakiWirelessSSIDSchedules(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_schedules.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_schedules.test", "ranges.0.end_day", "Tuesday"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_schedules.test", "ranges.0.end_time", "05:00:00"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_schedules.test", "ranges.0.start_day", "Tuesday"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_schedules.test", "ranges.0.start_time", "01:00:00"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDSchedulesPrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDSchedulesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDSchedulesPrerequisitesConfig = `
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDSchedulesConfig() string {
	config := `resource "meraki_wireless_ssid_schedules" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	ranges = [{` + "\n"
	config += `		end_day = "Tuesday"` + "\n"
	config += `		end_time = "05:00:00"` + "\n"
	config += `		start_day = "Tuesday"` + "\n"
	config += `		start_time = "01:00:00"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_schedules" "test" {
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_schedules.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
