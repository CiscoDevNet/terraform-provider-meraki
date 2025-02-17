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

type ApplianceSettings struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	ClientTrackingMethod types.String `tfsdk:"client_tracking_method"`
	DeploymentMode       types.String `tfsdk:"deployment_mode"`
	DynamicDnsEnabled    types.Bool   `tfsdk:"dynamic_dns_enabled"`
	DynamicDnsPrefix     types.String `tfsdk:"dynamic_dns_prefix"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceSettings) toBody(ctx context.Context, state ApplianceSettings) string {
	body := ""
	if !data.ClientTrackingMethod.IsNull() {
		body, _ = sjson.Set(body, "clientTrackingMethod", data.ClientTrackingMethod.ValueString())
	}
	if !data.DeploymentMode.IsNull() {
		body, _ = sjson.Set(body, "deploymentMode", data.DeploymentMode.ValueString())
	}
	if !data.DynamicDnsEnabled.IsNull() {
		body, _ = sjson.Set(body, "dynamicDns.enabled", data.DynamicDnsEnabled.ValueBool())
	}
	if !data.DynamicDnsPrefix.IsNull() {
		body, _ = sjson.Set(body, "dynamicDns.prefix", data.DynamicDnsPrefix.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("clientTrackingMethod"); value.Exists() && value.Value() != nil {
		data.ClientTrackingMethod = types.StringValue(value.String())
	} else {
		data.ClientTrackingMethod = types.StringNull()
	}
	if value := res.Get("deploymentMode"); value.Exists() && value.Value() != nil {
		data.DeploymentMode = types.StringValue(value.String())
	} else {
		data.DeploymentMode = types.StringNull()
	}
	if value := res.Get("dynamicDns.enabled"); value.Exists() && value.Value() != nil {
		data.DynamicDnsEnabled = types.BoolValue(value.Bool())
	} else {
		data.DynamicDnsEnabled = types.BoolNull()
	}
	if value := res.Get("dynamicDns.prefix"); value.Exists() && value.Value() != nil {
		data.DynamicDnsPrefix = types.StringValue(value.String())
	} else {
		data.DynamicDnsPrefix = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("clientTrackingMethod"); value.Exists() && !data.ClientTrackingMethod.IsNull() {
		data.ClientTrackingMethod = types.StringValue(value.String())
	} else {
		data.ClientTrackingMethod = types.StringNull()
	}
	if value := res.Get("deploymentMode"); value.Exists() && !data.DeploymentMode.IsNull() {
		data.DeploymentMode = types.StringValue(value.String())
	} else {
		data.DeploymentMode = types.StringNull()
	}
	if value := res.Get("dynamicDns.enabled"); value.Exists() && !data.DynamicDnsEnabled.IsNull() {
		data.DynamicDnsEnabled = types.BoolValue(value.Bool())
	} else {
		data.DynamicDnsEnabled = types.BoolNull()
	}
	if value := res.Get("dynamicDns.prefix"); value.Exists() && !data.DynamicDnsPrefix.IsNull() {
		data.DynamicDnsPrefix = types.StringValue(value.String())
	} else {
		data.DynamicDnsPrefix = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
