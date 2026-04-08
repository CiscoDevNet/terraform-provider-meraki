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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceDNSLocalProfileAssignments struct {
	Id             types.String                               `tfsdk:"id"`
	OrganizationId types.String                               `tfsdk:"organization_id"`
	Items          []ApplianceDNSLocalProfileAssignmentsItems `tfsdk:"items"`
}

type ApplianceDNSLocalProfileAssignmentsItems struct {
	AssignmentId types.String `tfsdk:"assignment_id"`
	NetworkId    types.String `tfsdk:"network_id"`
	ProfileId    types.String `tfsdk:"profile_id"`
}

type ApplianceDNSLocalProfileAssignmentsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceDNSLocalProfileAssignments) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/local/profiles/assignments/bulkCreate", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data ApplianceDNSLocalProfileAssignments) getDeletePath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/local/profiles/assignments/bulkDelete", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data ApplianceDNSLocalProfileAssignments) getAssignmentsPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/local/profiles/assignments", url.QueryEscape(data.OrganizationId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceDNSLocalProfileAssignments) toBody(ctx context.Context, state ApplianceDNSLocalProfileAssignments) string {
	body := ""
	{
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "network.id", item.NetworkId.ValueString())
			}
			if !item.ProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "profile.id", item.ProfileId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

func (data ApplianceDNSLocalProfileAssignments) toBodyDelete(ctx context.Context) string {
	body := ""
	{
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.AssignmentId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "assignmentId", item.AssignmentId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceDNSLocalProfileAssignments) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]ApplianceDNSLocalProfileAssignmentsItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceDNSLocalProfileAssignmentsItems{}
			if value := res.Get("assignmentId"); value.Exists() && value.Value() != nil {
				data.AssignmentId = types.StringValue(value.String())
			} else {
				data.AssignmentId = types.StringNull()
			}
			if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("profile.id"); value.Exists() && value.Value() != nil {
				data.ProfileId = types.StringValue(value.String())
			} else {
				data.ProfileId = types.StringNull()
			}
			(*parent).Items = append((*parent).Items, data)
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
func (data *ApplianceDNSLocalProfileAssignments) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"assignmentId", "network.id", "profile.id"}
		keyValues := [...]string{data.Items[i].AssignmentId.ValueString(), data.Items[i].NetworkId.ValueString(), data.Items[i].ProfileId.ValueString()}

		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Items[%d] = %+v",
				i,
				(*parent).Items[i],
			))
			(*parent).Items = slices.Delete((*parent).Items, i, i+1)
			i--

			continue
		}
		if value := res.Get("assignmentId"); value.Exists() && !data.AssignmentId.IsNull() {
			data.AssignmentId = types.StringValue(value.String())
		} else {
			data.AssignmentId = types.StringNull()
		}
		if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("profile.id"); value.Exists() && !data.ProfileId.IsNull() {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceDNSLocalProfileAssignments) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"assignmentId", "network.id", "profile.id"}
		keyValues := [...]string{data.Items[i].AssignmentId.ValueString(), data.Items[i].NetworkId.ValueString(), data.Items[i].ProfileId.ValueString()}

		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Items[%d] = %+v",
				i,
				(*parent).Items[i],
			))
			(*parent).Items = slices.Delete((*parent).Items, i, i+1)
			i--

			continue
		}
		if data.AssignmentId.IsUnknown() {
			if value := res.Get("assignmentId"); value.Exists() && !data.AssignmentId.IsNull() {
				data.AssignmentId = types.StringValue(value.String())
			} else {
				data.AssignmentId = types.StringNull()
			}
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceDNSLocalProfileAssignmentsIdentity) toIdentity(ctx context.Context, plan *ApplianceDNSLocalProfileAssignments) {
	data.OrganizationId = plan.OrganizationId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceDNSLocalProfileAssignments) fromIdentity(ctx context.Context, identity *ApplianceDNSLocalProfileAssignmentsIdentity) {
	data.OrganizationId = identity.OrganizationId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceDNSLocalProfileAssignments) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
