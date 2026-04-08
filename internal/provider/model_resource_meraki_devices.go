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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceDevices struct {
	Id             types.String           `tfsdk:"id"`
	OrganizationId types.String           `tfsdk:"organization_id"`
	Items          []ResourceDevicesItems `tfsdk:"items"`
}

type ResourceDevicesItems struct {
	Serial          types.String  `tfsdk:"serial"`
	Address         types.String  `tfsdk:"address"`
	FloorPlanId     types.String  `tfsdk:"floor_plan_id"`
	Lat             types.Float64 `tfsdk:"lat"`
	Lng             types.Float64 `tfsdk:"lng"`
	MoveMapMarker   types.Bool    `tfsdk:"move_map_marker"`
	Name            types.String  `tfsdk:"name"`
	Notes           types.String  `tfsdk:"notes"`
	SwitchProfileId types.String  `tfsdk:"switch_profile_id"`
	Tags            types.Set     `tfsdk:"tags"`
}

// End of section. //template:end types

func (data ResourceDevices) getPath() string {
	return fmt.Sprintf("/organizations/%v/devices", data.OrganizationId.ValueString())
}

func (data ResourceDevices) getItemPath(id string) string {
	return fmt.Sprintf("/devices/%v", id)
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceDevicesItems) toBody(ctx context.Context, state ResourceDevicesItems) string {
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

func (data *ResourceDevices) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceDevicesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceDevicesItems{}
		if value := res.Get("address"); value.Exists() && value.Value() != nil {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("lat"); value.Exists() && value.Value() != nil {
			data.Lat = types.Float64Value(value.Float())
		} else {
			data.Lat = types.Float64Null()
		}
		if value := res.Get("lng"); value.Exists() && value.Value() != nil {
			data.Lng = types.Float64Value(value.Float())
		} else {
			data.Lng = types.Float64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("notes"); value.Exists() && value.Value() != nil {
			data.Notes = types.StringValue(value.String())
		} else {
			data.Notes = types.StringNull()
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceDevices) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("serial").String() == (*parent).Items[i].Serial.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
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
		if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Serial.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceDevices) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceDevices) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("serial").String() == (*parent).Items[i].Serial.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("address"); value.Exists() && value.Value() != nil {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("lat"); value.Exists() && value.Value() != nil {
			data.Lat = types.Float64Value(value.Float())
		} else {
			data.Lat = types.Float64Null()
		}
		if value := res.Get("lng"); value.Exists() && value.Value() != nil {
			data.Lng = types.Float64Value(value.Float())
		} else {
			data.Lng = types.Float64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("notes"); value.Exists() && value.Value() != nil {
			data.Notes = types.StringValue(value.String())
		} else {
			data.Notes = types.StringNull()
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Serial.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceDevices) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceDevices) hasChanges(ctx context.Context, state *ResourceDevices, id string) bool {
	hasChanges := false

	item := ResourceDevicesItems{}
	for i := range data.Items {
		if data.Items[i].Serial.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceDevicesItems{}
	for i := range state.Items {
		if state.Items[i].Serial.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Address.Equal(stateItem.Address) {
		hasChanges = true
	}
	if !item.FloorPlanId.Equal(stateItem.FloorPlanId) {
		hasChanges = true
	}
	if !item.Lat.Equal(stateItem.Lat) {
		hasChanges = true
	}
	if !item.Lng.Equal(stateItem.Lng) {
		hasChanges = true
	}
	if !item.MoveMapMarker.Equal(stateItem.MoveMapMarker) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.Notes.Equal(stateItem.Notes) {
		hasChanges = true
	}
	if !item.SwitchProfileId.Equal(stateItem.SwitchProfileId) {
		hasChanges = true
	}
	if !item.Tags.Equal(stateItem.Tags) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
