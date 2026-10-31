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

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tidwall/sjson"
)

// Resources whose Read first gets the network (to find its organization) and then reads the settings
// from an organization-level endpoint. A deleted network already returns 404 on the first GET, which
// removes the resource from the state. A 404 from the organization endpoint does not mean that the
// settings are gone: for Air Marshal settings, the API returned 404 for an organization the key cannot
// access, and 200 with an empty list for a network that does not exist. The resource must then stay in
// the state and Read must return an error instead of planning a re-create on every run. A 400 from
// the organization endpoint still removes the resource, as before.
var orgReadResources = []struct {
	name      string
	config    string
	writePath string
	orgPath   string
	// orgBody returns the organization endpoint response for the last body written to writePath
	orgBody func(written string) string
}{
	{
		name: "meraki_appliance_firewall_multicast_forwarding",
		config: `
  rules = [{
    address     = "224.0.0.1"
    description = "test"
    vlan_ids    = ["1"]
  }]`,
		writePath: "/networks/N_1/appliance/firewall/multicastForwarding",
		orgPath:   "/organizations/123/appliance/firewall/multicastForwarding/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "network.id", "N_1") },
	},
	{
		name: "meraki_appliance_traffic_shaping_vpn_exclusions",
		config: `
  custom = [{
    destination = "192.168.3.0/24"
    port        = "8000"
    protocol    = "tcp"
  }]`,
		writePath: "/networks/N_1/appliance/trafficShaping/vpnExclusions",
		orgPath:   "/organizations/123/appliance/trafficShaping/vpnExclusions/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "networkId", "N_1") },
	},
	{
		name: "meraki_network_wireless_radio_rrm",
		config: `
  ai_enabled = true`,
		writePath: "/networks/N_1/wireless/radio/rrm",
		orgPath:   "/organizations/123/wireless/radio/rrm/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "networkId", "N_1") },
	},
	{
		name: "meraki_wireless_air_marshal_rule",
		config: `
  type         = "allow"
  match_string = "00:11:22:33:44:55"
  match_type   = "bssid"`,
		writePath: "/networks/N_1/wireless/airMarshal/rules",
		orgPath:   "/organizations/123/wireless/airMarshal/rules",
		orgBody:   func(w string) string { return orgItems(w, "ruleId", "R_1") },
	},
	{
		name: "meraki_wireless_air_marshal_settings",
		config: `
  default_policy = "allow"`,
		writePath: "/networks/N_1/wireless/airMarshal/settings",
		orgPath:   "/organizations/123/wireless/airMarshal/settings/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "networkId", "N_1") },
	},
	{
		name: "meraki_wireless_location_scanning",
		config: `
  enabled     = true
  api_enabled = false`,
		writePath: "/networks/N_1/wireless/location/scanning",
		orgPath:   "/organizations/123/wireless/location/scanning/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "networkId", "N_1") },
	},
	{
		name: "meraki_wireless_ssid_open_roaming",
		config: `
  number    = "0"
  enabled   = false
  tenant_id = "12345"`,
		writePath: "/networks/N_1/wireless/ssids/0/openRoaming",
		orgPath:   "/organizations/123/wireless/ssids/openRoaming/byNetwork",
		orgBody: func(w string) string {
			return fmt.Sprintf(`{"items":[{"networkId":"N_1","ssids":[{"number":0,"openRoaming":%s}]}]}`, w)
		},
	},
	{
		name: "meraki_wireless_zigbee",
		config: `
  enabled                       = false
  defaults_channel              = "25"
  defaults_transmit_power_level = 10`,
		writePath: "/networks/N_1/wireless/zigbee",
		orgPath:   "/organizations/123/wireless/zigbee/byNetwork",
		orgBody:   func(w string) string { return orgItems(w, "networkId", "N_1") },
	},
}

func TestUnitMerakiOrganizationReadErrors(t *testing.T) {
	for _, r := range orgReadResources {
		t.Run(r.name, func(t *testing.T) {
			api := newFakeOrgReadAPI(t, r.writePath, r.orgPath, r.orgBody)
			srv := httptest.NewServer(api)
			defer srv.Close()
			if !strings.HasPrefix(srv.URL, "http://127.0.0.1:") {
				t.Fatalf("fake API must be local, got %s", srv.URL)
			}
			config := fmt.Sprintf(`
provider "meraki" {
  base_url = "%s/api/v1"
  api_key  = "dummy-api-key"
  retries  = 0
}

resource %q "test" {
  network_id = "N_1"%s
}
`, srv.URL, r.name, r.config)

			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					// Both GETs succeed: create and refresh converge
					{Config: config},
					// The organization endpoint is not accessible: Read fails
					{
						PreConfig:   func() { api.set(http.StatusOK, http.StatusNotFound) },
						Config:      config,
						ExpectError: regexp.MustCompile(`StatusCode\s+404`),
					},
					// Access restored: the resource is still in the state, so nothing changes
					{
						PreConfig: func() { api.set(http.StatusOK, http.StatusOK) },
						Config:    config,
						PlanOnly:  true,
					},
					// A 400 from the organization endpoint still removes the resource from the state (unchanged)
					{
						PreConfig:          func() { api.set(http.StatusOK, http.StatusBadRequest) },
						Config:             config,
						PlanOnly:           true,
						ExpectNonEmptyPlan: true,
					},
					// The network is gone (404 on the first GET): the resource is still removed from the state
					{
						PreConfig:          func() { api.set(http.StatusNotFound, http.StatusOK) },
						Config:             config,
						PlanOnly:           true,
						ExpectNonEmptyPlan: true,
					},
				},
			})
		})
	}
}

// orgItems wraps the written settings as the only item of an organization-level list response
func orgItems(written, key, value string) string {
	item, _ := sjson.Set(written, key, value)
	return `{"items":[` + item + `]}`
}

type fakeOrgReadAPI struct {
	t         *testing.T
	writePath string
	orgPath   string
	orgBody   func(string) string

	mu            sync.Mutex
	networkStatus int
	orgStatus     int
	written       string
}

func newFakeOrgReadAPI(t *testing.T, writePath, orgPath string, orgBody func(string) string) *fakeOrgReadAPI {
	return &fakeOrgReadAPI{t: t, writePath: writePath, orgPath: orgPath, orgBody: orgBody,
		networkStatus: http.StatusOK, orgStatus: http.StatusOK, written: "{}"}
}

func (f *fakeOrgReadAPI) set(networkStatus, orgStatus int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.networkStatus, f.orgStatus = networkStatus, orgStatus
}

func (f *fakeOrgReadAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	path := strings.TrimPrefix(r.URL.Path, "/api/v1")
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodGet && path == "/networks/N_1":
		if f.networkStatus != http.StatusOK {
			w.WriteHeader(f.networkStatus)
			fmt.Fprint(w, `{"errors":["Not found"]}`)
			return
		}
		fmt.Fprint(w, `{"id":"N_1","organizationId":"123"}`)
	case (r.Method == http.MethodPut || r.Method == http.MethodPost) && path == f.writePath:
		body, _ := io.ReadAll(r.Body)
		f.written = string(body)
		if r.Method == http.MethodPost {
			f.written, _ = sjson.Set(f.written, "ruleId", "R_1")
		}
		fmt.Fprint(w, f.written)
	case (r.Method == http.MethodPut || r.Method == http.MethodDelete) && strings.HasPrefix(path, f.writePath+"/"):
		w.WriteHeader(http.StatusNoContent)
	case r.Method == http.MethodGet && path == f.orgPath:
		if f.orgStatus != http.StatusOK {
			w.WriteHeader(f.orgStatus)
			fmt.Fprint(w, `{"errors":["Not found"]}`)
			return
		}
		fmt.Fprint(w, f.orgBody(f.written))
	default:
		f.t.Errorf("unexpected request: %s %s", r.Method, r.URL.String())
		w.WriteHeader(http.StatusTeapot)
	}
}
