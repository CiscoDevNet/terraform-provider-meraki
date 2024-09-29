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

type NetworkWebhookPayloadTemplate struct {
	Id          types.String                           `tfsdk:"id"`
	NetworkId   types.String                           `tfsdk:"network_id"`
	Body        types.String                           `tfsdk:"body"`
	BodyFile    types.String                           `tfsdk:"body_file"`
	HeadersFile types.String                           `tfsdk:"headers_file"`
	Name        types.String                           `tfsdk:"name"`
	Headers     []NetworkWebhookPayloadTemplateHeaders `tfsdk:"headers"`
}

type NetworkWebhookPayloadTemplateHeaders struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkWebhookPayloadTemplate) getPath() string {
	return fmt.Sprintf("/networks/%v/webhooks/payloadTemplates", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkWebhookPayloadTemplate) toBody(ctx context.Context, state NetworkWebhookPayloadTemplate) string {
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

func (data *NetworkWebhookPayloadTemplate) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Headers = make([]NetworkWebhookPayloadTemplateHeaders, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkWebhookPayloadTemplateHeaders{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkWebhookPayloadTemplate) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBodyPartial
