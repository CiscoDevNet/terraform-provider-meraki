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

func TestAccMerakiNetworkFloorPlan(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_floor_plan.test", "name", "HQ Floor Plan"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkFloorPlanPrerequisitesConfig + testAccMerakiNetworkFloorPlanConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkFloorPlanPrerequisitesConfig + testAccMerakiNetworkFloorPlanConfig_all(),
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

const testAccMerakiNetworkFloorPlanPrerequisitesConfig = `
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

func testAccMerakiNetworkFloorPlanConfig_minimum() string {
	config := `resource "meraki_network_floor_plan" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	image_contents = "R0lGODdhAQABAIEAAP///wAAAAAAAAAAACwAAAAAAQABAAAIBAABBAQAOw=="` + "\n"
	config += `	name = "HQ Floor Plan"` + "\n"
	config += `	center_lat = 37.770040510499996` + "\n"
	config += `	center_lng = -122.38714009525` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkFloorPlanConfig_all() string {
	config := `resource "meraki_network_floor_plan" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	image_contents = "R0lGODdhAQABAIEAAP///wAAAAAAAAAAACwAAAAAAQABAAAIBAABBAQAOw=="` + "\n"
	config += `	name = "HQ Floor Plan"` + "\n"
	config += `	bottom_left_corner_lat = 37.770040510499996` + "\n"
	config += `	bottom_left_corner_lng = -122.38714009525` + "\n"
	config += `	bottom_right_corner_lat = 37.770040510499996` + "\n"
	config += `	bottom_right_corner_lng = -121.38714009525` + "\n"
	config += `	top_left_corner_lat = 38.770040510499996` + "\n"
	config += `	top_left_corner_lng = -122.38714009525` + "\n"
	config += `	top_right_corner_lat = 38.770040510499996` + "\n"
	config += `	top_right_corner_lng = -121.38714009525` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
