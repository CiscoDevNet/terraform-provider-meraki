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

func TestAccDataSourceMerakiApplianceVLAN(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "appliance_ip", "192.168.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "dhcp_boot_options_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "dhcp_handling", "Run a DHCP server"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "dhcp_lease_time", "1 day"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "dns_nameservers", "upstream_dns"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "vlan_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "name", "My VLAN"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "subnet", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.autonomous", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.disabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.static_appliance_ip6", "2001:321:3c4d:15::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.static_prefix", "2001:321:3c4d:15::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.origin_type", "internet"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "ipv6_prefix_assignments.0.origin_interfaces.0", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_vlan.test", "mandatory_dhcp_enabled", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceVLANPrerequisitesConfig + testAccDataSourceMerakiApplianceVLANConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceVLANPrerequisitesConfig = `
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
resource "meraki_appliance_vlans_settings" "test" {
  network_id = meraki_network.test.id
  vlans_enabled = true
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceVLANConfig() string {
	config := `resource "meraki_appliance_vlan" "test" {` + "\n"
	config += `  network_id = meraki_appliance_vlans_settings.test.network_id` + "\n"
	config += `  appliance_ip = "192.168.1.2"` + "\n"
	config += `  dhcp_boot_options_enabled = false` + "\n"
	config += `  dhcp_handling = "Run a DHCP server"` + "\n"
	config += `  dhcp_lease_time = "1 day"` + "\n"
	config += `  dns_nameservers = "upstream_dns"` + "\n"
	config += `  vlan_id = "1234"` + "\n"
	config += `  name = "My VLAN"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += `  ipv6_enabled = true` + "\n"
	config += `  ipv6_prefix_assignments = [{` + "\n"
	config += `    autonomous = false` + "\n"
	config += `    disabled = false` + "\n"
	config += `    static_appliance_ip6 = "2001:321:3c4d:15::1"` + "\n"
	config += `    static_prefix = "2001:321:3c4d:15::/64"` + "\n"
	config += `    origin_type = "internet"` + "\n"
	config += `    origin_interfaces = ["wan1"]` + "\n"
	config += `  }]` + "\n"
	config += `  mandatory_dhcp_enabled = true` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_vlan" "test" {
			id = meraki_appliance_vlan.test.id
			network_id = meraki_appliance_vlans_settings.test.network_id
			depends_on = [meraki_appliance_vlan.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiApplianceVLANConfig() string {
	config := `resource "meraki_appliance_vlan" "test" {` + "\n"
	config += `  network_id = meraki_appliance_vlans_settings.test.network_id` + "\n"
	config += `  appliance_ip = "192.168.1.2"` + "\n"
	config += `  dhcp_boot_options_enabled = false` + "\n"
	config += `  dhcp_handling = "Run a DHCP server"` + "\n"
	config += `  dhcp_lease_time = "1 day"` + "\n"
	config += `  dns_nameservers = "upstream_dns"` + "\n"
	config += `  vlan_id = "1234"` + "\n"
	config += `  name = "My VLAN"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += `  ipv6_enabled = true` + "\n"
	config += `  ipv6_prefix_assignments = [{` + "\n"
	config += `    autonomous = false` + "\n"
	config += `    disabled = false` + "\n"
	config += `    static_appliance_ip6 = "2001:321:3c4d:15::1"` + "\n"
	config += `    static_prefix = "2001:321:3c4d:15::/64"` + "\n"
	config += `    origin_type = "internet"` + "\n"
	config += `    origin_interfaces = ["wan1"]` + "\n"
	config += `  }]` + "\n"
	config += `  mandatory_dhcp_enabled = true` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_vlan" "test" {
			name = meraki_appliance_vlan.test.name
			network_id = meraki_appliance_vlans_settings.test.network_id
			depends_on = [meraki_appliance_vlan.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
