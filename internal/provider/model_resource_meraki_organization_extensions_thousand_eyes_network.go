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

type OrganizationExtensionsThousandEyesNetwork struct {
	Id             types.String                                     `tfsdk:"id"`
	OrganizationId types.String                                     `tfsdk:"organization_id"`
	Enabled        types.Bool                                       `tfsdk:"enabled"`
	NetworkId      types.String                                     `tfsdk:"network_id"`
	Tests          []OrganizationExtensionsThousandEyesNetworkTests `tfsdk:"tests"`
}

type OrganizationExtensionsThousandEyesNetworkTests struct {
	NetworkId                types.String `tfsdk:"network_id"`
	TemplateId               types.String `tfsdk:"template_id"`
	TemplateUserInputsTenant types.String `tfsdk:"template_user_inputs_tenant"`
}

type OrganizationExtensionsThousandEyesNetworkIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationExtensionsThousandEyesNetwork) getPath() string {
	return fmt.Sprintf("/organizations/%v/extensions/thousandEyes/networks", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data OrganizationExtensionsThousandEyesNetwork) toBody(ctx context.Context, state OrganizationExtensionsThousandEyesNetwork) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "networkId", data.NetworkId.ValueString())
	}
	if len(data.Tests) > 0 && len(state.Tests) == 0 {
		body, _ = sjson.Set(body, "tests", []interface{}{})
		for _, item := range data.Tests {
			itemBody := ""
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "network.id", item.NetworkId.ValueString())
			}
			if !item.TemplateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "template.id", item.TemplateId.ValueString())
			}
			if !item.TemplateUserInputsTenant.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "template.userInputs.tenant", item.TemplateUserInputsTenant.ValueString())
			}
			body, _ = sjson.SetRaw(body, "tests.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationExtensionsThousandEyesNetwork) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationExtensionsThousandEyesNetwork) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationExtensionsThousandEyesNetwork) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationExtensionsThousandEyesNetworkIdentity) toIdentity(ctx context.Context, plan *OrganizationExtensionsThousandEyesNetwork) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationExtensionsThousandEyesNetwork) fromIdentity(ctx context.Context, identity *OrganizationExtensionsThousandEyesNetworkIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationExtensionsThousandEyesNetwork) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
