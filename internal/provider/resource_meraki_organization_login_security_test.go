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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiOrganizationLoginSecurity(t *testing.T) {
	if os.Getenv("ORGANIZATION") == "" {
        t.Skip("skipping test, set environment variable ORGANIZATION")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "account_lockout_attempts", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_account_lockout", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_different_passwords", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_idle_timeout", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_password_expiration", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_strong_passwords", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "enforce_two_factor_auth", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "idle_timeout_minutes", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "minimum_password_length", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "num_different_passwords", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_login_security.test", "password_expiration_days", "90"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationLoginSecurityPrerequisitesConfig+testAccMerakiOrganizationLoginSecurityConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationLoginSecurityPrerequisitesConfig+testAccMerakiOrganizationLoginSecurityConfig_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_organization_login_security.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: merakiOrganizationLoginSecurityImportStateIdFunc("meraki_organization_login_security.test"),
		ImportStateVerifyIgnore: []string{ "enforce_login_ip_ranges","api_authentication_ip_restrictions_for_keys_enabled", },
		Check: resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiOrganizationLoginSecurityImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]

		return fmt.Sprintf("%s", OrganizationId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiOrganizationLoginSecurityPrerequisitesConfig = `
resource "meraki_organization" "test" {
  name = "TF-Test1"
}

`
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationLoginSecurityConfig_minimum() string {
	config := `resource "meraki_organization_login_security" "test" {` + "\n"
	config += `  organization_id = meraki_organization.test.id` + "\n"
	config += `  account_lockout_attempts = 5` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationLoginSecurityConfig_all() string {
	config := `resource "meraki_organization_login_security" "test" {` + "\n"
	config += `  organization_id = meraki_organization.test.id` + "\n"
	config += `  account_lockout_attempts = 3` + "\n"
	config += `  enforce_account_lockout = true` + "\n"
	config += `  enforce_different_passwords = true` + "\n"
	config += `  enforce_idle_timeout = true` + "\n"
	config += `  enforce_password_expiration = true` + "\n"
	config += `  enforce_strong_passwords = true` + "\n"
	config += `  enforce_two_factor_auth = true` + "\n"
	config += `  idle_timeout_minutes = 30` + "\n"
	config += `  minimum_password_length = 12` + "\n"
	config += `  num_different_passwords = 3` + "\n"
	config += `  password_expiration_days = 90` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional



// End of section. //template:end testAccConfigAdditional
