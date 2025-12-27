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

func TestAccDataSourceMerakiOrganizationSAMLIdP(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_saml_idp.test", "slo_logout_url", "https://somewhere.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_saml_idp.test", "sso_login_url", "https://onelogin.com/trust/saml2/http-post/sso/3de5f942-e7b8-4cb9-94e3-85828111158b"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_saml_idp.test", "x509cert_sha1_fingerprint", "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:CA"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationSAMLIdPPrerequisitesConfig + testAccDataSourceMerakiOrganizationSAMLIdPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationSAMLIdPPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationSAMLIdPConfig() string {
	config := `resource "meraki_organization_saml_idp" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  slo_logout_url = "https://somewhere.com"` + "\n"
	config += `  sso_login_url = "https://onelogin.com/trust/saml2/http-post/sso/3de5f942-e7b8-4cb9-94e3-85828111158b"` + "\n"
	config += `  x509cert_sha1_fingerprint = "00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:CA"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_saml_idp" "test" {
			id = meraki_organization_saml_idp.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_saml_idp.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
