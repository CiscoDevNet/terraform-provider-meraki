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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDIdentityPSK struct {
	Id            types.String `tfsdk:"id"`
	NetworkId     types.String `tfsdk:"network_id"`
	Number        types.String `tfsdk:"number"`
	ExpiresAt     types.String `tfsdk:"expires_at"`
	GroupPolicyId types.String `tfsdk:"group_policy_id"`
	Name          types.String `tfsdk:"name"`
	Passphrase    types.String `tfsdk:"passphrase"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDIdentityPSK) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/identityPsks", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDIdentityPSK) toBody(ctx context.Context, state WirelessSSIDIdentityPSK) string {
	body := ""
	if !data.ExpiresAt.IsNull() {
		body, _ = sjson.Set(body, "expiresAt", data.ExpiresAt.ValueString())
	}
	if !data.GroupPolicyId.IsNull() {
		body, _ = sjson.Set(body, "groupPolicyId", data.GroupPolicyId.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Passphrase.IsNull() {
		body, _ = sjson.Set(body, "passphrase", data.Passphrase.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDIdentityPSK) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("expiresAt"); value.Exists() && value.Value() != nil {
		data.ExpiresAt = types.StringValue(value.String())
	} else {
		data.ExpiresAt = types.StringNull()
	}
	if value := res.Get("groupPolicyId"); value.Exists() && value.Value() != nil {
		data.GroupPolicyId = types.StringValue(value.String())
	} else {
		data.GroupPolicyId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("passphrase"); value.Exists() && value.Value() != nil {
		data.Passphrase = types.StringValue(value.String())
	} else {
		data.Passphrase = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDIdentityPSK) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("expiresAt"); value.Exists() && !data.ExpiresAt.IsNull() {
		data.ExpiresAt = types.StringValue(value.String())
	} else {
		data.ExpiresAt = types.StringNull()
	}
	if value := res.Get("groupPolicyId"); value.Exists() && !data.GroupPolicyId.IsNull() {
		data.GroupPolicyId = types.StringValue(value.String())
	} else {
		data.GroupPolicyId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("passphrase"); value.Exists() && !data.Passphrase.IsNull() {
		data.Passphrase = types.StringValue(value.String())
	} else {
		data.Passphrase = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDIdentityPSK) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data WirelessSSIDIdentityPSK) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
