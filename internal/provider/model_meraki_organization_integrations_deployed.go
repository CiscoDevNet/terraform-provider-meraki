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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationIntegrationsDeployed struct {
	Id                       types.String                            `tfsdk:"id"`
	OrganizationId           types.String                            `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                             `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                             `tfsdk:"meta_counts_items_total"`
	Items                    []OrganizationIntegrationsDeployedItems `tfsdk:"items"`
}

type OrganizationIntegrationsDeployedItems struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Provider types.String `tfsdk:"provider"`
	Type     types.String `tfsdk:"type"`
	Tags     types.List   `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationIntegrationsDeployed) getPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/deployed", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationIntegrationsDeployed) toBody(ctx context.Context, state OrganizationIntegrationsDeployed) string {
	body := ""
	if !data.MetaCountsItemsRemaining.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.remaining", data.MetaCountsItemsRemaining.ValueInt64())
	}
	if !data.MetaCountsItemsTotal.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.total", data.MetaCountsItemsTotal.ValueInt64())
	}
	if data.Items != nil {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Provider.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "provider", item.Provider.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Tags.IsNull() {
				var values []string
				item.Tags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "tags", values)
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationIntegrationsDeployed) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]OrganizationIntegrationsDeployedItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationIntegrationsDeployedItems{}
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
			if value := res.Get("provider"); value.Exists() && value.Value() != nil {
				data.Provider = types.StringValue(value.String())
			} else {
				data.Provider = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("tags"); value.Exists() && value.Value() != nil {
				data.Tags = helpers.GetStringList(value.Array())
			} else {
				data.Tags = types.ListNull(types.StringType)
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
func (data *OrganizationIntegrationsDeployed) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && !data.MetaCountsItemsRemaining.IsNull() {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && !data.MetaCountsItemsTotal.IsNull() {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"id", "name", "provider", "type"}
		keyValues := [...]string{data.Items[i].Id.ValueString(), data.Items[i].Name.ValueString(), data.Items[i].Provider.ValueString(), data.Items[i].Type.ValueString()}

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
		if value := res.Get("provider"); value.Exists() && !data.Provider.IsNull() {
			data.Provider = types.StringValue(value.String())
		} else {
			data.Provider = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationIntegrationsDeployed) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationIntegrationsDeployed) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
