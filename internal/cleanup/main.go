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

// Cleanup script for CI test environments.
// Deletes leftover test resources from the Meraki Dashboard organization
// to ensure a clean state before running acceptance tests.
//
// Required environment variables:
//   - MERAKI_API_KEY: Meraki Dashboard API key
//   - TF_VAR_test_org: Name of the test organization
//   - TF_VAR_test_network: Name of the test network
//
// Usage: go run internal/cleanup/main.go

package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// resource defines an organization-level resource to clean up.
type resource struct {
	name      string // Human-readable name for logging
	path      string // API path relative to /organizations/{orgId}
	idField   string // JSON field name containing the resource ID
	skipField string // If set, skip items where this boolean JSON field is true (e.g. "isDefaultGroup")
}

// phase1Resources must be deleted before phase2Resources due to dependencies.
var phase1Resources = []resource{
	{name: "Adaptive Policy", path: "/adaptivePolicy/policies", idField: "adaptivePolicyId"},
	{name: "Switch Ports Profiles Automation", path: "/switch/ports/profiles/automations", idField: "id"},
	{name: "Policy Object Group", path: "/policyObjects/groups", idField: "id"},
	{name: "DNS Local Record", path: "/appliance/dns/local/records", idField: "recordId"},
}

// phase2Resources can be deleted after phase1Resources are gone.
var phase2Resources = []resource{
	{name: "Adaptive Policy ACL", path: "/adaptivePolicy/acls", idField: "aclId"},
	{name: "Adaptive Policy Group", path: "/adaptivePolicy/groups", idField: "groupId", skipField: "isDefaultGroup"},
	{name: "Alerts Profile", path: "/alerts/profiles", idField: "id"},
	{name: "Auth RADIUS Server", path: "/auth/radius/servers", idField: "serverId"},
	{name: "Config Template", path: "/configTemplates", idField: "id"},
	{name: "Policy Object", path: "/policyObjects", idField: "id"},
	{name: "SAML IdP", path: "/saml/idps", idField: "idpId"},
	{name: "SAML Role", path: "/samlRoles", idField: "id"},
	{name: "Camera Role", path: "/camera/roles", idField: "id"},
	{name: "SM Admin Role", path: "/sm/admins/roles", idField: "roleId"},
	{name: "Switch Ports Profile", path: "/switch/ports/profiles", idField: "profileId"},
	{name: "Wireless Location Scanning Receiver", path: "/wireless/location/scanning/receivers", idField: "receiverId"},
	{name: "Wireless SSID Firewall Isolation Allowlist Entry", path: "/wireless/ssids/firewall/isolation/allowlist/entries", idField: "entryId"},
	{name: "DNS Local Profile", path: "/appliance/dns/local/profiles", idField: "profileId"},
	{name: "DNS Split Profile", path: "/appliance/dns/split/profiles", idField: "profileId"},
}

func main() {
	apiKey := os.Getenv("MERAKI_API_KEY")
	orgName := os.Getenv("TF_VAR_test_org")
	networkName := os.Getenv("TF_VAR_test_network")

	if apiKey == "" {
		fmt.Println("ERROR: MERAKI_API_KEY environment variable is required")
		os.Exit(1)
	}
	if orgName == "" {
		fmt.Println("ERROR: TF_VAR_test_org environment variable is required")
		os.Exit(1)
	}
	if networkName == "" {
		fmt.Println("ERROR: TF_VAR_test_network environment variable is required")
		os.Exit(1)
	}

	client, err := meraki.NewClient(apiKey, meraki.MaxRetries(3))
	if err != nil {
		fmt.Printf("ERROR: Failed to create Meraki client: %s\n", err)
		os.Exit(1)
	}

	// Look up organization ID by name
	orgId, err := getOrganizationId(client, orgName)
	if err != nil {
		fmt.Printf("ERROR: Failed to find organization '%s': %s\n", orgName, err)
		os.Exit(1)
	}
	fmt.Printf("Found organization '%s' with ID: %s\n", orgName, orgId)

	fmt.Println("\n========================================")
	fmt.Println("Phase 1: Deleting dependent resources...")
	fmt.Println("========================================")
	for _, r := range phase1Resources {
		cleanupResource(client, orgId, r)
	}

	fmt.Println("\n========================================")
	fmt.Println("Phase 2: Deleting remaining resources...")
	fmt.Println("========================================")
	for _, r := range phase2Resources {
		cleanupResource(client, orgId, r)
	}

	fmt.Println("\n========================================")
	fmt.Println("Phase 3: Deleting test network(s)...")
	fmt.Println("========================================")
	cleanupNetworks(client, orgId, networkName)

	fmt.Println("\nCleanup complete.")
}

// getOrganizationId looks up an organization by name and returns its ID.
func getOrganizationId(client meraki.Client, name string) (string, error) {
	res, err := client.Get("/organizations")
	if err != nil {
		return "", fmt.Errorf("failed to list organizations: %w", err)
	}

	var orgId string
	res.ForEach(func(k, v gjson.Result) bool {
		if v.Get("name").String() == name {
			orgId = v.Get("id").String()
			return false
		}
		return true
	})

	if orgId == "" {
		return "", fmt.Errorf("organization not found")
	}
	return orgId, nil
}

// cleanupResource lists all instances of a resource type and deletes them.
func cleanupResource(client meraki.Client, orgId string, r resource) {
	listPath := fmt.Sprintf("/organizations/%s%s", url.QueryEscape(orgId), r.path)

	res, err := client.Get(listPath)
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			return
		}
		fmt.Printf("  [%s] Warning: failed to list: %s\n", r.name, err)
		return
	}

	// Some endpoints wrap results in an "items" array
	items := res.Result
	if res.Get("items").Exists() {
		items = res.Get("items")
	}

	if !items.IsArray() || len(items.Array()) == 0 {
		fmt.Printf("  [%s] No resources found\n", r.name)
		return
	}

	count := 0
	items.ForEach(func(k, v gjson.Result) bool {
		id := v.Get(r.idField).String()
		if id == "" {
			fmt.Printf("  [%s] Warning: resource missing '%s' field, skipping\n", r.name, r.idField)
			return true
		}

		// Skip built-in/default resources that cannot be deleted
		if r.skipField != "" && v.Get(r.skipField).Bool() {
			return true
		}

		deletePath := fmt.Sprintf("%s/%s", listPath, url.QueryEscape(id))
		_, err := client.Delete(deletePath)
		if err != nil {
			if !strings.Contains(err.Error(), "StatusCode 404") {
				fmt.Printf("  [%s] Warning: failed to delete %s: %s\n", r.name, id, err)
			}
		} else {
			fmt.Printf("  [%s] Deleted: %s\n", r.name, id)
			count++
		}
		return true
	})

	if count > 0 {
		fmt.Printf("  [%s] Deleted %d resource(s)\n", r.name, count)
	}
}

// cleanupNetworks finds and deletes test networks by name.
func cleanupNetworks(client meraki.Client, orgId, networkName string) {
	listPath := fmt.Sprintf("/organizations/%s/networks", url.QueryEscape(orgId))

	res, err := client.Get(listPath)
	if err != nil {
		fmt.Printf("  [Network] Warning: failed to list networks: %s\n", err)
		return
	}

	// Some endpoints wrap results in an "items" array
	items := res.Result
	if res.Get("items").Exists() {
		items = res.Get("items")
	}

	count := 0
	items.ForEach(func(k, v gjson.Result) bool {
		name := v.Get("name").String()
		id := v.Get("id").String()

		if name == networkName {
			deletePath := fmt.Sprintf("%s/%s", listPath, url.QueryEscape(id))
			_, err := client.Delete(deletePath)
			if err != nil {
				if !strings.Contains(err.Error(), "StatusCode 404") {
					fmt.Printf("  [Network] Warning: failed to delete '%s' (%s): %s\n", name, id, err)
				}
			} else {
				fmt.Printf("  [Network] Deleted: '%s' (%s)\n", name, id)
				count++
			}
		}
		return true
	})

	if count == 0 {
		fmt.Printf("  [Network] No networks found with name '%s'\n", networkName)
	} else {
		fmt.Printf("  [Network] Deleted %d network(s)\n", count)
	}
}
