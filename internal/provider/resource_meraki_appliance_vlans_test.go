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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiApplianceVLANs(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceVLANsPrerequisitesConfig + testAccMerakiApplianceVLANsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceVLANsPrerequisitesConfig + testAccMerakiApplianceVLANsConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:      "meraki_appliance_vlans.test",
		ImportState:       true,
		ImportStateIdFunc: merakiApplianceVLANsImportStateIdFunc("meraki_appliance_vlans.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiApplianceVLANsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s,%s", OrganizationId, NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceVLANsPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceVLANsConfig_minimum() string {
	config := `resource "meraki_appliance_vlans" "test" {` + "\n"
	config += ` network_id = meraki_appliance_vlans_settings.test.network_id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  appliance_ip = "192.168.1.2"` + "\n"
	config += `  vlan_id = "1234"` + "\n"
	config += `  name = "My VLAN"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceVLANsConfig_all() string {
	config := `resource "meraki_appliance_vlans" "test" {` + "\n"
	config += ` network_id = meraki_appliance_vlans_settings.test.network_id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  appliance_ip = "192.168.1.2"` + "\n"
	config += `  dhcp_boot_options_enabled = false` + "\n"
	config += `  dhcp_handling = "Run a DHCP server"` + "\n"
	config += `  dhcp_lease_time = "1 day"` + "\n"
	config += `  vlan_id = "1234"` + "\n"
	config += `  name = "My VLAN"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += `  ipv6_enabled = true` + "\n"
	config += `  mandatory_dhcp_enabled = true` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
