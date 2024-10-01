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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkWebhookPayloadTemplates struct {
	NetworkId types.String                          `tfsdk:"network_id"`
	Items     []NetworkWebhookPayloadTemplatesItems `tfsdk:"items"`
}

type NetworkWebhookPayloadTemplatesItems struct {
	Id          types.String                            `tfsdk:"id"`
	Body        types.String                            `tfsdk:"body"`
	BodyFile    types.String                            `tfsdk:"body_file"`
	HeadersFile types.String                            `tfsdk:"headers_file"`
	Name        types.String                            `tfsdk:"name"`
	Headers     []NetworkWebhookPayloadTemplatesHeaders `tfsdk:"headers"`
}

type NetworkWebhookPayloadTemplatesHeaders struct {
	Name     types.String `tfsdk:"name"`
	Template types.String `tfsdk:"template"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkWebhookPayloadTemplates) getPath() string {
	return fmt.Sprintf("/networks/%v/webhooks/payloadTemplates", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkWebhookPayloadTemplates) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]NetworkWebhookPayloadTemplatesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := NetworkWebhookPayloadTemplatesItems{}
		data.Id = types.StringValue(res.Get("payloadTemplateId").String())
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
			data.Headers = make([]NetworkWebhookPayloadTemplatesHeaders, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkWebhookPayloadTemplatesHeaders{}
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
}

// End of section. //template:end fromBody
