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

type NetworkFirmwareUpgradesRollback struct {
	Id          types.String                             `tfsdk:"id"`
	NetworkId   types.String                             `tfsdk:"network_id"`
	Product     types.String                             `tfsdk:"product"`
	Time        types.String                             `tfsdk:"time"`
	ToVersionId types.String                             `tfsdk:"to_version_id"`
	Reasons     []NetworkFirmwareUpgradesRollbackReasons `tfsdk:"reasons"`
}

type NetworkFirmwareUpgradesRollbackReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type NetworkFirmwareUpgradesRollbackIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgradesRollback) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades/rollbacks", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgradesRollback) toBody(ctx context.Context, state NetworkFirmwareUpgradesRollback) string {
	body := ""
	if !data.Product.IsNull() {
		body, _ = sjson.Set(body, "product", data.Product.ValueString())
	}
	if !data.Time.IsNull() {
		body, _ = sjson.Set(body, "time", data.Time.ValueString())
	}
	if !data.ToVersionId.IsNull() {
		body, _ = sjson.Set(body, "toVersion.id", data.ToVersionId.ValueString())
	}
	{
		body, _ = sjson.Set(body, "reasons", []interface{}{})
		for _, item := range data.Reasons {
			itemBody := ""
			if !item.Category.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "category", item.Category.ValueString())
			}
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			body, _ = sjson.SetRaw(body, "reasons.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFirmwareUpgradesRollback) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("product"); value.Exists() && value.Value() != nil {
		data.Product = types.StringValue(value.String())
	} else {
		data.Product = types.StringNull()
	}
	if value := res.Get("time"); value.Exists() && value.Value() != nil {
		data.Time = types.StringValue(value.String())
	} else {
		data.Time = types.StringNull()
	}
	if value := res.Get("toVersion.id"); value.Exists() && value.Value() != nil {
		data.ToVersionId = types.StringValue(value.String())
	} else {
		data.ToVersionId = types.StringNull()
	}
	if value := res.Get("reasons"); value.Exists() && value.Value() != nil {
		data.Reasons = make([]NetworkFirmwareUpgradesRollbackReasons, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesRollbackReasons{}
			if value := res.Get("category"); value.Exists() && value.Value() != nil {
				data.Category = types.StringValue(value.String())
			} else {
				data.Category = types.StringNull()
			}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			(*parent).Reasons = append((*parent).Reasons, data)
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
func (data *NetworkFirmwareUpgradesRollback) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("product"); value.Exists() && !data.Product.IsNull() {
		data.Product = types.StringValue(value.String())
	} else {
		data.Product = types.StringNull()
	}
	if value := res.Get("time"); value.Exists() && !data.Time.IsNull() {
		data.Time = types.StringValue(value.String())
	} else {
		data.Time = types.StringNull()
	}
	if value := res.Get("toVersion.id"); value.Exists() && !data.ToVersionId.IsNull() {
		data.ToVersionId = types.StringValue(value.String())
	} else {
		data.ToVersionId = types.StringNull()
	}
	for i := 0; i < len(data.Reasons); i++ {
		keys := [...]string{"category", "comment"}
		keyValues := [...]string{data.Reasons[i].Category.ValueString(), data.Reasons[i].Comment.ValueString()}

		parent := &data
		data := (*parent).Reasons[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("reasons").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Reasons[%d] = %+v",
				i,
				(*parent).Reasons[i],
			))
			(*parent).Reasons = slices.Delete((*parent).Reasons, i, i+1)
			i--

			continue
		}
		if value := res.Get("category"); value.Exists() && !data.Category.IsNull() {
			data.Category = types.StringValue(value.String())
		} else {
			data.Category = types.StringNull()
		}
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		(*parent).Reasons[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFirmwareUpgradesRollback) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesRollbackIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgradesRollback) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgradesRollback) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesRollbackIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgradesRollback) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
