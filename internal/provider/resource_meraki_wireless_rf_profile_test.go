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

func TestAccMerakiWirelessRFProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "band_selection_type", "ap"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "client_balancing_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "min_bitrate_type", "band"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "name", "Main Office"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "ap_band_settings_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "ap_band_settings_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "five_ghz_settings_channel_width", "auto"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "five_ghz_settings_max_power", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "five_ghz_settings_min_bitrate", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "five_ghz_settings_min_power", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "five_ghz_settings_rxsop", "-95"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_0_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_0_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_0_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_1_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_1_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_1_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_10_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_10_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_10_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_11_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_11_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_11_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_12_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_12_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_12_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_13_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_13_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_13_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_14_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_14_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_14_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_2_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_2_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_2_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_3_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_3_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_3_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_4_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_4_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_4_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_5_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_5_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_5_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_6_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_6_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_6_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_7_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_7_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_7_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_8_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_8_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_8_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_9_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_9_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "per_ssid_settings_9_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "transmission_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "two_four_ghz_settings_ax_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "two_four_ghz_settings_max_power", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "two_four_ghz_settings_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "two_four_ghz_settings_min_power", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "two_four_ghz_settings_rxsop", "-95"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "is_indoor_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_rf_profile.test", "is_outdoor_default", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessRFProfilePrerequisitesConfig+testAccMerakiWirelessRFProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessRFProfilePrerequisitesConfig+testAccMerakiWirelessRFProfileConfig_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_wireless_rf_profile.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: merakiWirelessRFProfileImportStateIdFunc("meraki_wireless_rf_profile.test"),
		ImportStateVerifyIgnore: []string{  },
		Check: resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessRFProfileImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId,Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessRFProfilePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}

`
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessRFProfileConfig_minimum() string {
	config := `resource "meraki_wireless_rf_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  band_selection_type = "ap"` + "\n"
	config += `  name = "Main Office"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessRFProfileConfig_all() string {
	config := `resource "meraki_wireless_rf_profile" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  band_selection_type = "ap"` + "\n"
	config += `  client_balancing_enabled = true` + "\n"
	config += `  min_bitrate_type = "band"` + "\n"
	config += `  name = "Main Office"` + "\n"
	config += `  ap_band_settings_band_operation_mode = "5ghz"` + "\n"
	config += `  ap_band_settings_band_steering_enabled = true` + "\n"
	config += `  ap_band_settings_bands_enabled = ["5"]` + "\n"
	config += `  five_ghz_settings_channel_width = "auto"` + "\n"
	config += `  five_ghz_settings_max_power = 30` + "\n"
	config += `  five_ghz_settings_min_bitrate = 12` + "\n"
	config += `  five_ghz_settings_min_power = 8` + "\n"
	config += `  five_ghz_settings_rxsop = -95` + "\n"
	config += `  five_ghz_settings_valid_auto_channels = [36]` + "\n"
	config += `  per_ssid_settings_0_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_0_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_0_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_0_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_1_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_1_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_1_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_1_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_10_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_10_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_10_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_10_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_11_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_11_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_11_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_11_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_12_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_12_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_12_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_12_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_13_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_13_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_13_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_13_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_14_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_14_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_14_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_14_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_2_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_2_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_2_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_2_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_3_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_3_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_3_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_3_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_4_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_4_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_4_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_4_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_5_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_5_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_5_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_5_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_6_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_6_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_6_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_6_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_7_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_7_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_7_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_7_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_8_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_8_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_8_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_8_bands_enabled = ["5"]` + "\n"
	config += `  per_ssid_settings_9_band_operation_mode = "5ghz"` + "\n"
	config += `  per_ssid_settings_9_band_steering_enabled = true` + "\n"
	config += `  per_ssid_settings_9_min_bitrate = 11` + "\n"
	config += `  per_ssid_settings_9_bands_enabled = ["5"]` + "\n"
	config += `  transmission_enabled = true` + "\n"
	config += `  two_four_ghz_settings_ax_enabled = true` + "\n"
	config += `  two_four_ghz_settings_max_power = 30` + "\n"
	config += `  two_four_ghz_settings_min_bitrate = 11` + "\n"
	config += `  two_four_ghz_settings_min_power = 5` + "\n"
	config += `  two_four_ghz_settings_rxsop = -95` + "\n"
	config += `  two_four_ghz_settings_valid_auto_channels = [1]` + "\n"
	config += `  is_indoor_default = false` + "\n"
	config += `  is_outdoor_default = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional



// End of section. //template:end testAccConfigAdditional
