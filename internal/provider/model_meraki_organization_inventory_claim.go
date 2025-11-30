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

type OrganizationInventoryClaim struct {
	Id             types.String                         `tfsdk:"id"`
	OrganizationId types.String                         `tfsdk:"organization_id"`
	Licenses       []OrganizationInventoryClaimLicenses `tfsdk:"licenses"`
	Orders         types.Set                            `tfsdk:"orders"`
	Serials        types.Set                            `tfsdk:"serials"`
}

type OrganizationInventoryClaimLicenses struct {
	Key  types.String `tfsdk:"key"`
	Mode types.String `tfsdk:"mode"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationInventoryClaim) getPath() string {
	return fmt.Sprintf("/organizations/%v/inventory/claim", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data OrganizationInventoryClaim) getReleasePath() string {
	return fmt.Sprintf("/organizations/%v/inventory/release", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data OrganizationInventoryClaim) getDevicesPath() string {
	return fmt.Sprintf("/organizations/%v/inventory/devices", url.QueryEscape(data.OrganizationId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationInventoryClaim) toBody(ctx context.Context, state OrganizationInventoryClaim) string {
	body := ""
	if data.Licenses != nil {
		body, _ = sjson.Set(body, "licenses", []interface{}{})
		for _, item := range data.Licenses {
			itemBody := ""
			if !item.Key.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "key", item.Key.ValueString())
			}
			if !item.Mode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mode", item.Mode.ValueString())
			}
			body, _ = sjson.SetRaw(body, "licenses.-1", itemBody)
		}
	}
	if !data.Orders.IsNull() {
		var values []string
		data.Orders.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "orders", values)
	}
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationInventoryClaim) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("licenses"); value.Exists() && value.Value() != nil {
		data.Licenses = make([]OrganizationInventoryClaimLicenses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationInventoryClaimLicenses{}
			if value := res.Get("key"); value.Exists() && value.Value() != nil {
				data.Key = types.StringValue(value.String())
			} else {
				data.Key = types.StringNull()
			}
			if value := res.Get("mode"); value.Exists() && value.Value() != nil {
				data.Mode = types.StringValue(value.String())
			} else {
				data.Mode = types.StringNull()
			}
			(*parent).Licenses = append((*parent).Licenses, data)
			return true
		})
	}
	if value := res.Get("orders"); value.Exists() && value.Value() != nil {
		data.Orders = helpers.GetStringSet(value.Array())
	} else {
		data.Orders = types.SetNull(types.StringType)
	}
	if value := res.Get("serials"); value.Exists() && value.Value() != nil {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationInventoryClaim) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Licenses); i++ {
		keys := [...]string{"key"}
		keyValues := [...]string{data.Licenses[i].Key.ValueString()}

		parent := &data
		data := (*parent).Licenses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("licenses").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Licenses[%d] = %+v",
				i,
				(*parent).Licenses[i],
			))
			(*parent).Licenses = slices.Delete((*parent).Licenses, i, i+1)
			i--

			continue
		}
		if value := res.Get("key"); value.Exists() && !data.Key.IsNull() {
			data.Key = types.StringValue(value.String())
		} else {
			data.Key = types.StringNull()
		}
		if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
			data.Mode = types.StringValue(value.String())
		} else {
			data.Mode = types.StringNull()
		}
		(*parent).Licenses[i] = data
	}
	if value := res.Get("orders"); value.Exists() && !data.Orders.IsNull() {
		data.Orders = helpers.GetStringSet(value.Array())
	} else {
		data.Orders = types.SetNull(types.StringType)
	}
	if value := res.Get("serials"); value.Exists() && !data.Serials.IsNull() {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationInventoryClaim) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationInventoryClaim) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
