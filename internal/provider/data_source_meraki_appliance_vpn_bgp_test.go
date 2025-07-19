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

func TestAccDataSourceMerakiApplianceVPNBGP(t *testing.T) {
	if os.Getenv("APPLIANCE_VPN_BGP") == "" {
		t.Skip("skipping test, set environment variable APPLIANCE_VPN_BGP")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "as_number", "64515"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "ibgp_hold_timer", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.allow_transit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.ebgp_hold_timer", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.ebgp_multihop", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.ip", "10.10.10.22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.multi_exit_discriminator", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.next_hop_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.receive_limit", "120"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.remote_as_number", "64343"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.source_interface", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.weight", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.authentication_password", "abc123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.ipv6_address", "2002::1234:abcd:ffff:c0a8:101"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.ttl_security_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vpn_bgp.test", "neighbors.0.path_prepend.0", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceVPNBGPPrerequisitesConfig + testAccDataSourceMerakiApplianceVPNBGPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceVPNBGPPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
resource "meraki_appliance_site_to_site_vpn" "test" {
  network_id = meraki_network.test.id
  mode       = "hub"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceVPNBGPConfig() string {
	config := `resource "meraki_appliance_vpn_bgp" "test" {` + "\n"
	config += `  network_id = meraki_appliance_site_to_site_vpn.test.network_id` + "\n"
	config += `  as_number = 64515` + "\n"
	config += `  enabled = true` + "\n"
	config += `  ibgp_hold_timer = 120` + "\n"
	config += `  neighbors = [{` + "\n"
	config += `    allow_transit = true` + "\n"
	config += `    ebgp_hold_timer = 180` + "\n"
	config += `    ebgp_multihop = 2` + "\n"
	config += `    ip = "10.10.10.22"` + "\n"
	config += `    multi_exit_discriminator = 2` + "\n"
	config += `    next_hop_ip = "1.2.3.4"` + "\n"
	config += `    receive_limit = 120` + "\n"
	config += `    remote_as_number = 64343` + "\n"
	config += `    source_interface = "wan1"` + "\n"
	config += `    weight = 10` + "\n"
	config += `    authentication_password = "abc123"` + "\n"
	config += `    ipv6_address = "2002::1234:abcd:ffff:c0a8:101"` + "\n"
	config += `    ttl_security_enabled = false` + "\n"
	config += `    path_prepend = [1]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_vpn_bgp" "test" {
			network_id = meraki_appliance_site_to_site_vpn.test.network_id
			depends_on = [meraki_appliance_vpn_bgp.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
