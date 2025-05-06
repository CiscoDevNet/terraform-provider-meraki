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

func TestAccDataSourceMerakiNetworkAlertsSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "default_destinations_all_admins", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "default_destinations_snmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "default_destinations_emails.0", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "muting_by_port_schedules_enabled", "true"))
	// disable alerts checks as default alerts are returned as well
	// checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "alerts.0.enabled", "true"))
	// checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "alerts.0.type", "gatewayDown"))
	// checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "alerts.0.alert_destinations_all_admins", "false"))
	// checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "alerts.0.alert_destinations_snmp", "false"))
	// checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_alerts_settings.test", "alerts.0.alert_destinations_emails.0", "miles@meraki.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkAlertsSettingsPrerequisitesConfig + testAccDataSourceMerakiNetworkAlertsSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkAlertsSettingsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkAlertsSettingsConfig() string {
	config := `resource "meraki_network_alerts_settings" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  default_destinations_all_admins = true` + "\n"
	config += `  default_destinations_snmp = true` + "\n"
	config += `  default_destinations_emails = ["miles@meraki.com"]` + "\n"
	config += `  muting_by_port_schedules_enabled = true` + "\n"
	config += `  alerts = [{` + "\n"
	config += `    enabled = true` + "\n"
	config += `    type = "gatewayDown"` + "\n"
	config += `    alert_destinations_all_admins = false` + "\n"
	config += `    alert_destinations_snmp = false` + "\n"
	config += `    alert_destinations_emails = ["miles@meraki.com"]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_alerts_settings" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_network_alerts_settings.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
