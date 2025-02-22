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

func TestAccDataSourceMerakiCameraQualityRetention(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_camera_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_camera_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "audio_recording_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "motion_based_retention_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "motion_detector_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_camera_quality_retention.test", "restricted_bandwidth_mode_enabled", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiCameraQualityRetentionPrerequisitesConfig + testAccDataSourceMerakiCameraQualityRetentionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiCameraQualityRetentionPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiCameraQualityRetentionConfig() string {
	config := `resource "meraki_camera_quality_retention" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  audio_recording_enabled = false` + "\n"
	config += `  motion_based_retention_enabled = false` + "\n"
	config += `  motion_detector_version = 2` + "\n"
	config += `  quality = "Standard"` + "\n"
	config += `  resolution = "1280x720"` + "\n"
	config += `  restricted_bandwidth_mode_enabled = false` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_camera_quality_retention" "test" {
			serial = tolist(meraki_network_device_claim.test.serials)[0]
			depends_on = [meraki_camera_quality_retention.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
