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

func TestAccDataSourceMerakiOrganizationSNMP(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_snmp.test", "v2c_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_snmp.test", "v3_auth_mode", "SHA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_snmp.test", "v3_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_snmp.test", "v3_priv_mode", "AES128"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationSNMPPrerequisitesConfig + testAccDataSourceMerakiOrganizationSNMPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationSNMPPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationSNMPConfig() string {
	config := `resource "meraki_organization_snmp" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  v2c_enabled = false` + "\n"
	config += `  v3_auth_mode = "SHA"` + "\n"
	config += `  v3_auth_pass = "password"` + "\n"
	config += `  v3_enabled = true` + "\n"
	config += `  v3_priv_mode = "AES128"` + "\n"
	config += `  v3_priv_pass = "password"` + "\n"
	config += `  peer_ips = ["123.123.123.1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_snmp" "test" {
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_snmp.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
