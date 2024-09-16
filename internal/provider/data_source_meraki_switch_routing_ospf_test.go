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

func TestAccDataSourceMerakiSwitchRoutingOSPF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "dead_timer_in_seconds", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "hello_timer_in_seconds", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "md5_authentication_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "md5_authentication_key_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "md5_authentication_key_passphrase", "abc1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_dead_timer_in_seconds", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_hello_timer_in_seconds", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_areas.0.area_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_areas.0.area_name", "V3 Backbone"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "v3_areas.0.area_type", "normal"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "areas.0.area_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "areas.0.area_name", "Backbone"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_routing_ospf.test", "areas.0.area_type", "normal"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSwitchRoutingOSPFPrerequisitesConfig + testAccDataSourceMerakiSwitchRoutingOSPFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSwitchRoutingOSPFPrerequisitesConfig = `
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

func testAccDataSourceMerakiSwitchRoutingOSPFConfig() string {
	config := `resource "meraki_switch_routing_ospf" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	dead_timer_in_seconds = 40` + "\n"
	config += `	enabled = true` + "\n"
	config += `	hello_timer_in_seconds = 10` + "\n"
	config += `	md5_authentication_enabled = true` + "\n"
	config += `	md5_authentication_key_id = 1` + "\n"
	config += `	md5_authentication_key_passphrase = "abc1234"` + "\n"
	config += `	v3_dead_timer_in_seconds = 40` + "\n"
	config += `	v3_enabled = true` + "\n"
	config += `	v3_hello_timer_in_seconds = 10` + "\n"
	config += `	v3_areas = [{` + "\n"
	config += `		area_id = "0"` + "\n"
	config += `		area_name = "V3 Backbone"` + "\n"
	config += `		area_type = "normal"` + "\n"
	config += `	}]` + "\n"
	config += `	areas = [{` + "\n"
	config += `		area_id = "0"` + "\n"
	config += `		area_name = "Backbone"` + "\n"
	config += `		area_type = "normal"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_routing_ospf" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_switch_routing_ospf.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig