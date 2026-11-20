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

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiOrganizationPoliciesGlobalFirewallRule(t *testing.T) {
	if os.Getenv("ORGANIZATION_POLICIES_GLOBAL") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION_POLICIES_GLOBAL")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "description", "This is rule 1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "name", "Allow developers"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "policy", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "destinations_criteria_address_ranges.0", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policies_global_firewall_rule.test", "sources_criteria_address_ranges.0", "1.1.1.1"))

	var steps []resource.TestStep
	var tfVersion *goversion.Version
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationPoliciesGlobalFirewallRulePrerequisitesConfig + testAccMerakiOrganizationPoliciesGlobalFirewallRuleConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	importStateVerifyIgnore := []string{}
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_organization_policies_global_firewall_rule.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiOrganizationPoliciesGlobalFirewallRuleImportStateIdFunc("meraki_organization_policies_global_firewall_rule.test"),
		ImportStateVerifyIgnore: importStateVerifyIgnore,
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			terraformVersionCapture{Version: &tfVersion},
		},
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiOrganizationPoliciesGlobalFirewallRuleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", OrganizationId, Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiOrganizationPoliciesGlobalFirewallRulePrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_organization_policies_global_firewall_ruleset" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Terraform Test Ruleset"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationPoliciesGlobalFirewallRuleConfig_minimum() string {
	config := `resource "meraki_organization_policies_global_firewall_rule" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  name = "Allow developers"` + "\n"
	config += `  policy = "deny"` + "\n"
	config += `  priority = 1` + "\n"
	config += `  ruleset_id = meraki_organization_policies_global_firewall_ruleset.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccMerakiOrganizationPoliciesGlobalFirewallRuleConfig_all() string {
	config := `resource "meraki_organization_policies_global_firewall_rule" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  description = "This is rule 1"` + "\n"
	config += `  enabled = true` + "\n"
	config += `  name = "Allow developers"` + "\n"
	config += `  policy = "deny"` + "\n"
	config += `  priority = 1` + "\n"
	config += `  ruleset_id = meraki_organization_policies_global_firewall_ruleset.test.id` + "\n"
	config += `  destinations_criteria_address_ranges = ["1.1.1.1"]` + "\n"
	config += `  destinations_match_criteria = ["addressRanges"]` + "\n"
	config += `  sources_criteria_address_ranges = ["1.1.1.1"]` + "\n"
	config += `  sources_match_criteria = ["addressRanges"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
