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

func TestAccMerakiSensorAlertsProfiles(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_sensor_1_serial") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_sensor_1_serial")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSensorAlertsProfilesPrerequisitesConfig+testAccMerakiSensorAlertsProfilesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSensorAlertsProfilesPrerequisitesConfig+testAccMerakiSensorAlertsProfilesConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_sensor_alerts_profiles.test",
		ImportState: true,
		ImportStateIdFunc: merakiSensorAlertsProfilesImportStateIdFunc("meraki_sensor_alerts_profiles.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSensorAlertsProfilesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s,%s", OrganizationId,NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSensorAlertsProfilesPrerequisitesConfig = `
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

func testAccMerakiSensorAlertsProfilesConfig_minimum() string {
	config := `resource "meraki_sensor_alerts_profiles" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  name = "My Sensor Alert Profile"` + "\n"
	config += `  conditions = [{` + "\n"
	config += `    direction = "below"` + "\n"
	config += `    metric = "temperature"` + "\n"
	config += `    threshold_temperature_celsius = 10.5` + "\n"
	config += `  }]` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSensorAlertsProfilesConfig_all() string {
	config := `resource "meraki_sensor_alerts_profiles" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  include_sensor_url = true` + "\n"
	config += `  message = "Check with Miles on what to do."` + "\n"
	config += `  name = "My Sensor Alert Profile"` + "\n"
	config += `  conditions = [{` + "\n"
	config += `    direction = "above"` + "\n"
	config += `    duration = 60` + "\n"
	config += `    metric = "temperature"` + "\n"
	config += `    threshold_temperature_celsius = 20.5` + "\n"
	config += `  }]` + "\n"
	config += `  serials = tolist(meraki_network_device_claim.test.serials)` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
