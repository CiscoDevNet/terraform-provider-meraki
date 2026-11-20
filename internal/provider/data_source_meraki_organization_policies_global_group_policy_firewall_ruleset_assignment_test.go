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

func TestAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignment(t *testing.T) {
	if os.Getenv("ORGANIZATION_POLICIES_GLOBAL") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION_POLICIES_GLOBAL")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_policies_global_group_policy_firewall_ruleset_assignment.test", "priority", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentPrerequisitesConfig + testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_organization_policies_global_firewall_ruleset" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Terraform Test Ruleset"
}
resource "meraki_organization_policies_global_group_policy" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Terraform Test Policy"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentConfig() string {
	config := `resource "meraki_organization_policies_global_group_policy_firewall_ruleset_assignment" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  policy_id = meraki_organization_policies_global_group_policy.test.id` + "\n"
	config += `  priority = 1` + "\n"
	config += `  ruleset_id = meraki_organization_policies_global_firewall_ruleset.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_policies_global_group_policy_firewall_ruleset_assignment" "test" {
			id = meraki_organization_policies_global_group_policy_firewall_ruleset_assignment.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_policies_global_group_policy_firewall_ruleset_assignment.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
