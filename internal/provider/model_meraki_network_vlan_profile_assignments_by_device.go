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

type NetworkVLANProfileAssignmentsByDevice struct {
	NetworkId types.String                                 `tfsdk:"network_id"`
	Items     []NetworkVLANProfileAssignmentsByDeviceItems `tfsdk:"items"`
}

type NetworkVLANProfileAssignmentsByDeviceItems struct {
	Id                   types.String `tfsdk:"id"`
	Mac                  types.String `tfsdk:"mac"`
	Name                 types.String `tfsdk:"name"`
	ProductType          types.String `tfsdk:"product_type"`
	Serial               types.String `tfsdk:"serial"`
	StackId              types.String `tfsdk:"stack_id"`
	VlanProfileIname     types.String `tfsdk:"vlan_profile_iname"`
	VlanProfileIsDefault types.Bool   `tfsdk:"vlan_profile_is_default"`
	VlanProfileName      types.String `tfsdk:"vlan_profile_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkVLANProfileAssignmentsByDevice) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles/assignments/byDevice", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkVLANProfileAssignmentsByDevice) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]NetworkVLANProfileAssignmentsByDeviceItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := NetworkVLANProfileAssignmentsByDeviceItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("mac"); value.Exists() && value.Value() != nil {
			data.Mac = types.StringValue(value.String())
		} else {
			data.Mac = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("productType"); value.Exists() && value.Value() != nil {
			data.ProductType = types.StringValue(value.String())
		} else {
			data.ProductType = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && value.Value() != nil {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		if value := res.Get("stack.id"); value.Exists() && value.Value() != nil {
			data.StackId = types.StringValue(value.String())
		} else {
			data.StackId = types.StringNull()
		}
		if value := res.Get("vlanProfile.iname"); value.Exists() && value.Value() != nil {
			data.VlanProfileIname = types.StringValue(value.String())
		} else {
			data.VlanProfileIname = types.StringNull()
		}
		if value := res.Get("vlanProfile.isDefault"); value.Exists() && value.Value() != nil {
			data.VlanProfileIsDefault = types.BoolValue(value.Bool())
		} else {
			data.VlanProfileIsDefault = types.BoolNull()
		}
		if value := res.Get("vlanProfile.name"); value.Exists() && value.Value() != nil {
			data.VlanProfileName = types.StringValue(value.String())
		} else {
			data.VlanProfileName = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
