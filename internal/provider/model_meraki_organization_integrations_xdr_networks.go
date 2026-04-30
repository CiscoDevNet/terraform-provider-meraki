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

type OrganizationIntegrationsXDRNetworks struct {
	Id             types.String                                  `tfsdk:"id"`
	OrganizationId types.String                                  `tfsdk:"organization_id"`
	Networks       []OrganizationIntegrationsXDRNetworksNetworks `tfsdk:"networks"`
}

type OrganizationIntegrationsXDRNetworksNetworks struct {
	NetworkId    types.String `tfsdk:"network_id"`
	ProductTypes types.List   `tfsdk:"product_types"`
}

type OrganizationIntegrationsXDRNetworksIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationIntegrationsXDRNetworks) getPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/xdr/networks/enable", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data OrganizationIntegrationsXDRNetworks) getDisablePath() string {
	return fmt.Sprintf("/organizations/%v/integrations/xdr/networks/disable", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data OrganizationIntegrationsXDRNetworks) getNetworksPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/xdr/networks", url.QueryEscape(data.OrganizationId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationIntegrationsXDRNetworks) toBody(ctx context.Context, state OrganizationIntegrationsXDRNetworks) string {
	body := ""
	{
		body, _ = sjson.Set(body, "networks", []interface{}{})
		for _, item := range data.Networks {
			itemBody := ""
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkId", item.NetworkId.ValueString())
			}
			if !item.ProductTypes.IsNull() {
				var values []string
				item.ProductTypes.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "productTypes", values)
			}
			body, _ = sjson.SetRaw(body, "networks.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationIntegrationsXDRNetworks) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("networks"); value.Exists() && value.Value() != nil {
		data.Networks = make([]OrganizationIntegrationsXDRNetworksNetworks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationIntegrationsXDRNetworksNetworks{}
			if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("productTypes"); value.Exists() && value.Value() != nil {
				data.ProductTypes = helpers.GetStringList(value.Array())
			} else {
				data.ProductTypes = types.ListNull(types.StringType)
			}
			(*parent).Networks = append((*parent).Networks, data)
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
func (data *OrganizationIntegrationsXDRNetworks) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Networks); i++ {
		keys := [...]string{"networkId"}
		keyValues := [...]string{data.Networks[i].NetworkId.ValueString()}

		parent := &data
		data := (*parent).Networks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("networks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Networks[%d] = %+v",
				i,
				(*parent).Networks[i],
			))
			(*parent).Networks = slices.Delete((*parent).Networks, i, i+1)
			i--

			continue
		}
		if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("productTypes"); value.Exists() && !data.ProductTypes.IsNull() {
			data.ProductTypes = helpers.GetStringList(value.Array())
		} else {
			data.ProductTypes = types.ListNull(types.StringType)
		}
		(*parent).Networks[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationIntegrationsXDRNetworks) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationIntegrationsXDRNetworksIdentity) toIdentity(ctx context.Context, plan *OrganizationIntegrationsXDRNetworks) {
	data.OrganizationId = plan.OrganizationId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationIntegrationsXDRNetworks) fromIdentity(ctx context.Context, identity *OrganizationIntegrationsXDRNetworksIdentity) {
	data.OrganizationId = identity.OrganizationId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationIntegrationsXDRNetworks) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
