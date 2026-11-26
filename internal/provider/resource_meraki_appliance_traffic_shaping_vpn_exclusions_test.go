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

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiApplianceTrafficShapingVPNExclusions(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_network_2") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_network_2 and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "applications.0.id", "meraki:layer7/application/2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "applications.0.name", "File Transfer Protocol"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "applications.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "applications.0.source_cidr", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "applications.0.source_port", "8000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "custom.0.destination", "192.168.3.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "custom.0.port", "8000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "custom.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "custom.0.source_cidr", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "custom.0.source_port", "8000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "major_applications.0.id", "meraki:vpnExclusions/application/2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "major_applications.0.name", "Office 365 Sharepoint"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "major_applications.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_vpn_exclusions.test", "major_applications.0.source_vlan_id", "20"))

	var steps []resource.TestStep
	var tfVersion *goversion.Version
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceTrafficShapingVPNExclusionsPrerequisitesConfig + testAccMerakiApplianceTrafficShapingVPNExclusionsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceTrafficShapingVPNExclusionsPrerequisitesConfig + testAccMerakiApplianceTrafficShapingVPNExclusionsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	importStateVerifyIgnore := []string{}
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_traffic_shaping_vpn_exclusions.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceTrafficShapingVPNExclusionsImportStateIdFunc("meraki_appliance_traffic_shaping_vpn_exclusions.test"),
		ImportStateVerifyIgnore: importStateVerifyIgnore,
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			terraformVersionCapture{Version: &tfVersion},
		},
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiApplianceTrafficShapingVPNExclusionsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceTrafficShapingVPNExclusionsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_network_2" {}
variable "test_appliance_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_appliance_1_serial]
}
resource "meraki_network" "test_spoke" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network_2
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}
resource "meraki_appliance_site_to_site_vpn" "test" {
  network_id = meraki_network.test.id
  mode       = "hub"
  subnets    = [
    {
      local_subnet = "192.168.10.0/24"
      use_vpn      = true
    }
  ]
  depends_on = [
    meraki_appliance_vlan.test,
  ]
}
resource "meraki_appliance_vlans_settings" "test" {
  network_id = meraki_network.test.id
  vlans_enabled = true
}
resource "meraki_appliance_vlan" "test" {
  network_id   = meraki_network.test.id
  name         = "test vlan 10"
  vlan_id      = 10
  appliance_ip = "192.168.10.1"
  subnet       = "192.168.10.0/24"
  depends_on   = [
    meraki_network_device_claim.test,
    meraki_appliance_vlans_settings.test,
  ]
}
resource "meraki_appliance_vlans_settings" "test_spoke" {
  network_id = meraki_network.test_spoke.id
  vlans_enabled = true
}
resource "meraki_appliance_vlan" "test_spoke" {
  network_id   = meraki_network.test_spoke.id
  name         = "test vlan 20"
  vlan_id      = 20
  appliance_ip = "192.168.20.1"
  subnet       = "192.168.20.0/24"
  depends_on   = [
    meraki_appliance_vlans_settings.test_spoke,
  ]
}
resource "meraki_appliance_site_to_site_vpn" "test_spoke" {
  network_id = meraki_network.test_spoke.id
  mode       = "spoke"
  hubs       = [
    {
      hub_id            = meraki_network.test.id
      use_default_route = true
    }
  ]
  subnets = [
    {
      local_subnet  = "192.168.20.0/24"
      use_vpn       = true
    }
  ]
  depends_on = [
    meraki_appliance_site_to_site_vpn.test,
    meraki_appliance_vlan.test_spoke,
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceTrafficShapingVPNExclusionsConfig_minimum() string {
	config := `resource "meraki_appliance_traffic_shaping_vpn_exclusions" "test" {` + "\n"
	config += `  network_id = meraki_appliance_site_to_site_vpn.test_spoke.network_id` + "\n"
	config += `  custom = [{` + "\n"
	config += `    destination = "192.168.3.0/24"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccMerakiApplianceTrafficShapingVPNExclusionsConfig_all() string {
	config := `resource "meraki_appliance_traffic_shaping_vpn_exclusions" "test" {` + "\n"
	config += `  network_id = meraki_appliance_site_to_site_vpn.test_spoke.network_id` + "\n"
	config += `  applications = [{` + "\n"
	config += `    id = "meraki:layer7/application/2"` + "\n"
	config += `    name = "File Transfer Protocol"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    source_cidr = "192.168.1.0/24"` + "\n"
	config += `    source_port = "8000"` + "\n"
	config += `  }]` + "\n"
	config += `  custom = [{` + "\n"
	config += `    destination = "192.168.3.0/24"` + "\n"
	config += `    port = "8000"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    source_cidr = "192.168.1.0/24"` + "\n"
	config += `    source_port = "8000"` + "\n"
	config += `  }]` + "\n"
	config += `  major_applications = [{` + "\n"
	config += `    id = "meraki:vpnExclusions/application/2"` + "\n"
	config += `    name = "Office 365 Sharepoint"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    source_vlan_id = "20"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
