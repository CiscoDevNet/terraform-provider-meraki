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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiSwitchStack(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_stack.test", "name", "A cool stack"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSwitchStackPrerequisitesConfig + testAccDataSourceMerakiSwitchStackConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceMerakiSwitchStackPrerequisitesConfig + testAccNamedDataSourceMerakiSwitchStackConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSwitchStackPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = ["Q5KD-PCG4-HB8R", "Q5KD-CU8N-DEDR"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSwitchStackConfig() string {
	config := `resource "meraki_switch_stack" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	name = "A cool stack"` + "\n"
	config += `	serials = meraki_network_device_claim.test.serials` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_stack" "test" {
			id = meraki_switch_stack.test.id
			network_id = meraki_network.test.id
		}
	`
	return config
}

func testAccNamedDataSourceMerakiSwitchStackConfig() string {
	config := `resource "meraki_switch_stack" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	name = "A cool stack"` + "\n"
	config += `	serials = meraki_network_device_claim.test.serials` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_stack" "test" {
			name = meraki_switch_stack.test.name
			network_id = meraki_network.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
