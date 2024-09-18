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

func TestAccMerakiOrganizationPolicyObject(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policy_object.test", "category", "network"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policy_object.test", "cidr", "10.0.0.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policy_object.test", "name", "Web Servers - Datacenter 10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_policy_object.test", "type", "cidr"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationPolicyObjectPrerequisitesConfig + testAccMerakiOrganizationPolicyObjectConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationPolicyObjectPrerequisitesConfig + testAccMerakiOrganizationPolicyObjectConfig_all(),
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

const testAccMerakiOrganizationPolicyObjectPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationPolicyObjectConfig_minimum() string {
	config := `resource "meraki_organization_policy_object" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	category = "network"` + "\n"
	config += `	cidr = "10.0.1.0/24"` + "\n"
	config += `	name = "Web Servers - Datacenter 10"` + "\n"
	config += `	type = "cidr"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationPolicyObjectConfig_all() string {
	config := `resource "meraki_organization_policy_object" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	category = "network"` + "\n"
	config += `	cidr = "10.0.0.0/24"` + "\n"
	config += `	name = "Web Servers - Datacenter 10"` + "\n"
	config += `	type = "cidr"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
