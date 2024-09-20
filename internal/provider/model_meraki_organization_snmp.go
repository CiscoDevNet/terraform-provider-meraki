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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationSNMP struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	V2cEnabled     types.Bool   `tfsdk:"v2c_enabled"`
	V3AuthMode     types.String `tfsdk:"v3_auth_mode"`
	V3AuthPass     types.String `tfsdk:"v3_auth_pass"`
	V3Enabled      types.Bool   `tfsdk:"v3_enabled"`
	V3PrivMode     types.String `tfsdk:"v3_priv_mode"`
	V3PrivPass     types.String `tfsdk:"v3_priv_pass"`
	PeerIps        types.List   `tfsdk:"peer_ips"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationSNMP) getPath() string {
	return fmt.Sprintf("/organizations/%v/snmp", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationSNMP) toBody(ctx context.Context, state OrganizationSNMP) string {
	body := ""
	if !data.V2cEnabled.IsNull() {
		body, _ = sjson.Set(body, "v2cEnabled", data.V2cEnabled.ValueBool())
	}
	if !data.V3AuthMode.IsNull() {
		body, _ = sjson.Set(body, "v3AuthMode", data.V3AuthMode.ValueString())
	}
	if !data.V3AuthPass.IsNull() {
		body, _ = sjson.Set(body, "v3AuthPass", data.V3AuthPass.ValueString())
	}
	if !data.V3Enabled.IsNull() {
		body, _ = sjson.Set(body, "v3Enabled", data.V3Enabled.ValueBool())
	}
	if !data.V3PrivMode.IsNull() {
		body, _ = sjson.Set(body, "v3PrivMode", data.V3PrivMode.ValueString())
	}
	if !data.V3PrivPass.IsNull() {
		body, _ = sjson.Set(body, "v3PrivPass", data.V3PrivPass.ValueString())
	}
	if !data.PeerIps.IsNull() {
		var values []string
		data.PeerIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "peerIps", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationSNMP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("v2cEnabled"); value.Exists() {
		data.V2cEnabled = types.BoolValue(value.Bool())
	} else {
		data.V2cEnabled = types.BoolNull()
	}
	if value := res.Get("v3AuthMode"); value.Exists() {
		data.V3AuthMode = types.StringValue(value.String())
	} else {
		data.V3AuthMode = types.StringNull()
	}
	if value := res.Get("v3Enabled"); value.Exists() {
		data.V3Enabled = types.BoolValue(value.Bool())
	} else {
		data.V3Enabled = types.BoolNull()
	}
	if value := res.Get("v3PrivMode"); value.Exists() {
		data.V3PrivMode = types.StringValue(value.String())
	} else {
		data.V3PrivMode = types.StringNull()
	}
	if value := res.Get("peerIps"); value.Exists() {
		data.PeerIps = helpers.GetStringList(value.Array())
	} else {
		data.PeerIps = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationSNMP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("v2cEnabled"); value.Exists() && !data.V2cEnabled.IsNull() {
		data.V2cEnabled = types.BoolValue(value.Bool())
	} else {
		data.V2cEnabled = types.BoolNull()
	}
	if value := res.Get("v3AuthMode"); value.Exists() && !data.V3AuthMode.IsNull() {
		data.V3AuthMode = types.StringValue(value.String())
	} else {
		data.V3AuthMode = types.StringNull()
	}
	if value := res.Get("v3Enabled"); value.Exists() && !data.V3Enabled.IsNull() {
		data.V3Enabled = types.BoolValue(value.Bool())
	} else {
		data.V3Enabled = types.BoolNull()
	}
	if value := res.Get("v3PrivMode"); value.Exists() && !data.V3PrivMode.IsNull() {
		data.V3PrivMode = types.StringValue(value.String())
	} else {
		data.V3PrivMode = types.StringNull()
	}
	if value := res.Get("peerIps"); value.Exists() && !data.PeerIps.IsNull() {
		data.PeerIps = helpers.GetStringList(value.Array())
	} else {
		data.PeerIps = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
