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

func TestAccMerakiSwitchRoutingMulticast(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_routing_multicast.test", "default_settings_flood_unknown_multicast_traffic_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_routing_multicast.test", "default_settings_igmp_snooping_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_routing_multicast.test", "overrides.0.flood_unknown_multicast_traffic_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_routing_multicast.test", "overrides.0.igmp_snooping_enabled", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchRoutingMulticastPrerequisitesConfig + testAccMerakiSwitchRoutingMulticastConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchRoutingMulticastPrerequisitesConfig + testAccMerakiSwitchRoutingMulticastConfig_all(),
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

const testAccMerakiSwitchRoutingMulticastPrerequisitesConfig = `
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
  serials    = ["Q3LN-KJJ5-45ZR"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchRoutingMulticastConfig_minimum() string {
	config := `resource "meraki_switch_routing_multicast" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	default_settings_flood_unknown_multicast_traffic_enabled = true` + "\n"
	config += `	default_settings_igmp_snooping_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchRoutingMulticastConfig_all() string {
	config := `resource "meraki_switch_routing_multicast" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	default_settings_flood_unknown_multicast_traffic_enabled = true` + "\n"
	config += `	default_settings_igmp_snooping_enabled = true` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		flood_unknown_multicast_traffic_enabled = true` + "\n"
	config += `		igmp_snooping_enabled = true` + "\n"
	config += `		switches = [tolist(meraki_network_device_claim.test.serials)[0]]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
