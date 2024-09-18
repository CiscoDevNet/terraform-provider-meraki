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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiOrganizationAdaptivePolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_adaptive_policy.test", "last_entry_rule", "allow"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationAdaptivePolicyPrerequisitesConfig + testAccMerakiOrganizationAdaptivePolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationAdaptivePolicyPrerequisitesConfig + testAccMerakiOrganizationAdaptivePolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiOrganizationAdaptivePolicyPrerequisitesConfig = `
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
resource "meraki_organization_adaptive_policy_group" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Employee Group"
  sgt             = 1000
  policy_objects = [
    {
      id   = meraki_organization_policy_object.test.id
      name = meraki_organization_policy_object.test.name
    }
  ]
}
resource "meraki_organization_adaptive_policy_acl" "test" {
  organization_id = data.meraki_organization.test.id
  ip_version      = "ipv4"
  name            = "Block sensitive web traffic"
  rules = [
    {
      dst_port = "80"
      policy   = "deny"
      protocol = "tcp"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationAdaptivePolicyConfig_minimum() string {
	config := `resource "meraki_organization_adaptive_policy" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	destination_group_id = meraki_organization_adaptive_policy_group.test.id` + "\n"
	config += `	destination_group_name = meraki_organization_adaptive_policy_group.test.name` + "\n"
	config += `	destination_group_sgt = meraki_organization_adaptive_policy_group.test.sgt` + "\n"
	config += `	source_group_id = meraki_organization_adaptive_policy_group.test.id` + "\n"
	config += `	source_group_name = meraki_organization_adaptive_policy_group.test.name` + "\n"
	config += `	source_group_sgt = meraki_organization_adaptive_policy_group.test.sgt` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationAdaptivePolicyConfig_all() string {
	config := `resource "meraki_organization_adaptive_policy" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	last_entry_rule = "allow"` + "\n"
	config += `	destination_group_id = meraki_organization_adaptive_policy_group.test.id` + "\n"
	config += `	destination_group_name = meraki_organization_adaptive_policy_group.test.name` + "\n"
	config += `	destination_group_sgt = meraki_organization_adaptive_policy_group.test.sgt` + "\n"
	config += `	source_group_id = meraki_organization_adaptive_policy_group.test.id` + "\n"
	config += `	source_group_name = meraki_organization_adaptive_policy_group.test.name` + "\n"
	config += `	source_group_sgt = meraki_organization_adaptive_policy_group.test.sgt` + "\n"
	config += `	acls = [{` + "\n"
	config += `		id = meraki_organization_adaptive_policy_acl.test.id` + "\n"
	config += `		name = meraki_organization_adaptive_policy_acl.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll