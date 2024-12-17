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

func TestAccMerakiWirelessSSIDHotspot20(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "network_access_type", "Private network"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "operator_name", "Meraki Product Management"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "venue_name", "SF Branch"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "venue_type", "Unspecified Assembly"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "mcc_mncs.0.mcc", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "mcc_mncs.0.mnc", "456"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "nai_realms.0.format", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "nai_realms.0.realm", ""))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_hotspot_20.test", "nai_realms.0.methods.0.id", "1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessSSIDHotspot20PrerequisitesConfig + testAccMerakiWirelessSSIDHotspot20Config_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDHotspot20PrerequisitesConfig + testAccMerakiWirelessSSIDHotspot20Config_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_ssid_hotspot_20.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessSSIDHotspot20ImportStateIdFunc("meraki_wireless_ssid_hotspot_20.test"),
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

func merakiWirelessSSIDHotspot20ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Number := primary.Attributes["number"]

		return fmt.Sprintf("%s,%s", NetworkId, Number), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessSSIDHotspot20PrerequisitesConfig = `
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
resource "meraki_wireless_ssid" "test" {
  network_id = meraki_network.test.id
  number     = "0"
  name       = "My SSID"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessSSIDHotspot20Config_minimum() string {
	config := `resource "meraki_wireless_ssid_hotspot_20" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `  enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessSSIDHotspot20Config_all() string {
	config := `resource "meraki_wireless_ssid_hotspot_20" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `  enabled = true` + "\n"
	config += `  network_access_type = "Private network"` + "\n"
	config += `  operator_name = "Meraki Product Management"` + "\n"
	config += `  venue_name = "SF Branch"` + "\n"
	config += `  venue_type = "Unspecified Assembly"` + "\n"
	config += `  domains = ["meraki.local"]` + "\n"
	config += `  mcc_mncs = [{` + "\n"
	config += `    mcc = "123"` + "\n"
	config += `    mnc = "456"` + "\n"
	config += `  }]` + "\n"
	config += `  nai_realms = [{` + "\n"
	config += `    format = "1"` + "\n"
	config += `    realm = ""` + "\n"
	config += `    methods = [{` + "\n"
	config += `      id = "1"` + "\n"
	config += `      authentication_types_non_eap_inner_authentication = ["MSCHAPV2"]` + "\n"
	config += `      authentication_types_eap_inner_authentication = ["EAP-TLS"]` + "\n"
	config += `      authentication_types_credentials = ["SIM"]` + "\n"
	config += `      authentication_types_tunneled_eap_method_credentials = ["USIM"]` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `  roam_consort_ois = ["ABC123"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
