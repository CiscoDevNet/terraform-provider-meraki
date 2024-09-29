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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationPolicyObject struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Category       types.String `tfsdk:"category"`
	Cidr           types.String `tfsdk:"cidr"`
	Fqdn           types.String `tfsdk:"fqdn"`
	Ip             types.String `tfsdk:"ip"`
	Mask           types.String `tfsdk:"mask"`
	Name           types.String `tfsdk:"name"`
	Type           types.String `tfsdk:"type"`
	GroupIds       types.Set    `tfsdk:"group_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationPolicyObject) getPath() string {
	return fmt.Sprintf("/organizations/%v/policyObjects", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationPolicyObject) toBody(ctx context.Context, state OrganizationPolicyObject) string {
	body := ""
	if !data.Category.IsNull() {
		body, _ = sjson.Set(body, "category", data.Category.ValueString())
	}
	if !data.Cidr.IsNull() {
		body, _ = sjson.Set(body, "cidr", data.Cidr.ValueString())
	}
	if !data.Fqdn.IsNull() {
		body, _ = sjson.Set(body, "fqdn", data.Fqdn.ValueString())
	}
	if !data.Ip.IsNull() {
		body, _ = sjson.Set(body, "ip", data.Ip.ValueString())
	}
	if !data.Mask.IsNull() {
		body, _ = sjson.Set(body, "mask", data.Mask.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.GroupIds.IsNull() {
		var values []string
		data.GroupIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "groupIds", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationPolicyObject) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("category"); value.Exists() && value.Value() != nil {
		data.Category = types.StringValue(value.String())
	} else {
		data.Category = types.StringNull()
	}
	if value := res.Get("cidr"); value.Exists() && value.Value() != nil {
		data.Cidr = types.StringValue(value.String())
	} else {
		data.Cidr = types.StringNull()
	}
	if value := res.Get("fqdn"); value.Exists() && value.Value() != nil {
		data.Fqdn = types.StringValue(value.String())
	} else {
		data.Fqdn = types.StringNull()
	}
	if value := res.Get("ip"); value.Exists() && value.Value() != nil {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get("mask"); value.Exists() && value.Value() != nil {
		data.Mask = types.StringValue(value.String())
	} else {
		data.Mask = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && value.Value() != nil {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("groupIds"); value.Exists() && value.Value() != nil {
		data.GroupIds = helpers.GetStringSet(value.Array())
	} else {
		data.GroupIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationPolicyObject) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("category"); value.Exists() && !data.Category.IsNull() {
		data.Category = types.StringValue(value.String())
	} else {
		data.Category = types.StringNull()
	}
	if value := res.Get("cidr"); value.Exists() && !data.Cidr.IsNull() {
		data.Cidr = types.StringValue(value.String())
	} else {
		data.Cidr = types.StringNull()
	}
	if value := res.Get("fqdn"); value.Exists() && !data.Fqdn.IsNull() {
		data.Fqdn = types.StringValue(value.String())
	} else {
		data.Fqdn = types.StringNull()
	}
	if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get("mask"); value.Exists() && !data.Mask.IsNull() {
		data.Mask = types.StringValue(value.String())
	} else {
		data.Mask = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("groupIds"); value.Exists() && !data.GroupIds.IsNull() {
		data.GroupIds = helpers.GetStringSet(value.Array())
	} else {
		data.GroupIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
