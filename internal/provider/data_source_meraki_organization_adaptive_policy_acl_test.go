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

func TestAccDataSourceMerakiOrganizationAdaptivePolicyACL(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "description", "Blocks sensitive web traffic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "ip_version", "ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "name", "Block sensitive web traffic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "rules.0.dst_port", "22-30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "rules.0.policy", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_adaptive_policy_acl.test", "rules.0.src_port", "1,33"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationAdaptivePolicyACLPrerequisitesConfig + testAccDataSourceMerakiOrganizationAdaptivePolicyACLConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationAdaptivePolicyACLPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationAdaptivePolicyACLConfig() string {
	config := `resource "meraki_organization_adaptive_policy_acl" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  description = "Blocks sensitive web traffic"` + "\n"
	config += `  ip_version = "ipv6"` + "\n"
	config += `  name = "Block sensitive web traffic"` + "\n"
	config += `  rules = [{` + "\n"
	config += `    dst_port = "22-30"` + "\n"
	config += `    policy = "deny"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    src_port = "1,33"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_adaptive_policy_acl" "test" {
			id = meraki_organization_adaptive_policy_acl.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_adaptive_policy_acl.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiOrganizationAdaptivePolicyACLConfig() string {
	config := `resource "meraki_organization_adaptive_policy_acl" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  description = "Blocks sensitive web traffic"` + "\n"
	config += `  ip_version = "ipv6"` + "\n"
	config += `  name = "Block sensitive web traffic"` + "\n"
	config += `  rules = [{` + "\n"
	config += `    dst_port = "22-30"` + "\n"
	config += `    policy = "deny"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    src_port = "1,33"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_adaptive_policy_acl" "test" {
			name = meraki_organization_adaptive_policy_acl.test.name
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_adaptive_policy_acl.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
