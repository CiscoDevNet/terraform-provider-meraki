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

func TestAccMerakiApplianceThirdPartyVPNPeers(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ike_version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.local_id", "myMXId@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.name", "Peer Name"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.public_ip", "123.123.123.1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.remote_id", "miles@meraki.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.secret", "Sample Password"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_child_lifetime", "28800"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_ike_lifetime", "28800"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_child_auth_algo.0", "sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_child_cipher_algo.0", "aes128"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_child_pfs_group.0", "disabled"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_ike_auth_algo.0", "sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_ike_cipher_algo.0", "tripledes"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_ike_diffie_hellman_group.0", "group2"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.ipsec_policies_ike_prf_algo.0", "prfsha1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.network_tags.0", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_third_party_vpn_peers.test", "peers.0.private_subnets.0", "192.168.1.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceThirdPartyVPNPeersPrerequisitesConfig + testAccMerakiApplianceThirdPartyVPNPeersConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceThirdPartyVPNPeersPrerequisitesConfig + testAccMerakiApplianceThirdPartyVPNPeersConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_third_party_vpn_peers.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceThirdPartyVPNPeersImportStateIdFunc("meraki_appliance_third_party_vpn_peers.test"),
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

func merakiApplianceThirdPartyVPNPeersImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]

		return fmt.Sprintf("%s", OrganizationId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceThirdPartyVPNPeersPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceThirdPartyVPNPeersConfig_minimum() string {
	config := `resource "meraki_appliance_third_party_vpn_peers" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	peers = [{` + "\n"
	config += `		name = "Peer Name"` + "\n"
	config += `		public_ip = "123.123.1.1"` + "\n"
	config += `		secret = "Sample Password"` + "\n"
	config += `		private_subnets = ["192.168.1.0/24"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceThirdPartyVPNPeersConfig_all() string {
	config := `resource "meraki_appliance_third_party_vpn_peers" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	peers = [{` + "\n"
	config += `		ike_version = "2"` + "\n"
	config += `		local_id = "myMXId@meraki.com"` + "\n"
	config += `		name = "Peer Name"` + "\n"
	config += `		public_ip = "123.123.123.1"` + "\n"
	config += `		remote_id = "miles@meraki.com"` + "\n"
	config += `		secret = "Sample Password"` + "\n"
	config += `		ipsec_policies_child_lifetime = 28800` + "\n"
	config += `		ipsec_policies_ike_lifetime = 28800` + "\n"
	config += `		ipsec_policies_child_auth_algo = ["sha1"]` + "\n"
	config += `		ipsec_policies_child_cipher_algo = ["aes128"]` + "\n"
	config += `		ipsec_policies_child_pfs_group = ["disabled"]` + "\n"
	config += `		ipsec_policies_ike_auth_algo = ["sha1"]` + "\n"
	config += `		ipsec_policies_ike_cipher_algo = ["tripledes"]` + "\n"
	config += `		ipsec_policies_ike_diffie_hellman_group = ["group2"]` + "\n"
	config += `		ipsec_policies_ike_prf_algo = ["prfsha1"]` + "\n"
	config += `		network_tags = ["none"]` + "\n"
	config += `		private_subnets = ["192.168.1.0/24"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
