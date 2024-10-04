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

func TestAccDataSourceMerakiApplianceSSID(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "auth_mode", "8021x-radius"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "name", "My SSID"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "visible", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "wpa_encryption_mode", "WPA2 only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "radius_servers.0.host", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_ssid.test", "radius_servers.0.port", "1000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceSSIDPrerequisitesConfig + testAccDataSourceMerakiApplianceSSIDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceSSIDPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceSSIDConfig() string {
	config := `resource "meraki_appliance_ssid" "test" {` + "\n"
	config += `	network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `	number = "1"` + "\n"
	config += `	auth_mode = "8021x-radius"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	name = "My SSID"` + "\n"
	config += `	visible = true` + "\n"
	config += `	wpa_encryption_mode = "WPA2 only"` + "\n"
	config += `	radius_servers = [{` + "\n"
	config += `		host = "0.0.0.0"` + "\n"
	config += `		port = 1000` + "\n"
	config += `		secret = "secret"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_ssid" "test" {
			network_id = meraki_network_device_claim.test.network_id
			number = "1"
			depends_on = [meraki_appliance_ssid.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
