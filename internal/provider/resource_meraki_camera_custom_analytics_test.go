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

func TestAccMerakiCameraCustomAnalytics(t *testing.T) {
	if os.Getenv("CAMERA_CUSTOM_ANALYTICS") == "" {
		t.Skip("skipping test, set environment variable CAMERA_CUSTOM_ANALYTICS")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_camera_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_camera_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_custom_analytics.test", "artifact_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_custom_analytics.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_custom_analytics.test", "parameters.0.name", "detection_threshold"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_custom_analytics.test", "parameters.0.value", "0.5"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiCameraCustomAnalyticsPrerequisitesConfig + testAccMerakiCameraCustomAnalyticsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiCameraCustomAnalyticsPrerequisitesConfig + testAccMerakiCameraCustomAnalyticsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_camera_custom_analytics.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiCameraCustomAnalyticsImportStateIdFunc("meraki_camera_custom_analytics.test"),
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

func merakiCameraCustomAnalyticsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		Serial := primary.Attributes["serial"]

		return fmt.Sprintf("%s", Serial), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiCameraCustomAnalyticsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_camera_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_camera_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiCameraCustomAnalyticsConfig_minimum() string {
	config := `resource "meraki_camera_custom_analytics" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiCameraCustomAnalyticsConfig_all() string {
	config := `resource "meraki_camera_custom_analytics" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  artifact_id = "1"` + "\n"
	config += `  enabled = true` + "\n"
	config += `  parameters = [{` + "\n"
	config += `    name = "detection_threshold"` + "\n"
	config += `    value = "0.5"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
