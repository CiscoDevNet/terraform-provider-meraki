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

type DataSourceAppliancePorts struct {
	NetworkId types.String                    `tfsdk:"network_id"`
	Items     []DataSourceAppliancePortsItems `tfsdk:"items"`
}

type DataSourceAppliancePortsItems struct {
	Id                  types.String `tfsdk:"id"`
	AccessPolicy        types.String `tfsdk:"access_policy"`
	AllowedVlans        types.String `tfsdk:"allowed_vlans"`
	DropUntaggedTraffic types.Bool   `tfsdk:"drop_untagged_traffic"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Number              types.Int64  `tfsdk:"number"`
	Type                types.String `tfsdk:"type"`
	Vlan                types.Int64  `tfsdk:"vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceAppliancePorts) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/ports", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceAppliancePorts) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceAppliancePortsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceAppliancePortsItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("accessPolicy"); value.Exists() && value.Value() != nil {
			data.AccessPolicy = types.StringValue(value.String())
		} else {
			data.AccessPolicy = types.StringNull()
		}
		if value := res.Get("allowedVlans"); value.Exists() && value.Value() != nil {
			data.AllowedVlans = types.StringValue(value.String())
		} else {
			data.AllowedVlans = types.StringNull()
		}
		if value := res.Get("dropUntaggedTraffic"); value.Exists() && value.Value() != nil {
			data.DropUntaggedTraffic = types.BoolValue(value.Bool())
		} else {
			data.DropUntaggedTraffic = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("number"); value.Exists() && value.Value() != nil {
			data.Number = types.Int64Value(value.Int())
		} else {
			data.Number = types.Int64Null()
		}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
			data.Vlan = types.Int64Value(value.Int())
		} else {
			data.Vlan = types.Int64Null()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
