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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationIntegrationsDeployable struct {
	Id                       types.String                              `tfsdk:"id"`
	OrganizationId           types.String                              `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                               `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                               `tfsdk:"meta_counts_items_total"`
	Items                    []OrganizationIntegrationsDeployableItems `tfsdk:"items"`
}

type OrganizationIntegrationsDeployableItems struct {
	IsCiscoProduct   types.Bool   `tfsdk:"is_cisco_product"`
	IsDeployable     types.Bool   `tfsdk:"is_deployable"`
	LogoUrl          types.String `tfsdk:"logo_url"`
	Name             types.String `tfsdk:"name"`
	Provider         types.String `tfsdk:"provider"`
	RedirectUrl      types.String `tfsdk:"redirect_url"`
	ReleaseType      types.String `tfsdk:"release_type"`
	ShortDescription types.String `tfsdk:"short_description"`
	Type             types.String `tfsdk:"type"`
	Tags             types.List   `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationIntegrationsDeployable) getPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/deployable", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationIntegrationsDeployable) toBody(ctx context.Context, state OrganizationIntegrationsDeployable) string {
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
			if !item.IsCiscoProduct.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isCiscoProduct", item.IsCiscoProduct.ValueBool())
			}
			if !item.IsDeployable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isDeployable", item.IsDeployable.ValueBool())
			}
			if !item.LogoUrl.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logoUrl", item.LogoUrl.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Provider.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "provider", item.Provider.ValueString())
			}
			if !item.RedirectUrl.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "redirectUrl", item.RedirectUrl.ValueString())
			}
			if !item.ReleaseType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "releaseType", item.ReleaseType.ValueString())
			}
			if !item.ShortDescription.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "shortDescription", item.ShortDescription.ValueString())
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

func (data *OrganizationIntegrationsDeployable) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Items = make([]OrganizationIntegrationsDeployableItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationIntegrationsDeployableItems{}
			if value := res.Get("isCiscoProduct"); value.Exists() && value.Value() != nil {
				data.IsCiscoProduct = types.BoolValue(value.Bool())
			} else {
				data.IsCiscoProduct = types.BoolNull()
			}
			if value := res.Get("isDeployable"); value.Exists() && value.Value() != nil {
				data.IsDeployable = types.BoolValue(value.Bool())
			} else {
				data.IsDeployable = types.BoolNull()
			}
			if value := res.Get("logoUrl"); value.Exists() && value.Value() != nil {
				data.LogoUrl = types.StringValue(value.String())
			} else {
				data.LogoUrl = types.StringNull()
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
			if value := res.Get("redirectUrl"); value.Exists() && value.Value() != nil {
				data.RedirectUrl = types.StringValue(value.String())
			} else {
				data.RedirectUrl = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortDescription"); value.Exists() && value.Value() != nil {
				data.ShortDescription = types.StringValue(value.String())
			} else {
				data.ShortDescription = types.StringNull()
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
func (data *OrganizationIntegrationsDeployable) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
		keys := [...]string{"isCiscoProduct", "isDeployable", "logoUrl", "name", "provider", "redirectUrl", "releaseType", "shortDescription", "type"}
		keyValues := [...]string{strconv.FormatBool(data.Items[i].IsCiscoProduct.ValueBool()), strconv.FormatBool(data.Items[i].IsDeployable.ValueBool()), data.Items[i].LogoUrl.ValueString(), data.Items[i].Name.ValueString(), data.Items[i].Provider.ValueString(), data.Items[i].RedirectUrl.ValueString(), data.Items[i].ReleaseType.ValueString(), data.Items[i].ShortDescription.ValueString(), data.Items[i].Type.ValueString()}

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
		if value := res.Get("isCiscoProduct"); value.Exists() && !data.IsCiscoProduct.IsNull() {
			data.IsCiscoProduct = types.BoolValue(value.Bool())
		} else {
			data.IsCiscoProduct = types.BoolNull()
		}
		if value := res.Get("isDeployable"); value.Exists() && !data.IsDeployable.IsNull() {
			data.IsDeployable = types.BoolValue(value.Bool())
		} else {
			data.IsDeployable = types.BoolNull()
		}
		if value := res.Get("logoUrl"); value.Exists() && !data.LogoUrl.IsNull() {
			data.LogoUrl = types.StringValue(value.String())
		} else {
			data.LogoUrl = types.StringNull()
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
		if value := res.Get("redirectUrl"); value.Exists() && !data.RedirectUrl.IsNull() {
			data.RedirectUrl = types.StringValue(value.String())
		} else {
			data.RedirectUrl = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortDescription"); value.Exists() && !data.ShortDescription.IsNull() {
			data.ShortDescription = types.StringValue(value.String())
		} else {
			data.ShortDescription = types.StringNull()
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
func (data *OrganizationIntegrationsDeployable) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationIntegrationsDeployable) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
