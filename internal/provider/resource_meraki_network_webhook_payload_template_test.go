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

func TestAccMerakiNetworkWebhookPayloadTemplate(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_webhook_payload_template.test", "body", "{\"event_type\":\"{{alertTypeId}}\",\"client_payload\":{\"text\":\"{{alertData}}\"}}"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_webhook_payload_template.test", "name", "Custom Template"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_webhook_payload_template.test", "headers.0.name", "Authorization"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_webhook_payload_template.test", "headers.0.template", "Bearer {{sharedSecret}}"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkWebhookPayloadTemplatePrerequisitesConfig + testAccMerakiNetworkWebhookPayloadTemplateConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkWebhookPayloadTemplatePrerequisitesConfig + testAccMerakiNetworkWebhookPayloadTemplateConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_webhook_payload_template.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkWebhookPayloadTemplateImportStateIdFunc("meraki_network_webhook_payload_template.test"),
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

func merakiNetworkWebhookPayloadTemplateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s", NetworkId, id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkWebhookPayloadTemplatePrerequisitesConfig = `
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

func testAccMerakiNetworkWebhookPayloadTemplateConfig_minimum() string {
	config := `resource "meraki_network_webhook_payload_template" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  body = "{}"` + "\n"
	config += `  name = "Custom Template"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkWebhookPayloadTemplateConfig_all() string {
	config := `resource "meraki_network_webhook_payload_template" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  body = "{\"event_type\":\"{{alertTypeId}}\",\"client_payload\":{\"text\":\"{{alertData}}\"}}"` + "\n"
	config += `  name = "Custom Template"` + "\n"
	config += `  headers = [{` + "\n"
	config += `    name = "Authorization"` + "\n"
	config += `    template = "Bearer {{sharedSecret}}"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
