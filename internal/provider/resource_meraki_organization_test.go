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

func TestAccMerakiOrganization(t *testing.T) {
	if os.Getenv("ORGANIZATION") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization.test", "name", "My organization"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization.test", "management_details.0.name", "MSP ID"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization.test", "management_details.0.value", "123456"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_organization.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiOrganizationImportStateIdFunc("meraki_organization.test"),
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

func merakiOrganizationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s", id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationConfig_minimum() string {
	config := `resource "meraki_organization" "test" {` + "\n"
	config += `  name = "My organization"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationConfig_all() string {
	config := `resource "meraki_organization" "test" {` + "\n"
	config += `  name = "My organization"` + "\n"
	config += `  management_details = [{` + "\n"
	config += `    name = "MSP ID"` + "\n"
	config += `    value = "123456"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
