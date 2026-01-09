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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationAuthRADIUSServer struct {
	Id             types.String                        `tfsdk:"id"`
	OrganizationId types.String                        `tfsdk:"organization_id"`
	Address        types.String                        `tfsdk:"address"`
	Name           types.String                        `tfsdk:"name"`
	Secret         types.String                        `tfsdk:"secret"`
	Modes          []OrganizationAuthRADIUSServerModes `tfsdk:"modes"`
}

type OrganizationAuthRADIUSServerModes struct {
	Mode types.String `tfsdk:"mode"`
	Port types.Int64  `tfsdk:"port"`
}

type OrganizationAuthRADIUSServerIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationAuthRADIUSServer) getPath() string {
	return fmt.Sprintf("/organizations/%v/auth/radius/servers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationAuthRADIUSServer) toBody(ctx context.Context, state OrganizationAuthRADIUSServer) string {
	body := ""
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, "address", data.Address.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Secret.IsNull() {
		body, _ = sjson.Set(body, "secret", data.Secret.ValueString())
	}
	{
		body, _ = sjson.Set(body, "modes", []interface{}{})
		for _, item := range data.Modes {
			itemBody := ""
			if !item.Mode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mode", item.Mode.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "modes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationAuthRADIUSServer) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("address"); value.Exists() && value.Value() != nil {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("modes"); value.Exists() && value.Value() != nil {
		data.Modes = make([]OrganizationAuthRADIUSServerModes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationAuthRADIUSServerModes{}
			if value := res.Get("mode"); value.Exists() && value.Value() != nil {
				data.Mode = types.StringValue(value.String())
			} else {
				data.Mode = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && value.Value() != nil {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).Modes = append((*parent).Modes, data)
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
func (data *OrganizationAuthRADIUSServer) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := 0; i < len(data.Modes); i++ {
		keys := [...]string{"mode", "port"}
		keyValues := [...]string{data.Modes[i].Mode.ValueString(), strconv.FormatInt(data.Modes[i].Port.ValueInt64(), 10)}

		parent := &data
		data := (*parent).Modes[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("modes").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Modes[%d] = %+v",
				i,
				(*parent).Modes[i],
			))
			(*parent).Modes = slices.Delete((*parent).Modes, i, i+1)
			i--

			continue
		}
		if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
			data.Mode = types.StringValue(value.String())
		} else {
			data.Mode = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		(*parent).Modes[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationAuthRADIUSServer) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationAuthRADIUSServerIdentity) toIdentity(ctx context.Context, plan *OrganizationAuthRADIUSServer) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationAuthRADIUSServer) fromIdentity(ctx context.Context, identity *OrganizationAuthRADIUSServerIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationAuthRADIUSServer) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
