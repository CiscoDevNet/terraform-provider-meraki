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

func TestAccDataSourceMerakiDeviceManagementInterface(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_static_gateway_ip", "1.2.3.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_static_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_static_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_using_static_ip", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_vlan", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_management_interface.test", "wan1_static_dns.0", "1.2.3.2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiDeviceManagementInterfacePrerequisitesConfig + testAccDataSourceMerakiDeviceManagementInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiDeviceManagementInterfacePrerequisitesConfig = `
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

func testAccDataSourceMerakiDeviceManagementInterfaceConfig() string {
	config := `resource "meraki_device_management_interface" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	wan1_static_gateway_ip = "1.2.3.1"` + "\n"
	config += `	wan1_static_ip = "1.2.3.4"` + "\n"
	config += `	wan1_static_subnet_mask = "255.255.255.0"` + "\n"
	config += `	wan1_using_static_ip = true` + "\n"
	config += `	wan1_vlan = 7` + "\n"
	config += `	wan1_static_dns = ["1.2.3.2"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_device_management_interface" "test" {
			serial = tolist(meraki_network_device_claim.test.serials)[0]
			depends_on = [meraki_device_management_interface.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
