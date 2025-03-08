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

func TestAccMerakiCameraDeviceWirelessProfiles(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_camera_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_camera_1_serial")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiCameraDeviceWirelessProfilesPrerequisitesConfig + testAccMerakiCameraDeviceWirelessProfilesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiCameraDeviceWirelessProfilesPrerequisitesConfig + testAccMerakiCameraDeviceWirelessProfilesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_camera_device_wireless_profiles.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiCameraDeviceWirelessProfilesImportStateIdFunc("meraki_camera_device_wireless_profiles.test"),
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

func merakiCameraDeviceWirelessProfilesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		Serial := primary.Attributes["serial"]

		return fmt.Sprintf("%s", Serial), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiCameraDeviceWirelessProfilesPrerequisitesConfig = `
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
resource "meraki_camera_wireless_profile" "test" {
  network_id = meraki_network.test.id
  name = "wireless profile A"
  ssid_auth_mode = "psk"
  ssid_name = "ssid test"
  ssid_psk = "password123"
}
resource "meraki_camera_wireless_profile" "test2" {
  network_id = meraki_network.test.id
  name = "wireless profile B"
  ssid_auth_mode = "psk"
  ssid_name = "ssid test2"
  ssid_psk = "password123"
}
resource "meraki_camera_wireless_profile" "test3" {
  network_id = meraki_network.test.id
  name = "wireless profile C"
  ssid_auth_mode = "psk"
  ssid_name = "ssid test3"
  ssid_psk = "password123"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiCameraDeviceWirelessProfilesConfig_minimum() string {
	config := `resource "meraki_camera_device_wireless_profiles" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  ids_primary = meraki_camera_wireless_profile.test.id` + "\n"
	config += `  ids_secondary = meraki_camera_wireless_profile.test2.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiCameraDeviceWirelessProfilesConfig_all() string {
	config := `resource "meraki_camera_device_wireless_profiles" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  ids_backup = meraki_camera_wireless_profile.test3.id` + "\n"
	config += `  ids_primary = meraki_camera_wireless_profile.test.id` + "\n"
	config += `  ids_secondary = meraki_camera_wireless_profile.test2.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
