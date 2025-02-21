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

func TestAccMerakiSwitchOrganizationPortsProfilesAutomation(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_3_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_3_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_organization_ports_profiles_automation.test", "description", "A full length description of the automation."))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_organization_ports_profiles_automation.test", "name", "Automation 1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_organization_ports_profiles_automation.test", "rules.0.priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_organization_ports_profiles_automation.test", "rules.0.conditions.0.attribute", "LLDP system description"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_organization_ports_profiles_automation.test", "rules.0.conditions.0.values.0", "Meraki MR*"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchOrganizationPortsProfilesAutomationPrerequisitesConfig + testAccMerakiSwitchOrganizationPortsProfilesAutomationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchOrganizationPortsProfilesAutomationPrerequisitesConfig + testAccMerakiSwitchOrganizationPortsProfilesAutomationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_switch_organization_ports_profiles_automation.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiSwitchOrganizationPortsProfilesAutomationImportStateIdFunc("meraki_switch_organization_ports_profiles_automation.test"),
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

func merakiSwitchOrganizationPortsProfilesAutomationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", OrganizationId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchOrganizationPortsProfilesAutomationPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_3_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_switch_3_serial]
}
resource "meraki_switch_organization_ports_profile" "test" {
  organization_id = data.meraki_organization.test.id
  name = "Phone"
  port_type = "access"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchOrganizationPortsProfilesAutomationConfig_minimum() string {
	config := `resource "meraki_switch_organization_ports_profiles_automation" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  name = "Automation 1"` + "\n"
	config += `  rules = [{` + "\n"
	config += `    priority = 1` + "\n"
	config += `    profile_id = resource.meraki_switch_organization_ports_profile.test.id` + "\n"
	config += `    profile_name = resource.meraki_switch_organization_ports_profile.test.name` + "\n"
	config += `    conditions = [{` + "\n"
	config += `    attribute = "LLDP system description"` + "\n"
	config += `    values = ["Meraki MR*"]` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchOrganizationPortsProfilesAutomationConfig_all() string {
	config := `resource "meraki_switch_organization_ports_profiles_automation" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  description = "A full length description of the automation."` + "\n"
	config += `  name = "Automation 1"` + "\n"
	config += `  rules = [{` + "\n"
	config += `    priority = 1` + "\n"
	config += `    profile_id = resource.meraki_switch_organization_ports_profile.test.id` + "\n"
	config += `    profile_name = resource.meraki_switch_organization_ports_profile.test.name` + "\n"
	config += `    conditions = [{` + "\n"
	config += `      attribute = "LLDP system description"` + "\n"
	config += `      values = ["Meraki MR*"]` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
