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

func TestAccMerakiNetworkWirelessRadioRRM(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "ai_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "busy_hour_minimize_changes_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "busy_hour_schedule_mode", "automatic"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "busy_hour_schedule_manual_end", "15:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "busy_hour_schedule_manual_start", "10:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "channel_avoidance_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_wireless_radio_rrm.test", "fra_enabled", "false"))

	var steps []resource.TestStep
	var tfVersion *goversion.Version
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkWirelessRadioRRMPrerequisitesConfig + testAccMerakiNetworkWirelessRadioRRMConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkWirelessRadioRRMPrerequisitesConfig + testAccMerakiNetworkWirelessRadioRRMConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

func merakiNetworkWirelessRadioRRMImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkWirelessRadioRRMPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkWirelessRadioRRMConfig_minimum() string {
	config := `resource "meraki_network_wireless_radio_rrm" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  ai_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccMerakiNetworkWirelessRadioRRMConfig_all() string {
	config := `resource "meraki_network_wireless_radio_rrm" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  ai_enabled = true` + "\n"
	config += `  busy_hour_minimize_changes_enabled = true` + "\n"
	config += `  busy_hour_schedule_mode = "automatic"` + "\n"
	config += `  busy_hour_schedule_manual_end = "15:00"` + "\n"
	config += `  busy_hour_schedule_manual_start = "10:00"` + "\n"
	config += `  channel_avoidance_enabled = true` + "\n"
	config += `  fra_enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
