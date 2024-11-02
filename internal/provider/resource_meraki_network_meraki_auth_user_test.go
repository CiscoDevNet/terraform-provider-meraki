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

func TestAccMerakiNetworkMerakiAuthUser(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_meraki_auth_user.test", "account_type", "802.1X"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_meraki_auth_user.test", "email", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_meraki_auth_user.test", "is_admin", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_meraki_auth_user.test", "name", "Miles Meraki"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_meraki_auth_user.test", "authorizations.0.expires_at", "2018-03-13T00:00:00.090210Z"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkMerakiAuthUserPrerequisitesConfig + testAccMerakiNetworkMerakiAuthUserConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkMerakiAuthUserPrerequisitesConfig + testAccMerakiNetworkMerakiAuthUserConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_meraki_auth_user.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkMerakiAuthUserImportStateIdFunc("meraki_network_meraki_auth_user.test"),
		ImportStateVerifyIgnore: []string{"email_password_to_user", "password"},
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

func merakiNetworkMerakiAuthUserImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkMerakiAuthUserPrerequisitesConfig = `
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
  network_id          = meraki_network.test.id
  number              = "1"
  name                = "My SSID"
  auth_mode           = "8021x-meraki"
  wpa_encryption_mode = "WPA2 only"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiNetworkMerakiAuthUserConfig_minimum() string {
	config := `resource "meraki_network_meraki_auth_user" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	email = "miles@meraki.com"` + "\n"
	config += `	name = "Miles Meraki"` + "\n"
	config += `	password = "Cisco123&"` + "\n"
	config += `	authorizations = [{` + "\n"
	config += `		ssid_number = meraki_wireless_ssid.test.number` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkMerakiAuthUserConfig_all() string {
	config := `resource "meraki_network_meraki_auth_user" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	account_type = "802.1X"` + "\n"
	config += `	email = "miles@meraki.com"` + "\n"
	config += `	email_password_to_user = false` + "\n"
	config += `	is_admin = false` + "\n"
	config += `	name = "Miles Meraki"` + "\n"
	config += `	password = "Cisco123&"` + "\n"
	config += `	authorizations = [{` + "\n"
	config += `		expires_at = "2018-03-13T00:00:00.090210Z"` + "\n"
	config += `		ssid_number = meraki_wireless_ssid.test.number` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
