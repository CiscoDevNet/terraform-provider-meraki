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

func TestAccDataSourceMerakiSwitchAccessControlLists(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.comment", "Deny SSH"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.dst_cidr", "172.16.3.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.dst_port", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.ip_version", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.policy", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.src_cidr", "10.1.10.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.src_port", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_access_control_lists.test", "rules.0.vlan", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSwitchAccessControlListsPrerequisitesConfig + testAccDataSourceMerakiSwitchAccessControlListsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSwitchAccessControlListsPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSwitchAccessControlListsConfig() string {
	config := `resource "meraki_switch_access_control_lists" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	rules = [{` + "\n"
	config += `		comment = "Deny SSH"` + "\n"
	config += `		dst_cidr = "172.16.3.0/24"` + "\n"
	config += `		dst_port = "22"` + "\n"
	config += `		ip_version = "ipv4"` + "\n"
	config += `		policy = "deny"` + "\n"
	config += `		protocol = "tcp"` + "\n"
	config += `		src_cidr = "10.1.10.0/24"` + "\n"
	config += `		src_port = "any"` + "\n"
	config += `		vlan = "10"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_access_control_lists" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_switch_access_control_lists.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
