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

func TestAccDataSourceMerakiWirelessSSID(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "number", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "auth_mode", "psk"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "available_on_all_aps", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "band_selection", "5 GHz band only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "encryption_mode", "wpa"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "ip_assignment_mode", "Bridge mode"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "lan_isolation_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "mandatory_dhcp_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "min_bitrate", "5.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "name", "My SSID"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "per_client_bandwidth_limit_down", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "per_client_bandwidth_limit_up", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "per_ssid_bandwidth_limit_down", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "per_ssid_bandwidth_limit_up", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "psk", "deadbeef"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "splash_page", "Click-through splash page"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "use_vlan_tagging", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "visible", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "walled_garden_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "wpa_encryption_mode", "WPA2 only"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "dot11r_adaptive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "dot11r_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "dot11w_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "dot11w_required", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "speed_burst_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_wireless_ssid.test", "availability_tags.0", "tag1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiWirelessSSIDPrerequisitesConfig + testAccDataSourceMerakiWirelessSSIDConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiWirelessSSIDPrerequisitesConfig = `
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiWirelessSSIDConfig() string {
	config := `resource "meraki_wireless_ssid" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	number = "0"` + "\n"
	config += `	auth_mode = "psk"` + "\n"
	config += `	available_on_all_aps = false` + "\n"
	config += `	band_selection = "5 GHz band only"` + "\n"
	config += `	enabled = false` + "\n"
	config += `	encryption_mode = "wpa"` + "\n"
	config += `	ip_assignment_mode = "Bridge mode"` + "\n"
	config += `	lan_isolation_enabled = false` + "\n"
	config += `	mandatory_dhcp_enabled = false` + "\n"
	config += `	min_bitrate = 5.5` + "\n"
	config += `	name = "My SSID"` + "\n"
	config += `	per_client_bandwidth_limit_down = 0` + "\n"
	config += `	per_client_bandwidth_limit_up = 0` + "\n"
	config += `	per_ssid_bandwidth_limit_down = 0` + "\n"
	config += `	per_ssid_bandwidth_limit_up = 0` + "\n"
	config += `	psk = "deadbeef"` + "\n"
	config += `	splash_page = "Click-through splash page"` + "\n"
	config += `	use_vlan_tagging = false` + "\n"
	config += `	visible = false` + "\n"
	config += `	walled_garden_enabled = false` + "\n"
	config += `	wpa_encryption_mode = "WPA2 only"` + "\n"
	config += `	dot11r_adaptive = false` + "\n"
	config += `	dot11r_enabled = false` + "\n"
	config += `	dot11w_enabled = false` + "\n"
	config += `	dot11w_required = false` + "\n"
	config += `	speed_burst_enabled = false` + "\n"
	config += `	availability_tags = ["tag1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_wireless_ssid" "test" {
			network_id = meraki_network.test.id
			number = "0"
			depends_on = [meraki_wireless_ssid.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
