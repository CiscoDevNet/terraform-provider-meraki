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

type ApplianceUplinksNAT struct {
	Id        types.String                 `tfsdk:"id"`
	NetworkId types.String                 `tfsdk:"network_id"`
	Uplinks   []ApplianceUplinksNATUplinks `tfsdk:"uplinks"`
}

type ApplianceUplinksNATUplinks struct {
	Interface  types.String `tfsdk:"interface"`
	NatEnabled types.Bool   `tfsdk:"nat_enabled"`
}

type ApplianceUplinksNATIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceUplinksNAT) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/uplinks/nat", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceUplinksNAT) toBody(ctx context.Context, state ApplianceUplinksNAT) string {
	body := ""
	{
		body, _ = sjson.Set(body, "uplinks", []interface{}{})
		for _, item := range data.Uplinks {
			itemBody := ""
			if !item.Interface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface", item.Interface.ValueString())
			}
			if !item.NatEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nat.enabled", item.NatEnabled.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "uplinks.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data ApplianceUplinksNAT) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("uplinks"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "uplinks", "null")
		} else {
			body, _ = sjson.Set(body, "uplinks", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("interface"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "interface", "null")
					} else {
						body, _ = sjson.Set(body, "interface", value.String())
					}
				}
				if value := res.Get("nat.enabled"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "nat.enabled", "null")
					} else {
						body, _ = sjson.Set(body, "nat.enabled", value.Bool())
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "uplinks.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceUplinksNAT) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("uplinks"); value.Exists() && value.Value() != nil {
		data.Uplinks = make([]ApplianceUplinksNATUplinks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceUplinksNATUplinks{}
			if value := res.Get("interface"); value.Exists() && value.Value() != nil {
				data.Interface = types.StringValue(value.String())
			} else {
				data.Interface = types.StringNull()
			}
			if value := res.Get("nat.enabled"); value.Exists() && value.Value() != nil {
				data.NatEnabled = types.BoolValue(value.Bool())
			} else {
				data.NatEnabled = types.BoolNull()
			}
			(*parent).Uplinks = append((*parent).Uplinks, data)
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
func (data *ApplianceUplinksNAT) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Uplinks); i++ {
		keys := [...]string{"interface", "nat.enabled"}
		keyValues := [...]string{data.Uplinks[i].Interface.ValueString(), strconv.FormatBool(data.Uplinks[i].NatEnabled.ValueBool())}

		parent := &data
		data := (*parent).Uplinks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("uplinks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Uplinks[%d] = %+v",
				i,
				(*parent).Uplinks[i],
			))
			(*parent).Uplinks = slices.Delete((*parent).Uplinks, i, i+1)
			i--

			continue
		}
		if value := res.Get("interface"); value.Exists() && !data.Interface.IsNull() {
			data.Interface = types.StringValue(value.String())
		} else {
			data.Interface = types.StringNull()
		}
		if value := res.Get("nat.enabled"); value.Exists() && !data.NatEnabled.IsNull() {
			data.NatEnabled = types.BoolValue(value.Bool())
		} else {
			data.NatEnabled = types.BoolNull()
		}
		(*parent).Uplinks[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceUplinksNAT) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceUplinksNATIdentity) toIdentity(ctx context.Context, plan *ApplianceUplinksNAT) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceUplinksNAT) fromIdentity(ctx context.Context, identity *ApplianceUplinksNATIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceUplinksNAT) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
