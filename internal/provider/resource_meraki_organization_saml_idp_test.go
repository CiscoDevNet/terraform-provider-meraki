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

func TestAccMerakiOrganizationSAMLIdP(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_idp.test", "slo_logout_url", "https://somewhere.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_saml_idp.test", "x509cert_sha1_fingerprint", "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:AA"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationSAMLIdPPrerequisitesConfig + testAccMerakiOrganizationSAMLIdPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationSAMLIdPPrerequisitesConfig + testAccMerakiOrganizationSAMLIdPConfig_all(),
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

const testAccMerakiOrganizationSAMLIdPPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationSAMLIdPConfig_minimum() string {
	config := `resource "meraki_organization_saml_idp" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:AA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationSAMLIdPConfig_all() string {
	config := `resource "meraki_organization_saml_idp" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	slo_logout_url = "https://somewhere.com"` + "\n"
	config += `	x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:AA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
