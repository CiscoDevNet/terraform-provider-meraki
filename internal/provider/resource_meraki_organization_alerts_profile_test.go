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

func TestAccMerakiOrganizationAlertsProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "description", "WAN 1 high utilization"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "type", "wanUtilization"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "alert_condition_bit_rate_bps", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "alert_condition_duration", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "alert_condition_interface", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "alert_condition_window", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "recipients_emails.0", "admin@example.org"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organization_alerts_profile.test", "network_tags.0", "tag1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationAlertsProfilePrerequisitesConfig + testAccMerakiOrganizationAlertsProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationAlertsProfilePrerequisitesConfig + testAccMerakiOrganizationAlertsProfileConfig_all(),
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

const testAccMerakiOrganizationAlertsProfilePrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationAlertsProfileConfig_minimum() string {
	config := `resource "meraki_organization_alerts_profile" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	type = "wanUtilization"` + "\n"
	config += `	alert_condition_bit_rate_bps = 1000` + "\n"
	config += `	alert_condition_duration = 120` + "\n"
	config += `	alert_condition_interface = "wan2"` + "\n"
	config += `	alert_condition_window = 300` + "\n"
	config += `	recipients_emails = ["admin2@example.org"]` + "\n"
	config += `	network_tags = ["tag1"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationAlertsProfileConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll
