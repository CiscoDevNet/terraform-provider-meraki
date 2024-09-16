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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiSwitchRoutingMulticastRendezvousPoint(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_routing_multicast_rendezvous_point.test", "multicast_group", "Any"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchRoutingMulticastRendezvousPointPrerequisitesConfig + testAccMerakiSwitchRoutingMulticastRendezvousPointConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchRoutingMulticastRendezvousPointPrerequisitesConfig + testAccMerakiSwitchRoutingMulticastRendezvousPointConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchRoutingMulticastRendezvousPointPrerequisitesConfig = `
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
resource "meraki_switch_routing_interface" "test" {
  serial            = tolist(meraki_network_device_claim.test.serials)[0]
  default_gateway   = "192.168.1.1"
  interface_ip      = "192.168.1.2"
  multicast_routing = "enabled"
  name              = "L3 interface"
  subnet            = "192.168.1.0/24"
  vlan_id           = 100
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchRoutingMulticastRendezvousPointConfig_minimum() string {
	config := `resource "meraki_switch_routing_multicast_rendezvous_point" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	interface_ip = meraki_switch_routing_interface.test.interface_ip` + "\n"
	config += `	multicast_group = "Any"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchRoutingMulticastRendezvousPointConfig_all() string {
	config := `resource "meraki_switch_routing_multicast_rendezvous_point" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	interface_ip = meraki_switch_routing_interface.test.interface_ip` + "\n"
	config += `	multicast_group = "Any"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll