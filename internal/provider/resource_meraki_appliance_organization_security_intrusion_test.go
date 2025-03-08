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

func TestAccMerakiApplianceOrganizationSecurityIntrusion(t *testing.T) {
	if os.Getenv("APPLIANCE_ORGANIZATION_SECURITY_INTRUSION") == "" {
		t.Skip("skipping test, set environment variable APPLIANCE_ORGANIZATION_SECURITY_INTRUSION")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_organization_security_intrusion.test", "allowed_rules.0.message", "SQL sa login failed"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_organization_security_intrusion.test", "allowed_rules.0.rule_id", "meraki:intrusion/snort/GID/01/SID/688"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceOrganizationSecurityIntrusionPrerequisitesConfig + testAccMerakiApplianceOrganizationSecurityIntrusionConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceOrganizationSecurityIntrusionPrerequisitesConfig + testAccMerakiApplianceOrganizationSecurityIntrusionConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_organization_security_intrusion.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceOrganizationSecurityIntrusionImportStateIdFunc("meraki_appliance_organization_security_intrusion.test"),
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

func merakiApplianceOrganizationSecurityIntrusionImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]

		return fmt.Sprintf("%s", OrganizationId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceOrganizationSecurityIntrusionPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceOrganizationSecurityIntrusionConfig_minimum() string {
	config := `resource "meraki_appliance_organization_security_intrusion" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  allowed_rules = [{` + "\n"
	config += `    rule_id = "meraki:intrusion/snort/GID/01/SID/688"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceOrganizationSecurityIntrusionConfig_all() string {
	config := `resource "meraki_appliance_organization_security_intrusion" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  allowed_rules = [{` + "\n"
	config += `    message = "SQL sa login failed"` + "\n"
	config += `    rule_id = "meraki:intrusion/snort/GID/01/SID/688"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
