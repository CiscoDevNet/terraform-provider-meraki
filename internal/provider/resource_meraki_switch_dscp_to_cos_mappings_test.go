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

func TestAccMerakiSwitchDSCPToCoSMappings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dscp_to_cos_mappings.test", "mappings.0.cos", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dscp_to_cos_mappings.test", "mappings.0.dscp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dscp_to_cos_mappings.test", "mappings.0.title", "Video"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchDSCPToCoSMappingsPrerequisitesConfig + testAccMerakiSwitchDSCPToCoSMappingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchDSCPToCoSMappingsPrerequisitesConfig + testAccMerakiSwitchDSCPToCoSMappingsConfig_all(),
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

const testAccMerakiSwitchDSCPToCoSMappingsPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchDSCPToCoSMappingsConfig_minimum() string {
	config := `resource "meraki_switch_dscp_to_cos_mappings" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	mappings = [{` + "\n"
	config += `		cos = 1` + "\n"
	config += `		dscp = 1` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchDSCPToCoSMappingsConfig_all() string {
	config := `resource "meraki_switch_dscp_to_cos_mappings" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	mappings = [{` + "\n"
	config += `		cos = 1` + "\n"
	config += `		dscp = 1` + "\n"
	config += `		title = "Video"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll