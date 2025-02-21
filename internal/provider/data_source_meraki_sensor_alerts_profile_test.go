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

func TestAccDataSourceMerakiSensorAlertsProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_sensor_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_sensor_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "name", "My Sensor Alert Profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "recipients_emails.0", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "recipients_sms_numbers.0", "+15555555555"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "conditions.0.direction", "above"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "conditions.0.duration", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "conditions.0.metric", "temperature"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sensor_alerts_profile.test", "conditions.0.threshold_temperature_celsius", "20.5"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSensorAlertsProfilePrerequisitesConfig + testAccDataSourceMerakiSensorAlertsProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSensorAlertsProfilePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_sensor_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_sensor_1_serial]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSensorAlertsProfileConfig() string {
	config := `resource "meraki_sensor_alerts_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "My Sensor Alert Profile"` + "\n"
	config += `  recipients_emails = ["miles@meraki.com"]` + "\n"
	config += `  recipients_sms_numbers = ["+15555555555"]` + "\n"
	config += `  conditions = [{` + "\n"
	config += `    direction = "above"` + "\n"
	config += `    duration = 60` + "\n"
	config += `    metric = "temperature"` + "\n"
	config += `    threshold_temperature_celsius = 20.5` + "\n"
	config += `  }]` + "\n"
	config += `  serials = tolist(meraki_network_device_claim.test.serials)` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_sensor_alerts_profile" "test" {
			id = meraki_sensor_alerts_profile.test.id
			network_id = meraki_network.test.id
			depends_on = [meraki_sensor_alerts_profile.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiSensorAlertsProfileConfig() string {
	config := `resource "meraki_sensor_alerts_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "My Sensor Alert Profile"` + "\n"
	config += `  recipients_emails = ["miles@meraki.com"]` + "\n"
	config += `  recipients_sms_numbers = ["+15555555555"]` + "\n"
	config += `  conditions = [{` + "\n"
	config += `    direction = "above"` + "\n"
	config += `    duration = 60` + "\n"
	config += `    metric = "temperature"` + "\n"
	config += `    threshold_temperature_celsius = 20.5` + "\n"
	config += `  }]` + "\n"
	config += `  serials = tolist(meraki_network_device_claim.test.serials)` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_sensor_alerts_profile" "test" {
			name = meraki_sensor_alerts_profile.test.name
			network_id = meraki_network.test.id
			depends_on = [meraki_sensor_alerts_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
