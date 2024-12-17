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

func TestAccMerakiApplianceTrafficShapingUplinkBandwidth(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "cellular_limit_down", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "cellular_limit_up", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "wan1_limit_down", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "wan1_limit_up", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "wan2_limit_down", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_traffic_shaping_uplink_bandwidth.test", "wan2_limit_up", "100000"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceTrafficShapingUplinkBandwidthPrerequisitesConfig + testAccMerakiApplianceTrafficShapingUplinkBandwidthConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceTrafficShapingUplinkBandwidthPrerequisitesConfig + testAccMerakiApplianceTrafficShapingUplinkBandwidthConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_traffic_shaping_uplink_bandwidth.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceTrafficShapingUplinkBandwidthImportStateIdFunc("meraki_appliance_traffic_shaping_uplink_bandwidth.test"),
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

func merakiApplianceTrafficShapingUplinkBandwidthImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceTrafficShapingUplinkBandwidthPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceTrafficShapingUplinkBandwidthConfig_minimum() string {
	config := `resource "meraki_appliance_traffic_shaping_uplink_bandwidth" "test" {` + "\n"
	config += `  network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `  wan1_limit_down = 10000` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceTrafficShapingUplinkBandwidthConfig_all() string {
	config := `resource "meraki_appliance_traffic_shaping_uplink_bandwidth" "test" {` + "\n"
	config += `  network_id = meraki_network_device_claim.test.network_id` + "\n"
	config += `  cellular_limit_down = 100000` + "\n"
	config += `  cellular_limit_up = 100000` + "\n"
	config += `  wan1_limit_down = 100000` + "\n"
	config += `  wan1_limit_up = 100000` + "\n"
	config += `  wan2_limit_down = 100000` + "\n"
	config += `  wan2_limit_up = 100000` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
