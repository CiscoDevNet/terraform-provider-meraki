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

func TestAccDataSourceMerakiSwitchOrganizationPortsProfile(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "description", "IP Phones for all office workers"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "is_organization_wide", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "name", "Phone"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_access_policy_type", "Open"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_allowed_vlans", "1-100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_dai_trusted", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_isolation_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_peer_sgt_capable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_poe_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_rstp_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_storm_control_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_stp_guard", "disabled"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_type", "access"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_udld", "Alert only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_vlan", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "port_voice_vlan", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_switch_organization_ports_profile.test", "tags.0", "tag1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSwitchOrganizationPortsProfilePrerequisitesConfig + testAccDataSourceMerakiSwitchOrganizationPortsProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSwitchOrganizationPortsProfilePrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSwitchOrganizationPortsProfileConfig() string {
	config := `resource "meraki_switch_organization_ports_profile" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	description = "IP Phones for all office workers"` + "\n"
	config += `	is_organization_wide = true` + "\n"
	config += `	name = "Phone"` + "\n"
	config += `	port_access_policy_type = "Open"` + "\n"
	config += `	port_allowed_vlans = "1-100"` + "\n"
	config += `	port_dai_trusted = false` + "\n"
	config += `	port_isolation_enabled = false` + "\n"
	config += `	port_peer_sgt_capable = false` + "\n"
	config += `	port_poe_enabled = true` + "\n"
	config += `	port_rstp_enabled = true` + "\n"
	config += `	port_storm_control_enabled = true` + "\n"
	config += `	port_stp_guard = "disabled"` + "\n"
	config += `	port_type = "access"` + "\n"
	config += `	port_udld = "Alert only"` + "\n"
	config += `	port_vlan = 10` + "\n"
	config += `	port_voice_vlan = 20` + "\n"
	config += `	tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_organization_ports_profile" "test" {
			id = meraki_switch_organization_ports_profile.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_switch_organization_ports_profile.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiSwitchOrganizationPortsProfileConfig() string {
	config := `resource "meraki_switch_organization_ports_profile" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	description = "IP Phones for all office workers"` + "\n"
	config += `	is_organization_wide = true` + "\n"
	config += `	name = "Phone"` + "\n"
	config += `	port_access_policy_type = "Open"` + "\n"
	config += `	port_allowed_vlans = "1-100"` + "\n"
	config += `	port_dai_trusted = false` + "\n"
	config += `	port_isolation_enabled = false` + "\n"
	config += `	port_peer_sgt_capable = false` + "\n"
	config += `	port_poe_enabled = true` + "\n"
	config += `	port_rstp_enabled = true` + "\n"
	config += `	port_storm_control_enabled = true` + "\n"
	config += `	port_stp_guard = "disabled"` + "\n"
	config += `	port_type = "access"` + "\n"
	config += `	port_udld = "Alert only"` + "\n"
	config += `	port_vlan = 10` + "\n"
	config += `	port_voice_vlan = 20` + "\n"
	config += `	tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_switch_organization_ports_profile" "test" {
			name = meraki_switch_organization_ports_profile.test.name
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_switch_organization_ports_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
