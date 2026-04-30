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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkVLANProfileAssignment struct {
	Id               types.String `tfsdk:"id"`
	NetworkId        types.String `tfsdk:"network_id"`
	VlanProfileIname types.String `tfsdk:"vlan_profile_iname"`
	Serials          types.Set    `tfsdk:"serials"`
	StackIds         types.Set    `tfsdk:"stack_ids"`
}

type NetworkVLANProfileAssignmentIdentity struct {
	NetworkId        types.String `tfsdk:"network_id"`
	VlanProfileIname types.String `tfsdk:"vlan_profile_iname"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkVLANProfileAssignment) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles/assignments/reassign", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkVLANProfileAssignment) toBody(ctx context.Context, state NetworkVLANProfileAssignment) string {
	body := ""
	if !data.VlanProfileIname.IsNull() {
		body, _ = sjson.Set(body, "vlanProfile.iname", data.VlanProfileIname.ValueString())
	}
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	} else {
		body, _ = sjson.Set(body, "serials", []interface{}{})
	}
	if !data.StackIds.IsNull() {
		var values []string
		data.StackIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "stackIds", values)
	} else {
		body, _ = sjson.Set(body, "stackIds", []interface{}{})
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkVLANProfileAssignment) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("vlanProfile.iname"); value.Exists() && value.Value() != nil {
		data.VlanProfileIname = types.StringValue(value.String())
	} else {
		data.VlanProfileIname = types.StringNull()
	}
	if value := res.Get("serials"); value.Exists() && value.Value() != nil {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
	if value := res.Get("stackIds"); value.Exists() && value.Value() != nil {
		data.StackIds = helpers.GetStringSet(value.Array())
	} else {
		data.StackIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkVLANProfileAssignment) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("vlanProfile.iname"); value.Exists() && !data.VlanProfileIname.IsNull() {
		data.VlanProfileIname = types.StringValue(value.String())
	} else {
		data.VlanProfileIname = types.StringNull()
	}
	if value := res.Get("serials"); value.Exists() && !data.Serials.IsNull() {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
	if value := res.Get("stackIds"); value.Exists() && !data.StackIds.IsNull() {
		data.StackIds = helpers.GetStringSet(value.Array())
	} else {
		data.StackIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkVLANProfileAssignment) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkVLANProfileAssignmentIdentity) toIdentity(ctx context.Context, plan *NetworkVLANProfileAssignment) {
	data.NetworkId = plan.NetworkId
	data.VlanProfileIname = plan.VlanProfileIname
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkVLANProfileAssignment) fromIdentity(ctx context.Context, identity *NetworkVLANProfileAssignmentIdentity) {
	data.NetworkId = identity.NetworkId
	data.VlanProfileIname = identity.VlanProfileIname
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkVLANProfileAssignment) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

func (data NetworkVLANProfileAssignment) getByDevicePath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles/assignments/byDevice", url.QueryEscape(data.NetworkId.ValueString()))
}

func (data *NetworkVLANProfileAssignment) fromByDeviceBody(ctx context.Context, res meraki.Res) {
	iname := data.VlanProfileIname.ValueString()
	if iname == "" {
		iname = data.Id.ValueString()
	}
	var serialValues []attr.Value
	var stackIdValues []attr.Value
	stackIdSet := map[string]bool{}
	for _, item := range res.Array() {
		profileIname := item.Get("vlanProfile.iname").String()
		if profileIname != iname {
			continue
		}
		if serial := item.Get("serial").String(); serial != "" {
			serialValues = append(serialValues, types.StringValue(serial))
		}
		if stackId := item.Get("stack.id").String(); stackId != "" && !stackIdSet[stackId] {
			stackIdSet[stackId] = true
			stackIdValues = append(stackIdValues, types.StringValue(stackId))
		}
	}
	if len(serialValues) > 0 {
		data.Serials = types.SetValueMust(types.StringType, serialValues)
	} else if !data.Serials.IsNull() {
		data.Serials = types.SetValueMust(types.StringType, []attr.Value{})
	}
	if !data.StackIds.IsNull() || len(stackIdValues) > 0 {
		data.StackIds = types.SetValueMust(types.StringType, stackIdValues)
	}
}

func (data NetworkVLANProfileAssignment) toRemovedBody(ctx context.Context, state NetworkVLANProfileAssignment) string {
	var stateSerials, planSerials []string
	state.Serials.ElementsAs(ctx, &stateSerials, false)
	data.Serials.ElementsAs(ctx, &planSerials, false)

	planSet := map[string]bool{}
	for _, s := range planSerials {
		planSet[s] = true
	}
	var removedSerials []string
	for _, s := range stateSerials {
		if !planSet[s] {
			removedSerials = append(removedSerials, s)
		}
	}

	var stateStacks, planStacks []string
	if !state.StackIds.IsNull() {
		state.StackIds.ElementsAs(ctx, &stateStacks, false)
	}
	if !data.StackIds.IsNull() {
		data.StackIds.ElementsAs(ctx, &planStacks, false)
	}

	planStackSet := map[string]bool{}
	for _, s := range planStacks {
		planStackSet[s] = true
	}
	var removedStacks []string
	for _, s := range stateStacks {
		if !planStackSet[s] {
			removedStacks = append(removedStacks, s)
		}
	}

	if len(removedSerials) == 0 && len(removedStacks) == 0 {
		return ""
	}

	body := ""
	body, _ = sjson.Set(body, "vlanProfile.iname", "Default")
	body, _ = sjson.Set(body, "serials", removedSerials)
	body, _ = sjson.Set(body, "stackIds", removedStacks)
	return body
}

func (data NetworkVLANProfileAssignment) toDefaultBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "vlanProfile.iname", "Default")
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	} else {
		body, _ = sjson.Set(body, "serials", []string{})
	}
	if !data.StackIds.IsNull() {
		var values []string
		data.StackIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "stackIds", values)
	} else {
		body, _ = sjson.Set(body, "stackIds", []string{})
	}
	return body
}
