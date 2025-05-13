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

func TestAccMerakiCameraQualityRetentionProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "audio_recording_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "cloud_archive_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "max_retention_days", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "motion_based_retention_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "motion_detector_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "name", "Sample quality retention profile"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "restricted_bandwidth_mode_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "smart_retention_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv12_mv22_mv72_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv12_mv22_mv72_resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv12_we_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv12_we_resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv13_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv13_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv13_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv13_m_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv21_mv71_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv21_mv71_resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv22_x_mv72_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv22_x_mv72_x_resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_m_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv23_x_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv32_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv32_resolution", "1080x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv33_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv33_resolution", "1080x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv33_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv33_m_resolution", "1080x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv52_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv52_resolution", "1280x720"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv53_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv53_x_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_m_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv63_x_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_m_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv73_x_resolution", "1920x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv84_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv84_x_resolution", "1440x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_resolution", "1080x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_m_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_m_resolution", "1080x1080"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_x_quality", "Standard"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_camera_quality_retention_profile.test", "video_settings_mv93_x_resolution", "1080x1080"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiCameraQualityRetentionProfilePrerequisitesConfig + testAccMerakiCameraQualityRetentionProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiCameraQualityRetentionProfilePrerequisitesConfig + testAccMerakiCameraQualityRetentionProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_camera_quality_retention_profile.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiCameraQualityRetentionProfileImportStateIdFunc("meraki_camera_quality_retention_profile.test"),
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

func merakiCameraQualityRetentionProfileImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId, Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiCameraQualityRetentionProfilePrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiCameraQualityRetentionProfileConfig_minimum() string {
	config := `resource "meraki_camera_quality_retention_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "Sample quality retention profile"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiCameraQualityRetentionProfileConfig_all() string {
	config := `resource "meraki_camera_quality_retention_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  audio_recording_enabled = false` + "\n"
	config += `  cloud_archive_enabled = false` + "\n"
	config += `  max_retention_days = 1` + "\n"
	config += `  motion_based_retention_enabled = false` + "\n"
	config += `  motion_detector_version = 2` + "\n"
	config += `  name = "Sample quality retention profile"` + "\n"
	config += `  restricted_bandwidth_mode_enabled = false` + "\n"
	config += `  smart_retention_enabled = false` + "\n"
	config += `  video_settings_mv12_mv22_mv72_quality = "Standard"` + "\n"
	config += `  video_settings_mv12_mv22_mv72_resolution = "1280x720"` + "\n"
	config += `  video_settings_mv12_we_quality = "Standard"` + "\n"
	config += `  video_settings_mv12_we_resolution = "1280x720"` + "\n"
	config += `  video_settings_mv13_quality = "Standard"` + "\n"
	config += `  video_settings_mv13_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv13_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv13_m_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv21_mv71_quality = "Standard"` + "\n"
	config += `  video_settings_mv21_mv71_resolution = "1280x720"` + "\n"
	config += `  video_settings_mv22_x_mv72_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv22_x_mv72_x_resolution = "1280x720"` + "\n"
	config += `  video_settings_mv23_quality = "Standard"` + "\n"
	config += `  video_settings_mv23_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv23_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv23_m_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv23_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv23_x_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv32_quality = "Standard"` + "\n"
	config += `  video_settings_mv32_resolution = "1080x1080"` + "\n"
	config += `  video_settings_mv33_quality = "Standard"` + "\n"
	config += `  video_settings_mv33_resolution = "1080x1080"` + "\n"
	config += `  video_settings_mv33_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv33_m_resolution = "1080x1080"` + "\n"
	config += `  video_settings_mv52_quality = "Standard"` + "\n"
	config += `  video_settings_mv52_resolution = "1280x720"` + "\n"
	config += `  video_settings_mv53_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv53_x_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv63_quality = "Standard"` + "\n"
	config += `  video_settings_mv63_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv63_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv63_m_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv63_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv63_x_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv73_quality = "Standard"` + "\n"
	config += `  video_settings_mv73_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv73_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv73_m_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv73_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv73_x_resolution = "1920x1080"` + "\n"
	config += `  video_settings_mv84_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv84_x_resolution = "1440x1080"` + "\n"
	config += `  video_settings_mv93_quality = "Standard"` + "\n"
	config += `  video_settings_mv93_resolution = "1080x1080"` + "\n"
	config += `  video_settings_mv93_m_quality = "Standard"` + "\n"
	config += `  video_settings_mv93_m_resolution = "1080x1080"` + "\n"
	config += `  video_settings_mv93_x_quality = "Standard"` + "\n"
	config += `  video_settings_mv93_x_resolution = "1080x1080"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
