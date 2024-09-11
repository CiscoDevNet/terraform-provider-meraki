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

func TestAccDataSourceMerakiNetwork(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network.test", "name", "Main Office"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network.test", "notes", "Additional description of the network"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network.test", "time_zone", "America/Los_Angeles"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network.test", "product_types.0", "switch"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network.test", "tags.0", "tag1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkPrerequisitesConfig + testAccDataSourceMerakiNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceMerakiNetworkPrerequisitesConfig + testAccNamedDataSourceMerakiNetworkConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkConfig() string {
	config := `resource "meraki_network" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	name = "Main Office"` + "\n"
	config += `	notes = "Additional description of the network"` + "\n"
	config += `	time_zone = "America/Los_Angeles"` + "\n"
	config += `	product_types = ["switch"]` + "\n"
	config += `	tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network" "test" {
			id = meraki_network.test.id
			organization_id = data.meraki_organization.test.id
		}
	`
	return config
}

func testAccNamedDataSourceMerakiNetworkConfig() string {
	config := `resource "meraki_network" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	name = "Main Office"` + "\n"
	config += `	notes = "Additional description of the network"` + "\n"
	config += `	time_zone = "America/Los_Angeles"` + "\n"
	config += `	product_types = ["switch"]` + "\n"
	config += `	tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network" "test" {
			name = meraki_network.test.name
			organization_id = data.meraki_organization.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
