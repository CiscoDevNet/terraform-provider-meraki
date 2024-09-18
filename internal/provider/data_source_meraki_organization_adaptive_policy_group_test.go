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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiOrganizationAdaptivePolicyGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_group.test", "description", "Group of XYZ Corp Employees"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_group.test", "name", "Employee Group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_group.test", "sgt", "1000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationAdaptivePolicyGroupPrerequisitesConfig + testAccDataSourceMerakiOrganizationAdaptivePolicyGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceMerakiOrganizationAdaptivePolicyGroupPrerequisitesConfig + testAccNamedDataSourceMerakiOrganizationAdaptivePolicyGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationAdaptivePolicyGroupPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_organization_policy_object" "test" {
  organization_id = data.meraki_organization.test.id
  category        = "adaptivePolicy"
  cidr            = "10.0.0.0/24"
  name            = "Web Servers - Datacenter 10"
  type            = "adaptivePolicyIpv4Cidr"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationAdaptivePolicyGroupConfig() string {
	config := `resource "meraki_organization_adaptive_policy_group" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	description = "Group of XYZ Corp Employees"` + "\n"
	config += `	name = "Employee Group"` + "\n"
	config += `	sgt = 1000` + "\n"
	config += `	policy_objects = [{` + "\n"
	config += `		id = meraki_organization_policy_object.test.id` + "\n"
	config += `		name = meraki_organization_policy_object.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_adaptive_policy_group" "test" {
			id = meraki_organization_adaptive_policy_group.test.id
			organization_id = data.meraki_organization.test.id
		}
	`
	return config
}

func testAccNamedDataSourceMerakiOrganizationAdaptivePolicyGroupConfig() string {
	config := `resource "meraki_organization_adaptive_policy_group" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	description = "Group of XYZ Corp Employees"` + "\n"
	config += `	name = "Employee Group"` + "\n"
	config += `	sgt = 1000` + "\n"
	config += `	policy_objects = [{` + "\n"
	config += `		id = meraki_organization_policy_object.test.id` + "\n"
	config += `		name = meraki_organization_policy_object.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_adaptive_policy_group" "test" {
			name = meraki_organization_adaptive_policy_group.test.name
			organization_id = data.meraki_organization.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig