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

func TestAccDataSourceMerakiWirelessRFProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "band_selection_type", "ap"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "client_balancing_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "min_bitrate_type", "band"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "name", "Main Office"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "ap_band_settings_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "ap_band_settings_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "ap_band_settings_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_channel_width", "auto"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_max_power", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_min_bitrate", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_min_power", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_rxsop", "-95"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "five_ghz_settings_valid_auto_channels.0", "36"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings0_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings0_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings0_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings0_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings1_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings1_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings1_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings1_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings10_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings10_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings10_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings10_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings11_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings11_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings11_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings11_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings12_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings12_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings12_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings12_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings13_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings13_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings13_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings13_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings14_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings14_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings14_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings14_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings2_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings2_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings2_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings2_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings3_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings3_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings3_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings3_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings4_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings4_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings4_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings4_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings5_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings5_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings5_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings5_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings6_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings6_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings6_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings6_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings7_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings7_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings7_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings7_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings8_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings8_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings8_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings8_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings9_band_operation_mode", "5ghz"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings9_band_steering_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings9_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "per_ssid_settings9_bands_enabled.0", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "transmission_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_ax_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_max_power", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_min_bitrate", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_min_power", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_rxsop", "-95"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_rf_profile.test", "two_four_ghz_settings_valid_auto_channels.0", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessRFProfilePrerequisitesConfig + testAccDataSourceMerakiWirelessRFProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceMerakiWirelessRFProfilePrerequisitesConfig + testAccNamedDataSourceMerakiWirelessRFProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessRFProfilePrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessRFProfileConfig() string {
	config := `resource "meraki_wireless_rf_profile" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	band_selection_type = "ap"` + "\n"
	config += `	client_balancing_enabled = true` + "\n"
	config += `	min_bitrate_type = "band"` + "\n"
	config += `	name = "Main Office"` + "\n"
	config += `	ap_band_settings_band_operation_mode = "5ghz"` + "\n"
	config += `	ap_band_settings_band_steering_enabled = true` + "\n"
	config += `	ap_band_settings_bands_enabled = ["5"]` + "\n"
	config += `	five_ghz_settings_channel_width = "auto"` + "\n"
	config += `	five_ghz_settings_max_power = 30` + "\n"
	config += `	five_ghz_settings_min_bitrate = 12` + "\n"
	config += `	five_ghz_settings_min_power = 8` + "\n"
	config += `	five_ghz_settings_rxsop = -95` + "\n"
	config += `	five_ghz_settings_valid_auto_channels = [36]` + "\n"
	config += `	per_ssid_settings0_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings0_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings0_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings0_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings1_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings1_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings1_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings1_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings10_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings10_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings10_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings10_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings11_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings11_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings11_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings11_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings12_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings12_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings12_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings12_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings13_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings13_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings13_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings13_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings14_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings14_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings14_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings14_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings2_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings2_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings2_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings2_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings3_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings3_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings3_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings3_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings4_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings4_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings4_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings4_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings5_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings5_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings5_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings5_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings6_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings6_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings6_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings6_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings7_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings7_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings7_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings7_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings8_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings8_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings8_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings8_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings9_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings9_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings9_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings9_bands_enabled = ["5"]` + "\n"
	config += `	transmission_enabled = true` + "\n"
	config += `	two_four_ghz_settings_ax_enabled = true` + "\n"
	config += `	two_four_ghz_settings_max_power = 30` + "\n"
	config += `	two_four_ghz_settings_min_bitrate = 11` + "\n"
	config += `	two_four_ghz_settings_min_power = 5` + "\n"
	config += `	two_four_ghz_settings_rxsop = -95` + "\n"
	config += `	two_four_ghz_settings_valid_auto_channels = [1]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_rf_profile" "test" {
			id = meraki_wireless_rf_profile.test.id
			network_id = meraki_network.test.id
			depends_on = [meraki_wireless_rf_profile.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiWirelessRFProfileConfig() string {
	config := `resource "meraki_wireless_rf_profile" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	band_selection_type = "ap"` + "\n"
	config += `	client_balancing_enabled = true` + "\n"
	config += `	min_bitrate_type = "band"` + "\n"
	config += `	name = "Main Office"` + "\n"
	config += `	ap_band_settings_band_operation_mode = "5ghz"` + "\n"
	config += `	ap_band_settings_band_steering_enabled = true` + "\n"
	config += `	ap_band_settings_bands_enabled = ["5"]` + "\n"
	config += `	five_ghz_settings_channel_width = "auto"` + "\n"
	config += `	five_ghz_settings_max_power = 30` + "\n"
	config += `	five_ghz_settings_min_bitrate = 12` + "\n"
	config += `	five_ghz_settings_min_power = 8` + "\n"
	config += `	five_ghz_settings_rxsop = -95` + "\n"
	config += `	five_ghz_settings_valid_auto_channels = [36]` + "\n"
	config += `	per_ssid_settings0_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings0_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings0_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings0_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings1_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings1_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings1_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings1_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings10_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings10_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings10_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings10_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings11_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings11_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings11_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings11_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings12_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings12_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings12_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings12_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings13_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings13_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings13_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings13_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings14_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings14_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings14_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings14_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings2_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings2_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings2_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings2_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings3_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings3_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings3_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings3_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings4_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings4_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings4_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings4_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings5_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings5_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings5_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings5_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings6_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings6_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings6_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings6_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings7_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings7_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings7_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings7_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings8_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings8_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings8_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings8_bands_enabled = ["5"]` + "\n"
	config += `	per_ssid_settings9_band_operation_mode = "5ghz"` + "\n"
	config += `	per_ssid_settings9_band_steering_enabled = true` + "\n"
	config += `	per_ssid_settings9_min_bitrate = 11` + "\n"
	config += `	per_ssid_settings9_bands_enabled = ["5"]` + "\n"
	config += `	transmission_enabled = true` + "\n"
	config += `	two_four_ghz_settings_ax_enabled = true` + "\n"
	config += `	two_four_ghz_settings_max_power = 30` + "\n"
	config += `	two_four_ghz_settings_min_bitrate = 11` + "\n"
	config += `	two_four_ghz_settings_min_power = 5` + "\n"
	config += `	two_four_ghz_settings_rxsop = -95` + "\n"
	config += `	two_four_ghz_settings_valid_auto_channels = [1]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_rf_profile" "test" {
			name = meraki_wireless_rf_profile.test.name
			network_id = meraki_network.test.id
			depends_on = [meraki_wireless_rf_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
