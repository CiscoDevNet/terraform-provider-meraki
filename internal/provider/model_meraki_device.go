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

type Device struct {
	Id              types.String  `tfsdk:"id"`
	Serial          types.String  `tfsdk:"serial"`
	Address         types.String  `tfsdk:"address"`
	FloorPlanId     types.String  `tfsdk:"floor_plan_id"`
	Lat             types.Float64 `tfsdk:"lat"`
	Lng             types.Float64 `tfsdk:"lng"`
	MoveMapMarker   types.Bool    `tfsdk:"move_map_marker"`
	Name            types.String  `tfsdk:"name"`
	Notes           types.String  `tfsdk:"notes"`
	SwitchProfileId types.String  `tfsdk:"switch_profile_id"`
	Tags            types.List    `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Device) getPath() string {
	return fmt.Sprintf("/devices/%v", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Device) toBody(ctx context.Context, state Device) string {
	body := ""
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, "address", data.Address.ValueString())
	}
	if !data.FloorPlanId.IsNull() {
		body, _ = sjson.Set(body, "floorPlanId", data.FloorPlanId.ValueString())
	}
	if !data.Lat.IsNull() {
		body, _ = sjson.Set(body, "lat", data.Lat.ValueFloat64())
	}
	if !data.Lng.IsNull() {
		body, _ = sjson.Set(body, "lng", data.Lng.ValueFloat64())
	}
	if !data.MoveMapMarker.IsNull() {
		body, _ = sjson.Set(body, "moveMapMarker", data.MoveMapMarker.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Notes.IsNull() {
		body, _ = sjson.Set(body, "notes", data.Notes.ValueString())
	}
	if !data.SwitchProfileId.IsNull() {
		body, _ = sjson.Set(body, "switchProfileId", data.SwitchProfileId.ValueString())
	}
	if !data.Tags.IsNull() {
		var values []string
		data.Tags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "tags", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Device) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("address"); value.Exists() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("floorPlanId"); value.Exists() {
		data.FloorPlanId = types.StringValue(value.String())
	} else {
		data.FloorPlanId = types.StringNull()
	}
	if value := res.Get("lat"); value.Exists() {
		data.Lat = types.Float64Value(value.Float())
	} else {
		data.Lat = types.Float64Null()
	}
	if value := res.Get("lng"); value.Exists() {
		data.Lng = types.Float64Value(value.Float())
	} else {
		data.Lng = types.Float64Null()
	}
	if value := res.Get("moveMapMarker"); value.Exists() {
		data.MoveMapMarker = types.BoolValue(value.Bool())
	} else {
		data.MoveMapMarker = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("notes"); value.Exists() {
		data.Notes = types.StringValue(value.String())
	} else {
		data.Notes = types.StringNull()
	}
	if value := res.Get("switchProfileId"); value.Exists() {
		data.SwitchProfileId = types.StringValue(value.String())
	} else {
		data.SwitchProfileId = types.StringNull()
	}
	if value := res.Get("tags"); value.Exists() {
		data.Tags = helpers.GetStringList(value.Array())
	} else {
		data.Tags = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *Device) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("floorPlanId"); value.Exists() && !data.FloorPlanId.IsNull() {
		data.FloorPlanId = types.StringValue(value.String())
	} else {
		data.FloorPlanId = types.StringNull()
	}
	if value := res.Get("lat"); value.Exists() && !data.Lat.IsNull() {
		data.Lat = types.Float64Value(value.Float())
	} else {
		data.Lat = types.Float64Null()
	}
	if value := res.Get("lng"); value.Exists() && !data.Lng.IsNull() {
		data.Lng = types.Float64Value(value.Float())
	} else {
		data.Lng = types.Float64Null()
	}
	if value := res.Get("moveMapMarker"); value.Exists() && !data.MoveMapMarker.IsNull() {
		data.MoveMapMarker = types.BoolValue(value.Bool())
	} else {
		data.MoveMapMarker = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("notes"); value.Exists() && !data.Notes.IsNull() {
		data.Notes = types.StringValue(value.String())
	} else {
		data.Notes = types.StringNull()
	}
	if value := res.Get("switchProfileId"); value.Exists() && !data.SwitchProfileId.IsNull() {
		data.SwitchProfileId = types.StringValue(value.String())
	} else {
		data.SwitchProfileId = types.StringNull()
	}
	if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
		data.Tags = helpers.GetStringList(value.Array())
	} else {
		data.Tags = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
