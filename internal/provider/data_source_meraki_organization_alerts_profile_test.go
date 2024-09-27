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

func TestAccDataSourceMerakiOrganizationAlertsProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "description", "WAN 1 high utilization"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "type", "wanUtilization"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "alert_condition_bit_rate_bps", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "alert_condition_duration", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "alert_condition_interface", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_alerts_profile.test", "alert_condition_window", "600"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationAlertsProfilePrerequisitesConfig + testAccDataSourceMerakiOrganizationAlertsProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationAlertsProfilePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationAlertsProfileConfig() string {
	config := `resource "meraki_organization_alerts_profile" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	description = "WAN 1 high utilization"` + "\n"
	config += `	type = "wanUtilization"` + "\n"
	config += `	alert_condition_bit_rate_bps = 10000` + "\n"
	config += `	alert_condition_duration = 60` + "\n"
	config += `	alert_condition_interface = "wan1"` + "\n"
	config += `	alert_condition_window = 600` + "\n"
	config += `	recipients_emails = ["admin@example.org"]` + "\n"
	config += `	network_tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_alerts_profile" "test" {
			id = meraki_organization_alerts_profile.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_alerts_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
