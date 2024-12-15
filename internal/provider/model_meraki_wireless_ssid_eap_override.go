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

type WirelessSSIDEAPOverride struct {
	Id                  types.String `tfsdk:"id"`
	NetworkId           types.String `tfsdk:"network_id"`
	Number              types.String `tfsdk:"number"`
	MaxRetries          types.Int64  `tfsdk:"max_retries"`
	Timeout             types.Int64  `tfsdk:"timeout"`
	EapolKeyRetries     types.Int64  `tfsdk:"eapol_key_retries"`
	EapolKeyTimeoutInMs types.Int64  `tfsdk:"eapol_key_timeout_in_ms"`
	IdentityRetries     types.Int64  `tfsdk:"identity_retries"`
	IdentityTimeout     types.Int64  `tfsdk:"identity_timeout"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDEAPOverride) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/eapOverride", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDEAPOverride) toBody(ctx context.Context, state WirelessSSIDEAPOverride) string {
	body := ""
	if !data.MaxRetries.IsNull() {
		body, _ = sjson.Set(body, "maxRetries", data.MaxRetries.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, "timeout", data.Timeout.ValueInt64())
	}
	if !data.EapolKeyRetries.IsNull() {
		body, _ = sjson.Set(body, "eapolKey.retries", data.EapolKeyRetries.ValueInt64())
	}
	if !data.EapolKeyTimeoutInMs.IsNull() {
		body, _ = sjson.Set(body, "eapolKey.timeoutInMs", data.EapolKeyTimeoutInMs.ValueInt64())
	}
	if !data.IdentityRetries.IsNull() {
		body, _ = sjson.Set(body, "identity.retries", data.IdentityRetries.ValueInt64())
	}
	if !data.IdentityTimeout.IsNull() {
		body, _ = sjson.Set(body, "identity.timeout", data.IdentityTimeout.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDEAPOverride) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("maxRetries"); value.Exists() && value.Value() != nil {
		data.MaxRetries = types.Int64Value(value.Int())
	} else {
		data.MaxRetries = types.Int64Null()
	}
	if value := res.Get("timeout"); value.Exists() && value.Value() != nil {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get("eapolKey.retries"); value.Exists() && value.Value() != nil {
		data.EapolKeyRetries = types.Int64Value(value.Int())
	} else {
		data.EapolKeyRetries = types.Int64Null()
	}
	if value := res.Get("eapolKey.timeoutInMs"); value.Exists() && value.Value() != nil {
		data.EapolKeyTimeoutInMs = types.Int64Value(value.Int())
	} else {
		data.EapolKeyTimeoutInMs = types.Int64Null()
	}
	if value := res.Get("identity.retries"); value.Exists() && value.Value() != nil {
		data.IdentityRetries = types.Int64Value(value.Int())
	} else {
		data.IdentityRetries = types.Int64Null()
	}
	if value := res.Get("identity.timeout"); value.Exists() && value.Value() != nil {
		data.IdentityTimeout = types.Int64Value(value.Int())
	} else {
		data.IdentityTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDEAPOverride) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("maxRetries"); value.Exists() && !data.MaxRetries.IsNull() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else {
		data.MaxRetries = types.Int64Null()
	}
	if value := res.Get("timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get("eapolKey.retries"); value.Exists() && !data.EapolKeyRetries.IsNull() {
		data.EapolKeyRetries = types.Int64Value(value.Int())
	} else {
		data.EapolKeyRetries = types.Int64Null()
	}
	if value := res.Get("eapolKey.timeoutInMs"); value.Exists() && !data.EapolKeyTimeoutInMs.IsNull() {
		data.EapolKeyTimeoutInMs = types.Int64Value(value.Int())
	} else {
		data.EapolKeyTimeoutInMs = types.Int64Null()
	}
	if value := res.Get("identity.retries"); value.Exists() && !data.IdentityRetries.IsNull() {
		data.IdentityRetries = types.Int64Value(value.Int())
	} else {
		data.IdentityRetries = types.Int64Null()
	}
	if value := res.Get("identity.timeout"); value.Exists() && !data.IdentityTimeout.IsNull() {
		data.IdentityTimeout = types.Int64Value(value.Int())
	} else {
		data.IdentityTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDEAPOverride) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
