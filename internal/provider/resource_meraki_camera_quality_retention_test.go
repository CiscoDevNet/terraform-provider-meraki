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

func TestAccMerakiCameraQualityRetention(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_camera_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_camera_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "audio_recording_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "motion_based_retention_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "motion_detector_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention.test", "restricted_bandwidth_mode_enabled", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiCameraQualityRetentionPrerequisitesConfig + testAccMerakiCameraQualityRetentionConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiCameraQualityRetentionPrerequisitesConfig + testAccMerakiCameraQualityRetentionConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_camera_quality_retention.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiCameraQualityRetentionImportStateIdFunc("meraki_camera_quality_retention.test"),
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

func merakiCameraQualityRetentionImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		Serial := primary.Attributes["serial"]

		return fmt.Sprintf("%s", Serial), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiCameraQualityRetentionPrerequisitesConfig = `
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

func testAccMerakiCameraQualityRetentionConfig_minimum() string {
	config := `resource "meraki_camera_quality_retention" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  audio_recording_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiCameraQualityRetentionConfig_all() string {
	config := `resource "meraki_camera_quality_retention" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  audio_recording_enabled = false` + "\n"
	config += `  motion_based_retention_enabled = false` + "\n"
	config += `  motion_detector_version = 2` + "\n"
	config += `  quality = "Standard"` + "\n"
	config += `  resolution = "1280x720"` + "\n"
	config += `  restricted_bandwidth_mode_enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
