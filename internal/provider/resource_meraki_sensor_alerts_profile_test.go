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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiSensorAlertsProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_sensor_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_sensor_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "name", "My Sensor Alert Profile"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "recipients_emails.0", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "recipients_sms_numbers.0", "+15555555555"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "conditions.0.direction", "above"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "conditions.0.duration", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "conditions.0.metric", "temperature"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_sensor_alerts_profile.test", "conditions.0.threshold_temperature_celsius", "20.5"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSensorAlertsProfilePrerequisitesConfig + testAccMerakiSensorAlertsProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSensorAlertsProfilePrerequisitesConfig + testAccMerakiSensorAlertsProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_sensor_alerts_profile.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiSensorAlertsProfileImportStateIdFunc("meraki_sensor_alerts_profile.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSensorAlertsProfileImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSensorAlertsProfilePrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSensorAlertsProfileConfig_minimum() string {
	config := `resource "meraki_sensor_alerts_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "My Sensor Alert Profile"` + "\n"
	config += `  conditions = [{` + "\n"
	config += `    direction = "below"` + "\n"
	config += `    metric = "temperature"` + "\n"
	config += `    threshold_temperature_celsius = 10.5` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSensorAlertsProfileConfig_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll
