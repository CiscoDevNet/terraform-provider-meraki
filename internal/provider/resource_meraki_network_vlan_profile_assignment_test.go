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

func TestAccMerakiNetworkVLANProfileAssignment(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkVLANProfileAssignmentPrerequisitesConfig + testAccMerakiNetworkVLANProfileAssignmentConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkVLANProfileAssignmentPrerequisitesConfig + testAccMerakiNetworkVLANProfileAssignmentConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_vlan_profile_assignment.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkVLANProfileAssignmentImportStateIdFunc("meraki_network_vlan_profile_assignment.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiNetworkVLANProfileAssignmentImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		VlanProfileIname := primary.Attributes["vlan_profile_iname"]

		return fmt.Sprintf("%s,%s", NetworkId, VlanProfileIname), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkVLANProfileAssignmentPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_switch_1_serial]
}
resource "meraki_network_vlan_profile" "default" {
  network_id = meraki_network.test.id
  iname      = "Default"
  name       = "Default Profile"
  vlan_names = [
    {
      name    = "named-1"
      vlan_id = "1"
    }
  ]
  vlan_groups = [
    {
      name     = "named-group-1"
      vlan_ids = "2,5-7"
    }
  ]
}
resource "meraki_network_vlan_profile" "test" {
  network_id = meraki_network_vlan_profile.default.network_id
  iname      = "Profile1"
  name       = "Test Profile"
  vlan_names = [
    {
      name    = "named-1"
      vlan_id = "1"
    }
  ]
  vlan_groups = [
    {
      name     = "named-group-1"
      vlan_ids = "2,5-7"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkVLANProfileAssignmentConfig_minimum() string {
	config := `resource "meraki_network_vlan_profile_assignment" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  vlan_profile_iname = meraki_network_vlan_profile.test.iname` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkVLANProfileAssignmentConfig_all() string {
	config := `resource "meraki_network_vlan_profile_assignment" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  vlan_profile_iname = meraki_network_vlan_profile.test.iname` + "\n"
	config += `  serials = [tolist(meraki_network_device_claim.test.serials)[0]]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
