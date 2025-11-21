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

func TestAccMerakiSwitchPortSchedule(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "name", "Weekdays schedule"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_friday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_friday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_friday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_monday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_monday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_monday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_saturday_active", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_saturday_from", "00:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_saturday_to", "24:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_sunday_active", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_sunday_from", "00:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_sunday_to", "24:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_thursday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_thursday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_thursday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_tuesday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_tuesday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_tuesday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_wednesday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_wednesday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port_schedule.test", "port_schedule_wednesday_to", "17:00"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchPortSchedulePrerequisitesConfig+testAccMerakiSwitchPortScheduleConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchPortSchedulePrerequisitesConfig+testAccMerakiSwitchPortScheduleConfig_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_switch_port_schedule.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: merakiSwitchPortScheduleImportStateIdFunc("meraki_switch_port_schedule.test"),
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

func merakiSwitchPortScheduleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId,Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchPortSchedulePrerequisitesConfig = `
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

func testAccMerakiSwitchPortScheduleConfig_minimum() string {
	config := `resource "meraki_switch_port_schedule" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "Weekdays schedule"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchPortScheduleConfig_all() string {
	config := `resource "meraki_switch_port_schedule" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "Weekdays schedule"` + "\n"
	config += `  port_schedule_friday_active = true` + "\n"
	config += `  port_schedule_friday_from = "09:00"` + "\n"
	config += `  port_schedule_friday_to = "17:00"` + "\n"
	config += `  port_schedule_monday_active = true` + "\n"
	config += `  port_schedule_monday_from = "09:00"` + "\n"
	config += `  port_schedule_monday_to = "17:00"` + "\n"
	config += `  port_schedule_saturday_active = false` + "\n"
	config += `  port_schedule_saturday_from = "00:00"` + "\n"
	config += `  port_schedule_saturday_to = "24:00"` + "\n"
	config += `  port_schedule_sunday_active = false` + "\n"
	config += `  port_schedule_sunday_from = "00:00"` + "\n"
	config += `  port_schedule_sunday_to = "24:00"` + "\n"
	config += `  port_schedule_thursday_active = true` + "\n"
	config += `  port_schedule_thursday_from = "09:00"` + "\n"
	config += `  port_schedule_thursday_to = "17:00"` + "\n"
	config += `  port_schedule_tuesday_active = true` + "\n"
	config += `  port_schedule_tuesday_from = "09:00"` + "\n"
	config += `  port_schedule_tuesday_to = "17:00"` + "\n"
	config += `  port_schedule_wednesday_active = true` + "\n"
	config += `  port_schedule_wednesday_from = "09:00"` + "\n"
	config += `  port_schedule_wednesday_to = "17:00"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional



// End of section. //template:end testAccConfigAdditional
