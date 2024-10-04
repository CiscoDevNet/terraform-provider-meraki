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

func TestAccDataSourceMerakiApplianceOrganizationSecurityIntrusion(t *testing.T) {
	if os.Getenv("APPLIANCE_ORGANIZATION_SECURITY_INTRUSION") == "" {
		t.Skip("skipping test, set environment variable APPLIANCE_ORGANIZATION_SECURITY_INTRUSION")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_organization_security_intrusion.test", "allowed_rules.0.message", "SQL sa login failed"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_organization_security_intrusion.test", "allowed_rules.0.rule_id", "meraki:intrusion/snort/GID/01/SID/688"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceOrganizationSecurityIntrusionPrerequisitesConfig + testAccDataSourceMerakiApplianceOrganizationSecurityIntrusionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceOrganizationSecurityIntrusionPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceOrganizationSecurityIntrusionConfig() string {
	config := `resource "meraki_appliance_organization_security_intrusion" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	allowed_rules = [{` + "\n"
	config += `		message = "SQL sa login failed"` + "\n"
	config += `		rule_id = "meraki:intrusion/snort/GID/01/SID/688"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_organization_security_intrusion" "test" {
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_appliance_organization_security_intrusion.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
