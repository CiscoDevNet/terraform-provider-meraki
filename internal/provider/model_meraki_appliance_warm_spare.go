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

type ApplianceWarmSpare struct {
	Id          types.String `tfsdk:"id"`
	NetworkId   types.String `tfsdk:"network_id"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	SpareSerial types.String `tfsdk:"spare_serial"`
	UplinkMode  types.String `tfsdk:"uplink_mode"`
	VirtualIp1  types.String `tfsdk:"virtual_ip1"`
	VirtualIp2  types.String `tfsdk:"virtual_ip2"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceWarmSpare) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/warmSpare", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceWarmSpare) toBody(ctx context.Context, state ApplianceWarmSpare) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.SpareSerial.IsNull() {
		body, _ = sjson.Set(body, "spareSerial", data.SpareSerial.ValueString())
	}
	if !data.UplinkMode.IsNull() {
		body, _ = sjson.Set(body, "uplinkMode", data.UplinkMode.ValueString())
	}
	if !data.VirtualIp1.IsNull() {
		body, _ = sjson.Set(body, "virtualIp1", data.VirtualIp1.ValueString())
	}
	if !data.VirtualIp2.IsNull() {
		body, _ = sjson.Set(body, "virtualIp2", data.VirtualIp2.ValueString())
	}
	return body
}

// End of section. //template:end toBody

func (data *ApplianceWarmSpare) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("spareSerial"); value.Exists() && value.Value() != nil {
		data.SpareSerial = types.StringValue(value.String())
	} else {
		data.SpareSerial = types.StringNull()
	}
	if value := res.Get("uplinkMode"); value.Exists() && value.Value() != nil {
		data.UplinkMode = types.StringValue(value.String())
	} else {
		data.UplinkMode = types.StringNull()
	}
	if value := res.Get("wan1.ip"); value.Exists() && value.Value() != nil {
		data.VirtualIp1 = types.StringValue(value.String())
	} else {
		data.VirtualIp1 = types.StringNull()
	}
	if value := res.Get("wan2.ip"); value.Exists() && value.Value() != nil {
		data.VirtualIp2 = types.StringValue(value.String())
	} else {
		data.VirtualIp2 = types.StringNull()
	}
}

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceWarmSpare) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("spareSerial"); value.Exists() && !data.SpareSerial.IsNull() {
		data.SpareSerial = types.StringValue(value.String())
	} else {
		data.SpareSerial = types.StringNull()
	}
	if value := res.Get("uplinkMode"); value.Exists() && !data.UplinkMode.IsNull() {
		data.UplinkMode = types.StringValue(value.String())
	} else {
		data.UplinkMode = types.StringNull()
	}
	if value := res.Get("wan1.ip"); value.Exists() && !data.VirtualIp1.IsNull() {
		data.VirtualIp1 = types.StringValue(value.String())
	} else {
		data.VirtualIp1 = types.StringNull()
	}
	if value := res.Get("wan2.ip"); value.Exists() && !data.VirtualIp2.IsNull() {
		data.VirtualIp2 = types.StringValue(value.String())
	} else {
		data.VirtualIp2 = types.StringNull()
	}
}
