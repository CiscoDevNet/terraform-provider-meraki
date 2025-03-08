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

func TestAccMerakiOrganizationSAMLRole(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_role.test", "org_access", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_role.test", "role", "myrole"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_role.test", "networks.0.access", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_role.test", "tags.0.access", "read-only"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_role.test", "tags.0.tag", "west"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationSAMLRolePrerequisitesConfig + testAccMerakiOrganizationSAMLRoleConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationSAMLRolePrerequisitesConfig + testAccMerakiOrganizationSAMLRoleConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_organization_saml_role.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiOrganizationSAMLRoleImportStateIdFunc("meraki_organization_saml_role.test"),
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

func merakiOrganizationSAMLRoleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", OrganizationId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiOrganizationSAMLRolePrerequisitesConfig = `
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

func testAccMerakiOrganizationSAMLRoleConfig_minimum() string {
	config := `resource "meraki_organization_saml_role" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  org_access = "read-only"` + "\n"
	config += `  role = "myrole"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationSAMLRoleConfig_all() string {
	config := `resource "meraki_organization_saml_role" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  org_access = "none"` + "\n"
	config += `  role = "myrole"` + "\n"
	config += `  networks = [{` + "\n"
	config += `    access = "full"` + "\n"
	config += `    id = meraki_network.test.id` + "\n"
	config += `  }]` + "\n"
	config += `  tags = [{` + "\n"
	config += `    access = "read-only"` + "\n"
	config += `    tag = "west"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
