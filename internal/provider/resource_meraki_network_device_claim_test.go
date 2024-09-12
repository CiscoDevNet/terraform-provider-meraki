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

func TestAccMerakiNetworkDeviceClaim(t *testing.T) {
	if os.Getenv("TF_VAR_CLAIM_SERIAL_1") == "" || os.Getenv("TF_VAR_CLAIM_SERIAL_2") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_CLAIM_SERIAL_1 or TF_VAR_CLAIM_SERIAL_2")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkDeviceClaimPrerequisitesConfig + testAccMerakiNetworkDeviceClaimConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkDeviceClaimPrerequisitesConfig + testAccMerakiNetworkDeviceClaimConfig_all(),
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

const testAccMerakiNetworkDeviceClaimPrerequisitesConfig = `
variable "CLAIM_SERIAL_1" {}
variable "CLAIM_SERIAL_2" {}

data "meraki_organization" "test" {
  name = "Dev"
}

resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["appliance", "switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkDeviceClaimConfig_minimum() string {
	config := `resource "meraki_network_device_claim" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	serials = [var.CLAIM_SERIAL_1]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

func testAccMerakiNetworkDeviceClaimConfig_all() string {
	config := `resource "meraki_network_device_claim" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	serials = [var.CLAIM_SERIAL_1, var.CLAIM_SERIAL_2]` + "\n"
	config += `}` + "\n"
	return config
}
