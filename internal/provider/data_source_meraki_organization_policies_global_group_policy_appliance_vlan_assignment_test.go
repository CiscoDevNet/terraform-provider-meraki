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

func TestAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment(t *testing.T) {
	if os.Getenv("ORGANIZATION_POLICIES_GLOBAL") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION_POLICIES_GLOBAL")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentPrerequisitesConfig + testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_organization_policies_global_group_policy" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Terraform Test Policy"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["appliance"]
}
resource "meraki_appliance_vlans_settings" "test" {
  network_id    = meraki_network.test.id
  vlans_enabled = true
}
resource "meraki_appliance_vlan" "test" {
  network_id   = meraki_appliance_vlans_settings.test.network_id
  vlan_id      = "20"
  name         = "Terraform Test VLAN"
  subnet       = "192.168.20.0/24"
  appliance_ip = "192.168.20.1"
}
data "meraki_appliance_vlan" "test" {
  network_id = meraki_appliance_vlan.test.network_id
  id         = meraki_appliance_vlan.test.vlan_id
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentConfig() string {
	config := `resource "meraki_organization_policies_global_group_policy_appliance_vlan_assignment" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  policy_id = meraki_organization_policies_global_group_policy.test.id` + "\n"
	config += `  interface_id = data.meraki_appliance_vlan.test.interface_id` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_policies_global_group_policy_appliance_vlan_assignment" "test" {
			id = meraki_organization_policies_global_group_policy_appliance_vlan_assignment.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_policies_global_group_policy_appliance_vlan_assignment.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
