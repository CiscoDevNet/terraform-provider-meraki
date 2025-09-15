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

type NetworkFloorPlan struct {
	Id                   types.String  `tfsdk:"id"`
	NetworkId            types.String  `tfsdk:"network_id"`
	FloorNumber          types.Float64 `tfsdk:"floor_number"`
	ImageContents        types.String  `tfsdk:"image_contents"`
	Name                 types.String  `tfsdk:"name"`
	BottomLeftCornerLat  types.Float64 `tfsdk:"bottom_left_corner_lat"`
	BottomLeftCornerLng  types.Float64 `tfsdk:"bottom_left_corner_lng"`
	BottomRightCornerLat types.Float64 `tfsdk:"bottom_right_corner_lat"`
	BottomRightCornerLng types.Float64 `tfsdk:"bottom_right_corner_lng"`
	CenterLat            types.Float64 `tfsdk:"center_lat"`
	CenterLng            types.Float64 `tfsdk:"center_lng"`
	TopLeftCornerLat     types.Float64 `tfsdk:"top_left_corner_lat"`
	TopLeftCornerLng     types.Float64 `tfsdk:"top_left_corner_lng"`
	TopRightCornerLat    types.Float64 `tfsdk:"top_right_corner_lat"`
	TopRightCornerLng    types.Float64 `tfsdk:"top_right_corner_lng"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFloorPlan) getPath() string {
	return fmt.Sprintf("/networks/%v/floorPlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFloorPlan) toBody(ctx context.Context, state NetworkFloorPlan) string {
	body := ""
	if !data.FloorNumber.IsNull() {
		body, _ = sjson.Set(body, "floorNumber", data.FloorNumber.ValueFloat64())
	}
	if !data.ImageContents.IsNull() {
		body, _ = sjson.Set(body, "imageContents", data.ImageContents.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.BottomLeftCornerLat.IsNull() {
		body, _ = sjson.Set(body, "bottomLeftCorner.lat", data.BottomLeftCornerLat.ValueFloat64())
	}
	if !data.BottomLeftCornerLng.IsNull() {
		body, _ = sjson.Set(body, "bottomLeftCorner.lng", data.BottomLeftCornerLng.ValueFloat64())
	}
	if !data.BottomRightCornerLat.IsNull() {
		body, _ = sjson.Set(body, "bottomRightCorner.lat", data.BottomRightCornerLat.ValueFloat64())
	}
	if !data.BottomRightCornerLng.IsNull() {
		body, _ = sjson.Set(body, "bottomRightCorner.lng", data.BottomRightCornerLng.ValueFloat64())
	}
	if !data.CenterLat.IsNull() {
		body, _ = sjson.Set(body, "center.lat", data.CenterLat.ValueFloat64())
	}
	if !data.CenterLng.IsNull() {
		body, _ = sjson.Set(body, "center.lng", data.CenterLng.ValueFloat64())
	}
	if !data.TopLeftCornerLat.IsNull() {
		body, _ = sjson.Set(body, "topLeftCorner.lat", data.TopLeftCornerLat.ValueFloat64())
	}
	if !data.TopLeftCornerLng.IsNull() {
		body, _ = sjson.Set(body, "topLeftCorner.lng", data.TopLeftCornerLng.ValueFloat64())
	}
	if !data.TopRightCornerLat.IsNull() {
		body, _ = sjson.Set(body, "topRightCorner.lat", data.TopRightCornerLat.ValueFloat64())
	}
	if !data.TopRightCornerLng.IsNull() {
		body, _ = sjson.Set(body, "topRightCorner.lng", data.TopRightCornerLng.ValueFloat64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFloorPlan) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("floorNumber"); value.Exists() && value.Value() != nil {
		data.FloorNumber = types.Float64Value(value.Float())
	} else {
		data.FloorNumber = types.Float64Null()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkFloorPlan) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("floorNumber"); value.Exists() && !data.FloorNumber.IsNull() {
		data.FloorNumber = types.Float64Value(value.Float())
	} else {
		data.FloorNumber = types.Float64Null()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFloorPlan) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data NetworkFloorPlan) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
