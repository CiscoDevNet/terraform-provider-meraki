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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationAdaptivePolicy struct {
	Id                   types.String                     `tfsdk:"id"`
	OrganizationId       types.String                     `tfsdk:"organization_id"`
	LastEntryRule        types.String                     `tfsdk:"last_entry_rule"`
	DestinationGroupId   types.String                     `tfsdk:"destination_group_id"`
	DestinationGroupName types.String                     `tfsdk:"destination_group_name"`
	DestinationGroupSgt  types.Int64                      `tfsdk:"destination_group_sgt"`
	SourceGroupId        types.String                     `tfsdk:"source_group_id"`
	SourceGroupName      types.String                     `tfsdk:"source_group_name"`
	SourceGroupSgt       types.Int64                      `tfsdk:"source_group_sgt"`
	Acls                 []OrganizationAdaptivePolicyAcls `tfsdk:"acls"`
}

type OrganizationAdaptivePolicyAcls struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationAdaptivePolicy) getPath() string {
	return fmt.Sprintf("/organizations/%v/adaptivePolicy/policies", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationAdaptivePolicy) toBody(ctx context.Context, state OrganizationAdaptivePolicy) string {
	body := ""
	if !data.LastEntryRule.IsNull() {
		body, _ = sjson.Set(body, "lastEntryRule", data.LastEntryRule.ValueString())
	}
	if !data.DestinationGroupId.IsNull() {
		body, _ = sjson.Set(body, "destinationGroup.id", data.DestinationGroupId.ValueString())
	}
	if !data.DestinationGroupName.IsNull() {
		body, _ = sjson.Set(body, "destinationGroup.name", data.DestinationGroupName.ValueString())
	}
	if !data.DestinationGroupSgt.IsNull() {
		body, _ = sjson.Set(body, "destinationGroup.sgt", data.DestinationGroupSgt.ValueInt64())
	}
	if !data.SourceGroupId.IsNull() {
		body, _ = sjson.Set(body, "sourceGroup.id", data.SourceGroupId.ValueString())
	}
	if !data.SourceGroupName.IsNull() {
		body, _ = sjson.Set(body, "sourceGroup.name", data.SourceGroupName.ValueString())
	}
	if !data.SourceGroupSgt.IsNull() {
		body, _ = sjson.Set(body, "sourceGroup.sgt", data.SourceGroupSgt.ValueInt64())
	}
	if len(data.Acls) > 0 {
		body, _ = sjson.Set(body, "acls", []interface{}{})
		for _, item := range data.Acls {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "acls.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationAdaptivePolicy) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("lastEntryRule"); value.Exists() && value.Value() != nil {
		data.LastEntryRule = types.StringValue(value.String())
	} else {
		data.LastEntryRule = types.StringNull()
	}
	if value := res.Get("destinationGroup.id"); value.Exists() && value.Value() != nil {
		data.DestinationGroupId = types.StringValue(value.String())
	} else {
		data.DestinationGroupId = types.StringNull()
	}
	if value := res.Get("destinationGroup.name"); value.Exists() && value.Value() != nil {
		data.DestinationGroupName = types.StringValue(value.String())
	} else {
		data.DestinationGroupName = types.StringNull()
	}
	if value := res.Get("destinationGroup.sgt"); value.Exists() && value.Value() != nil {
		data.DestinationGroupSgt = types.Int64Value(value.Int())
	} else {
		data.DestinationGroupSgt = types.Int64Null()
	}
	if value := res.Get("sourceGroup.id"); value.Exists() && value.Value() != nil {
		data.SourceGroupId = types.StringValue(value.String())
	} else {
		data.SourceGroupId = types.StringNull()
	}
	if value := res.Get("sourceGroup.name"); value.Exists() && value.Value() != nil {
		data.SourceGroupName = types.StringValue(value.String())
	} else {
		data.SourceGroupName = types.StringNull()
	}
	if value := res.Get("sourceGroup.sgt"); value.Exists() && value.Value() != nil {
		data.SourceGroupSgt = types.Int64Value(value.Int())
	} else {
		data.SourceGroupSgt = types.Int64Null()
	}
	if value := res.Get("acls"); value.Exists() && value.Value() != nil {
		data.Acls = make([]OrganizationAdaptivePolicyAcls, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationAdaptivePolicyAcls{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Acls = append((*parent).Acls, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationAdaptivePolicy) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("lastEntryRule"); value.Exists() && !data.LastEntryRule.IsNull() {
		data.LastEntryRule = types.StringValue(value.String())
	} else {
		data.LastEntryRule = types.StringNull()
	}
	if value := res.Get("destinationGroup.id"); value.Exists() && !data.DestinationGroupId.IsNull() {
		data.DestinationGroupId = types.StringValue(value.String())
	} else {
		data.DestinationGroupId = types.StringNull()
	}
	if value := res.Get("destinationGroup.name"); value.Exists() && !data.DestinationGroupName.IsNull() {
		data.DestinationGroupName = types.StringValue(value.String())
	} else {
		data.DestinationGroupName = types.StringNull()
	}
	if value := res.Get("destinationGroup.sgt"); value.Exists() && !data.DestinationGroupSgt.IsNull() {
		data.DestinationGroupSgt = types.Int64Value(value.Int())
	} else {
		data.DestinationGroupSgt = types.Int64Null()
	}
	if value := res.Get("sourceGroup.id"); value.Exists() && !data.SourceGroupId.IsNull() {
		data.SourceGroupId = types.StringValue(value.String())
	} else {
		data.SourceGroupId = types.StringNull()
	}
	if value := res.Get("sourceGroup.name"); value.Exists() && !data.SourceGroupName.IsNull() {
		data.SourceGroupName = types.StringValue(value.String())
	} else {
		data.SourceGroupName = types.StringNull()
	}
	if value := res.Get("sourceGroup.sgt"); value.Exists() && !data.SourceGroupSgt.IsNull() {
		data.SourceGroupSgt = types.Int64Value(value.Int())
	} else {
		data.SourceGroupSgt = types.Int64Null()
	}
	{
		l := len(res.Get("acls").Array())
		tflog.Debug(ctx, fmt.Sprintf("acls array resizing from %d to %d", len(data.Acls), l))
		if len(data.Acls) > l {
			data.Acls = data.Acls[:l]
		}
	}
	for i := range data.Acls {
		parent := &data
		data := (*parent).Acls[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("acls.%d", i))
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Acls[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationAdaptivePolicy) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationAdaptivePolicy) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
