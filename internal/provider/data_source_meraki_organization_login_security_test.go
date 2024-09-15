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

func TestAccDataSourceMerakiOrganizationLoginSecurity(t *testing.T) {
	if os.Getenv("ORGANIZATION") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "account_lockout_attempts", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_account_lockout", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_different_passwords", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_idle_timeout", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_password_expiration", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_strong_passwords", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "enforce_two_factor_auth", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "idle_timeout_minutes", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "num_different_passwords", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_login_security.test", "password_expiration_days", "90"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationLoginSecurityPrerequisitesConfig + testAccDataSourceMerakiOrganizationLoginSecurityConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationLoginSecurityPrerequisitesConfig = `
resource "meraki_organization" "test" {
  name = "TF-Test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationLoginSecurityConfig() string {
	config := `resource "meraki_organization_login_security" "test" {` + "\n"
	config += `	organization_id = meraki_organization.test.id` + "\n"
	config += `	account_lockout_attempts = 3` + "\n"
	config += `	enforce_account_lockout = true` + "\n"
	config += `	enforce_different_passwords = true` + "\n"
	config += `	enforce_idle_timeout = true` + "\n"
	config += `	enforce_password_expiration = true` + "\n"
	config += `	enforce_strong_passwords = true` + "\n"
	config += `	enforce_two_factor_auth = true` + "\n"
	config += `	idle_timeout_minutes = 30` + "\n"
	config += `	num_different_passwords = 3` + "\n"
	config += `	password_expiration_days = 90` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_login_security" "test" {
			organization_id = meraki_organization.test.id
			depends_on = [meraki_organization_login_security.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
