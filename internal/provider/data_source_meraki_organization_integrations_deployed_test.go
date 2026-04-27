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

func TestAccDataSourceMerakiOrganizationIntegrationsDeployed(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "meta_counts_items_remaining", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "meta_counts_items_total", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "items.0.id", "98765"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "items.0.name", "OAuth Application"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "items.0.provider", "partner"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "items.0.type", "OAuth"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_integrations_deployed.test", "items.0.tags.0", "Wayfinding"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationIntegrationsDeployedPrerequisitesConfig + testAccDataSourceMerakiOrganizationIntegrationsDeployedConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationIntegrationsDeployedPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationIntegrationsDeployedConfig() string {
	return `
		data "meraki_organization_integrations_deployed" "test" {
			organization_id = data.meraki_organization.test.id
		}
	`
}

// End of section. //template:end testAccDataSourceConfig
