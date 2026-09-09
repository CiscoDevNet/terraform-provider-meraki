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
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

// Actions require Terraform >= 1.14.0 (the `action` block and the ProviderWithActions
// interface are not available in older releases), hence the TerraformVersionChecks gate below.
func TestAccMerakiGenerateApplianceVMXAuthenticationToken(t *testing.T) {
	if os.Getenv("GENERATE_APPLIANCE_VMX_AUTHENTICATION_TOKEN") == "" {
		t.Skip("skipping test, set environment variable GENERATE_APPLIANCE_VMX_AUTHENTICATION_TOKEN")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMerakiGenerateApplianceVMXAuthenticationTokenPrerequisitesConfig + testAccMerakiGenerateApplianceVMXAuthenticationTokenConfig(),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiGenerateApplianceVMXAuthenticationTokenPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfig

// testAccMerakiGenerateApplianceVMXAuthenticationTokenConfig only sets the id/reference/mandatory top-level
// attributes. Nested (List/Set/Map) attributes are not yet supported in the generated test
// config; add them by hand outside this marked section if a future action needs them.
func testAccMerakiGenerateApplianceVMXAuthenticationTokenConfig() string {
	config := `action "meraki_generate_appliance_vmx_authentication_token" "test" {` + "\n"
	config += `  config {` + "\n"
	config += `    serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  }` + "\n"
	config += `}`
	return config
}

// End of section. //template:end testAccConfig
