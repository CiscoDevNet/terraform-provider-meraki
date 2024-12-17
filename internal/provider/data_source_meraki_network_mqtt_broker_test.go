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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiNetworkMQTTBroker(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "host", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "name", "MQTT_Broker_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "authentication_username", "milesmeraki"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "security_mode", "tls"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_mqtt_broker.test", "security_tls_verify_hostnames", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkMQTTBrokerPrerequisitesConfig + testAccDataSourceMerakiNetworkMQTTBrokerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkMQTTBrokerPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkMQTTBrokerConfig() string {
	config := `resource "meraki_network_mqtt_broker" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  host = "1.2.3.4"` + "\n"
	config += `  name = "MQTT_Broker_1"` + "\n"
	config += `  port = 443` + "\n"
	config += `  authentication_password = "*****"` + "\n"
	config += `  authentication_username = "milesmeraki"` + "\n"
	config += `  security_mode = "tls"` + "\n"
	config += `  security_tls_ca_certificate = "*****"` + "\n"
	config += `  security_tls_verify_hostnames = true` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_mqtt_broker" "test" {
			id = meraki_network_mqtt_broker.test.id
			network_id = meraki_network.test.id
			depends_on = [meraki_network_mqtt_broker.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiNetworkMQTTBrokerConfig() string {
	config := `resource "meraki_network_mqtt_broker" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  host = "1.2.3.4"` + "\n"
	config += `  name = "MQTT_Broker_1"` + "\n"
	config += `  port = 443` + "\n"
	config += `  authentication_password = "*****"` + "\n"
	config += `  authentication_username = "milesmeraki"` + "\n"
	config += `  security_mode = "tls"` + "\n"
	config += `  security_tls_ca_certificate = "*****"` + "\n"
	config += `  security_tls_verify_hostnames = true` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_mqtt_broker" "test" {
			name = meraki_network_mqtt_broker.test.name
			network_id = meraki_network.test.id
			depends_on = [meraki_network_mqtt_broker.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
