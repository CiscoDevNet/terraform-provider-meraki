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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiWirelessSSIDs(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessSSIDsPrerequisitesConfig+testAccMerakiWirelessSSIDsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDsPrerequisitesConfig+testAccMerakiWirelessSSIDsConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_wireless_ssids.test",
		ImportState: true,
		ImportStateIdFunc: merakiWirelessSSIDsImportStateIdFunc("meraki_wireless_ssids.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessSSIDsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s,%s", OrganizationId,NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessSSIDsPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessSSIDsConfig_minimum() string {
	config := `resource "meraki_wireless_ssids" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  number = "0"` + "\n"
	config += `  name = "My SSID"` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessSSIDsConfig_all() string {
	config := `resource "meraki_wireless_ssids" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  number = "0"` + "\n"
	config += `  auth_mode = "psk"` + "\n"
	config += `  available_on_all_aps = false` + "\n"
	config += `  band_selection = "5 GHz band only"` + "\n"
	config += `  enabled = false` + "\n"
	config += `  encryption_mode = "wpa"` + "\n"
	config += `  ip_assignment_mode = "Bridge mode"` + "\n"
	config += `  lan_isolation_enabled = false` + "\n"
	config += `  mandatory_dhcp_enabled = false` + "\n"
	config += `  min_bitrate = 5.5` + "\n"
	config += `  name = "My SSID"` + "\n"
	config += `  per_client_bandwidth_limit_down = 0` + "\n"
	config += `  per_client_bandwidth_limit_up = 0` + "\n"
	config += `  per_ssid_bandwidth_limit_down = 0` + "\n"
	config += `  per_ssid_bandwidth_limit_up = 0` + "\n"
	config += `  psk = "deadbeef"` + "\n"
	config += `  splash_page = "Click-through splash page"` + "\n"
	config += `  use_vlan_tagging = false` + "\n"
	config += `  visible = false` + "\n"
	config += `  walled_garden_enabled = false` + "\n"
	config += `  wpa_encryption_mode = "WPA2 only"` + "\n"
	config += `  dot11r_adaptive = false` + "\n"
	config += `  dot11r_enabled = false` + "\n"
	config += `  dot11w_enabled = false` + "\n"
	config += `  dot11w_required = false` + "\n"
	config += `  speed_burst_enabled = false` + "\n"
	config += `  availability_tags = ["tag1"]` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
