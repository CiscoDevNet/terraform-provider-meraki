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

func TestAccMerakiApplianceRFProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "name", "MX RF Profile"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "five_ghz_settings_ax_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "five_ghz_settings_min_bitrate", "48"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_1_band_operation_mode", "dual"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_1_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_2_band_operation_mode", "dual"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_2_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_3_band_operation_mode", "dual"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_3_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_4_band_operation_mode", "dual"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "per_ssid_settings_4_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "two_four_ghz_settings_ax_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_rf_profile.test", "two_four_ghz_settings_min_bitrate", "12"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceRFProfilePrerequisitesConfig + testAccMerakiApplianceRFProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceRFProfilePrerequisitesConfig + testAccMerakiApplianceRFProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_rf_profile.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceRFProfileImportStateIdFunc("meraki_appliance_rf_profile.test"),
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

func merakiApplianceRFProfileImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceRFProfilePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_appliance_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_appliance_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceRFProfileConfig_minimum() string {
	config := `resource "meraki_appliance_rf_profile" "test" {` + "\n"
	config += `	network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `	name = "MX RF Profile"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceRFProfileConfig_all() string {
	config := `resource "meraki_appliance_rf_profile" "test" {` + "\n"
	config += `	network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `	name = "MX RF Profile"` + "\n"
	config += `	five_ghz_settings_ax_enabled = true` + "\n"
	config += `	five_ghz_settings_min_bitrate = 48` + "\n"
	config += `	per_ssid_settings_1_band_operation_mode = "dual"` + "\n"
	config += `	per_ssid_settings_1_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings_2_band_operation_mode = "dual"` + "\n"
	config += `	per_ssid_settings_2_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings_3_band_operation_mode = "dual"` + "\n"
	config += `	per_ssid_settings_3_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings_4_band_operation_mode = "dual"` + "\n"
	config += `	per_ssid_settings_4_band_steering_enabled = true` + "\n"
	config += `	two_four_ghz_settings_ax_enabled = true` + "\n"
	config += `	two_four_ghz_settings_min_bitrate = 12` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll