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

type ResourceNetworkWebhookPayloadTemplates struct {
	Id             types.String                                  `tfsdk:"id"`
	OrganizationId types.String                                  `tfsdk:"organization_id"`
	NetworkId      types.String                                  `tfsdk:"network_id"`
	Items          []ResourceNetworkWebhookPayloadTemplatesItems `tfsdk:"items"`
}

type ResourceNetworkWebhookPayloadTemplatesItems struct {
	Id          types.String                                    `tfsdk:"id"`
	Body        types.String                                    `tfsdk:"body"`
	BodyFile    types.String                                    `tfsdk:"body_file"`
	HeadersFile types.String                                    `tfsdk:"headers_file"`
	Name        types.String                                    `tfsdk:"name"`
	Headers     []ResourceNetworkWebhookPayloadTemplatesHeaders `tfsdk:"headers"`
}

type ResourceNetworkWebhookPayloadTemplatesHeaders struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

type ResourceNetworkWebhookPayloadTemplatesIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceNetworkWebhookPayloadTemplates) getPath() string {
	return fmt.Sprintf("/networks/%v/webhooks/payloadTemplates", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceNetworkWebhookPayloadTemplatesItems) toBody(ctx context.Context, state ResourceNetworkWebhookPayloadTemplatesItems) string {
	body := ""
	if !data.Body.IsNull() {
		body, _ = sjson.Set(body, "body", data.Body.ValueString())
	}
	if !data.BodyFile.IsNull() {
		body, _ = sjson.Set(body, "bodyFile", data.BodyFile.ValueString())
	}
	if !data.HeadersFile.IsNull() {
		body, _ = sjson.Set(body, "headersFile", data.HeadersFile.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.Headers) > 0 {
		body, _ = sjson.Set(body, "headers", []interface{}{})
		for _, item := range data.Headers {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Template.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "template", item.Template.ValueString())
			}
			body, _ = sjson.SetRaw(body, "headers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceNetworkWebhookPayloadTemplates) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceNetworkWebhookPayloadTemplatesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceNetworkWebhookPayloadTemplatesItems{}
		if value := res.Get("body"); value.Exists() && value.Value() != nil {
			data.Body = types.StringValue(value.String())
		} else {
			data.Body = types.StringNull()
		}
		if value := res.Get("bodyFile"); value.Exists() && value.Value() != nil {
			data.BodyFile = types.StringValue(value.String())
		} else {
			data.BodyFile = types.StringNull()
		}
		if value := res.Get("headersFile"); value.Exists() && value.Value() != nil {
			data.HeadersFile = types.StringValue(value.String())
		} else {
			data.HeadersFile = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("headers"); value.Exists() && value.Value() != nil {
			data.Headers = make([]ResourceNetworkWebhookPayloadTemplatesHeaders, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceNetworkWebhookPayloadTemplatesHeaders{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("template"); value.Exists() && value.Value() != nil {
					data.Template = types.StringValue(value.String())
				} else {
					data.Template = types.StringNull()
				}
				(*parent).Headers = append((*parent).Headers, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("payloadTemplateId").String())
		index++
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceNetworkWebhookPayloadTemplates) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("payloadTemplateId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("body"); value.Exists() && !data.Body.IsNull() {
			data.Body = types.StringValue(value.String())
		} else {
			data.Body = types.StringNull()
		}
		if value := res.Get("bodyFile"); value.Exists() && !data.BodyFile.IsNull() {
			data.BodyFile = types.StringValue(value.String())
		} else {
			data.BodyFile = types.StringNull()
		}
		if value := res.Get("headersFile"); value.Exists() && !data.HeadersFile.IsNull() {
			data.HeadersFile = types.StringValue(value.String())
		} else {
			data.HeadersFile = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		for i := 0; i < len(data.Headers); i++ {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Headers[i].Name.ValueString()}

			parent := &data
			data := (*parent).Headers[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("headers").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Headers[%d] = %+v",
					i,
					(*parent).Headers[i],
				))
				(*parent).Headers = slices.Delete((*parent).Headers, i, i+1)
				i--

				continue
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("template"); value.Exists() && !data.Template.IsNull() {
				data.Template = types.StringValue(value.String())
			} else {
				data.Template = types.StringNull()
			}
			(*parent).Headers[i] = data
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceNetworkWebhookPayloadTemplates) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceNetworkWebhookPayloadTemplates) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("payloadTemplateId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("body"); value.Exists() && value.Value() != nil {
			data.Body = types.StringValue(value.String())
		} else {
			data.Body = types.StringNull()
		}
		if value := res.Get("bodyFile"); value.Exists() && value.Value() != nil {
			data.BodyFile = types.StringValue(value.String())
		} else {
			data.BodyFile = types.StringNull()
		}
		if value := res.Get("headersFile"); value.Exists() && value.Value() != nil {
			data.HeadersFile = types.StringValue(value.String())
		} else {
			data.HeadersFile = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("headers"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.Headers = make([]ResourceNetworkWebhookPayloadTemplatesHeaders, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceNetworkWebhookPayloadTemplatesHeaders{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("template"); value.Exists() && value.Value() != nil {
					data.Template = types.StringValue(value.String())
				} else {
					data.Template = types.StringNull()
				}
				(*parent).Headers = append((*parent).Headers, data)
				return true
			})
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ResourceNetworkWebhookPayloadTemplatesIdentity) toIdentity(ctx context.Context, plan *ResourceNetworkWebhookPayloadTemplates) {
	data.OrganizationId = plan.OrganizationId
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceNetworkWebhookPayloadTemplates) fromIdentity(ctx context.Context, identity *ResourceNetworkWebhookPayloadTemplatesIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceNetworkWebhookPayloadTemplates) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceNetworkWebhookPayloadTemplates) hasChanges(ctx context.Context, state *ResourceNetworkWebhookPayloadTemplates, id string) bool {
	hasChanges := false

	item := ResourceNetworkWebhookPayloadTemplatesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceNetworkWebhookPayloadTemplatesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Body.Equal(stateItem.Body) {
		hasChanges = true
	}
	if !item.BodyFile.Equal(stateItem.BodyFile) {
		hasChanges = true
	}
	if !item.HeadersFile.Equal(stateItem.HeadersFile) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if len(item.Headers) != len(stateItem.Headers) {
		hasChanges = true
	} else {
		for i := range item.Headers {
			if !item.Headers[i].Name.Equal(stateItem.Headers[i].Name) {
				hasChanges = true
			}
			if !item.Headers[i].Template.Equal(stateItem.Headers[i].Template) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
