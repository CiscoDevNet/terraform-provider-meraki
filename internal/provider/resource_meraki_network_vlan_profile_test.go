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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiNetworkVLANProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "iname", "Profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "name", "My VLAN profile name"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "vlan_groups.0.name", "named-group-1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "vlan_groups.0.vlan_ids", "2,5-7"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "vlan_names.0.name", "named-1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_vlan_profile.test", "vlan_names.0.vlan_id", "1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkVLANProfilePrerequisitesConfig + testAccMerakiNetworkVLANProfileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkVLANProfilePrerequisitesConfig + testAccMerakiNetworkVLANProfileConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkVLANProfilePrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_vlan_profile" "default" {
  network_id = meraki_network.test.id
  iname      = "Default"
  name       = "Default Profile"
  vlan_names = [
    {
      name = "named-1"
      vlan_id = "1"
    }
  ]
  vlan_groups = [
    {
      name = "named-group-1"
      vlan_ids = "2,5-7"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkVLANProfileConfig_minimum() string {
	config := `resource "meraki_network_vlan_profile" "test" {` + "\n"
	config += `	network_id = meraki_network_vlan_profile.default.network_id` + "\n"
	config += `	iname = "Profile1"` + "\n"
	config += `	name = "My VLAN profile name"` + "\n"
	config += `	vlan_groups = [{` + "\n"
	config += `		name = "named-group-1"` + "\n"
	config += `		vlan_ids = "2,5-7"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_names = [{` + "\n"
	config += `		name = "named-1"` + "\n"
	config += `		vlan_id = "1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkVLANProfileConfig_all() string {
	config := `resource "meraki_network_vlan_profile" "test" {` + "\n"
	config += `	network_id = meraki_network_vlan_profile.default.network_id` + "\n"
	config += `	iname = "Profile1"` + "\n"
	config += `	name = "My VLAN profile name"` + "\n"
	config += `	vlan_groups = [{` + "\n"
	config += `		name = "named-group-1"` + "\n"
	config += `		vlan_ids = "2,5-7"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_names = [{` + "\n"
	config += `		name = "named-1"` + "\n"
	config += `		vlan_id = "1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
