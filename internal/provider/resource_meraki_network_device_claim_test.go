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

func TestAccMerakiNetworkDeviceClaim(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_claim_serial_1") == "" || os.Getenv("TF_VAR_test_claim_serial_2") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_claim_serial_1 and TF_VAR_test_claim_serial_2")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_device_claim.test", "details_by_device.0.details.0.name", "username"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_device_claim.test", "details_by_device.0.details.0.value", "milesmeraki"))

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
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_device_claim.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkDeviceClaimImportStateIdFunc("meraki_network_device_claim.test"),
		ImportStateVerifyIgnore: []string{"details_by_device"},
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

func merakiNetworkDeviceClaimImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkDeviceClaimPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_claim_serial_1" {}
variable "test_claim_serial_2" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["appliance", "switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkDeviceClaimConfig_minimum() string {
	config := `resource "meraki_network_device_claim" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  serials = [var.test_claim_serial_1]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

func testAccMerakiNetworkDeviceClaimConfig_all() string {
	config := `resource "meraki_network_device_claim" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	serials = [var.test_claim_serial_1, var.test_claim_serial_2]` + "\n"
	config += `	details_by_device = [{` + "\n"
	config += `	  serial = var.test_claim_serial_2` + "\n"
	config += `	  details = [` + "\n"
	config += `	    {` + "\n"
	config += `	      name = "username"` + "\n"
	config += `	      value = "milesmeraki"` + "\n"
	config += `	    },` + "\n"
	config += `	    {` + "\n"
	config += `	      name = "password"` + "\n"
	config += `	      value = "cisco"` + "\n"
	config += `	    }` + "\n"
	config += `	  ]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
