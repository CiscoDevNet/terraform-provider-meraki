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

func TestAccMerakiOrganizationAdmin(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "authentication_method", "Email"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "email", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "name", "Miles Meraki"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "org_access", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "networks.0.access", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "tags.0.access", "read-only"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_admin.test", "tags.0.tag", "west"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationAdminPrerequisitesConfig + testAccMerakiOrganizationAdminConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationAdminPrerequisitesConfig + testAccMerakiOrganizationAdminConfig_all(),
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

const testAccMerakiOrganizationAdminPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationAdminConfig_minimum() string {
	config := `resource "meraki_organization_admin" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	email = "miles@meraki.com"` + "\n"
	config += `	name = "Miles Meraki"` + "\n"
	config += `	org_access = "full"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationAdminConfig_all() string {
	config := `resource "meraki_organization_admin" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	authentication_method = "Email"` + "\n"
	config += `	email = "miles@meraki.com"` + "\n"
	config += `	name = "Miles Meraki"` + "\n"
	config += `	org_access = "none"` + "\n"
	config += `	networks = [{` + "\n"
	config += `		access = "full"` + "\n"
	config += `		id = meraki_network.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	tags = [{` + "\n"
	config += `		access = "read-only"` + "\n"
	config += `		tag = "west"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
