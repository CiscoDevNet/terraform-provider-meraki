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

func TestAccDataSourceMerakiDevice(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device.test", "address", "1600 Pennsylvania Ave"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device.test", "lat", "37.4180951010362"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device.test", "lng", "-122.098531723022"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device.test", "name", "My AP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device.test", "notes", "My AP's note"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiDevicePrerequisitesConfig + testAccDataSourceMerakiDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiDevicePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_switch_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiDeviceConfig() string {
	config := `resource "meraki_device" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  address = "1600 Pennsylvania Ave"` + "\n"
	config += `  lat = 37.4180951010362` + "\n"
	config += `  lng = -122.098531723022` + "\n"
	config += `  name = "My AP"` + "\n"
	config += `  notes = "My AP's note"` + "\n"
	config += `  tags = ["recently-added"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_device" "test" {
			serial = tolist(meraki_network_device_claim.test.serials)[0]
			depends_on = [meraki_device.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
