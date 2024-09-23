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

func TestAccDataSourceMerakiWirelessSSIDVPN(t *testing.T) {
	if os.Getenv("WIRELESS_SSID_VPN") == "" {
		t.Skip("skipping test, set environment variable WIRELESS_SSID_VPN")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "concentrator_vlan_id", "44"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "failover_heartbeat_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "failover_idle_timeout", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "failover_request_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_rules.0.comment", "split tunnel rule 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_rules.0.dest_cidr", "1.1.1.1/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_rules.0.dest_port", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_rules.0.policy", "allow"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid_vpn.test", "split_tunnel_rules.0.protocol", "Any"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDVPNPrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDVPNConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDVPNPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_wireless_ssid" "test" {
  network_id         = meraki_network.test.id
  number             = "0"
  name               = "My SSID"
  ip_assignment_mode = "VPN"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDVPNConfig() string {
	config := `resource "meraki_wireless_ssid_vpn" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = meraki_wireless_ssid.test.id` + "\n"
	config += `	concentrator_network_id = meraki_network.test.id` + "\n"
	config += `	concentrator_vlan_id = 44` + "\n"
	config += `	failover_heartbeat_interval = 10` + "\n"
	config += `	failover_idle_timeout = 30` + "\n"
	config += `	failover_request_ip = "1.1.1.1"` + "\n"
	config += `	split_tunnel_enabled = true` + "\n"
	config += `	split_tunnel_rules = [{` + "\n"
	config += `		comment = "split tunnel rule 1"` + "\n"
	config += `		dest_cidr = "1.1.1.1/32"` + "\n"
	config += `		dest_port = "any"` + "\n"
	config += `		policy = "allow"` + "\n"
	config += `		protocol = "Any"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid_vpn" "test" {
			network_id = meraki_network.test.id
			number = meraki_wireless_ssid.test.id
			depends_on = [meraki_wireless_ssid_vpn.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
