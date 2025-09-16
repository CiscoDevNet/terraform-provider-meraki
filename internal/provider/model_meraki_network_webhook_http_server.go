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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkWebhookHTTPServer struct {
	Id                               types.String `tfsdk:"id"`
	NetworkId                        types.String `tfsdk:"network_id"`
	Name                             types.String `tfsdk:"name"`
	SharedSecret                     types.String `tfsdk:"shared_secret"`
	Url                              types.String `tfsdk:"url"`
	PayloadTemplateName              types.String `tfsdk:"payload_template_name"`
	PayloadTemplatePayloadTemplateId types.String `tfsdk:"payload_template_payload_template_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkWebhookHTTPServer) getPath() string {
	return fmt.Sprintf("/networks/%v/webhooks/httpServers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkWebhookHTTPServer) toBody(ctx context.Context, state NetworkWebhookHTTPServer) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.SharedSecret.IsNull() {
		body, _ = sjson.Set(body, "sharedSecret", data.SharedSecret.ValueString())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, "url", data.Url.ValueString())
	}
	if !data.PayloadTemplateName.IsNull() {
		body, _ = sjson.Set(body, "payloadTemplate.name", data.PayloadTemplateName.ValueString())
	}
	if !data.PayloadTemplatePayloadTemplateId.IsNull() {
		body, _ = sjson.Set(body, "payloadTemplate.payloadTemplateId", data.PayloadTemplatePayloadTemplateId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkWebhookHTTPServer) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("url"); value.Exists() && value.Value() != nil {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get("payloadTemplate.name"); value.Exists() && value.Value() != nil {
		data.PayloadTemplateName = types.StringValue(value.String())
	} else {
		data.PayloadTemplateName = types.StringNull()
	}
	if value := res.Get("payloadTemplate.payloadTemplateId"); value.Exists() && value.Value() != nil {
		data.PayloadTemplatePayloadTemplateId = types.StringValue(value.String())
	} else {
		data.PayloadTemplatePayloadTemplateId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkWebhookHTTPServer) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("url"); value.Exists() && !data.Url.IsNull() {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get("payloadTemplate.name"); value.Exists() && !data.PayloadTemplateName.IsNull() {
		data.PayloadTemplateName = types.StringValue(value.String())
	} else {
		data.PayloadTemplateName = types.StringNull()
	}
	if value := res.Get("payloadTemplate.payloadTemplateId"); value.Exists() && !data.PayloadTemplatePayloadTemplateId.IsNull() {
		data.PayloadTemplatePayloadTemplateId = types.StringValue(value.String())
	} else {
		data.PayloadTemplatePayloadTemplateId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkWebhookHTTPServer) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data NetworkWebhookHTTPServer) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
