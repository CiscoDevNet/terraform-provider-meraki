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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkFloorPlans struct {
	NetworkId types.String                       `tfsdk:"network_id"`
	Items     []DataSourceNetworkFloorPlansItems `tfsdk:"items"`
}

type DataSourceNetworkFloorPlansItems struct {
	Id                   types.String  `tfsdk:"id"`
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

func (data DataSourceNetworkFloorPlans) getPath() string {
	return fmt.Sprintf("/networks/%v/floorPlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkFloorPlans) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceNetworkFloorPlansItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceNetworkFloorPlansItems{}
		data.Id = types.StringValue(res.Get("floorPlanId").String())
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
