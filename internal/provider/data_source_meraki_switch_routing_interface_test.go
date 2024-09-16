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

func TestAccDataSourceMerakiSwitchRoutingInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "default_gateway", "192.168.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "interface_ip", "192.168.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "multicast_routing", "disabled"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "name", "L3 interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "subnet", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "vlan_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "ipv6_address", "1:2:3:4::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "ipv6_assignment_mode", "static"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "ipv6_gateway", "1:2:3:4::2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_interface.test", "ipv6_prefix", "1:2:3:4::/64"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSwitchRoutingInterfacePrerequisitesConfig + testAccDataSourceMerakiSwitchRoutingInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceMerakiSwitchRoutingInterfacePrerequisitesConfig + testAccNamedDataSourceMerakiSwitchRoutingInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSwitchRoutingInterfacePrerequisitesConfig = `
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
  serials    = ["Q5KD-PCG4-HB8R"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSwitchRoutingInterfaceConfig() string {
	config := `resource "meraki_switch_routing_interface" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	default_gateway = "192.168.1.1"` + "\n"
	config += `	interface_ip = "192.168.1.2"` + "\n"
	config += `	multicast_routing = "disabled"` + "\n"
	config += `	name = "L3 interface"` + "\n"
	config += `	subnet = "192.168.1.0/24"` + "\n"
	config += `	vlan_id = 100` + "\n"
	config += `	ipv6_address = "1:2:3:4::1"` + "\n"
	config += `	ipv6_assignment_mode = "static"` + "\n"
	config += `	ipv6_gateway = "1:2:3:4::2"` + "\n"
	config += `	ipv6_prefix = "1:2:3:4::/64"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_routing_interface" "test" {
			id = meraki_switch_routing_interface.test.id
			serial = tolist(meraki_network_device_claim.test.serials)[0]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiSwitchRoutingInterfaceConfig() string {
	config := `resource "meraki_switch_routing_interface" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	default_gateway = "192.168.1.1"` + "\n"
	config += `	interface_ip = "192.168.1.2"` + "\n"
	config += `	multicast_routing = "disabled"` + "\n"
	config += `	name = "L3 interface"` + "\n"
	config += `	subnet = "192.168.1.0/24"` + "\n"
	config += `	vlan_id = 100` + "\n"
	config += `	ipv6_address = "1:2:3:4::1"` + "\n"
	config += `	ipv6_assignment_mode = "static"` + "\n"
	config += `	ipv6_gateway = "1:2:3:4::2"` + "\n"
	config += `	ipv6_prefix = "1:2:3:4::/64"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_routing_interface" "test" {
			name = meraki_switch_routing_interface.test.name
			serial = tolist(meraki_network_device_claim.test.serials)[0]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
