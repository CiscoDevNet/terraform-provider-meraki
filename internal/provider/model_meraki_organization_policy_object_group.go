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

type OrganizationPolicyObjectGroup struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Category       types.String `tfsdk:"category"`
	Name           types.String `tfsdk:"name"`
	ObjectIds      types.Set    `tfsdk:"object_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationPolicyObjectGroup) getPath() string {
	return fmt.Sprintf("/organizations/%v/policyObjects/groups", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationPolicyObjectGroup) toBody(ctx context.Context, state OrganizationPolicyObjectGroup) string {
	body := ""
	if !data.Category.IsNull() {
		body, _ = sjson.Set(body, "category", data.Category.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.ObjectIds.IsNull() {
		var values []int64
		data.ObjectIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "objectIds", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationPolicyObjectGroup) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("category"); value.Exists() && value.Value() != nil {
		data.Category = types.StringValue(value.String())
	} else {
		data.Category = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("objectIds"); value.Exists() && value.Value() != nil {
		data.ObjectIds = helpers.GetInt64Set(value.Array())
	} else {
		data.ObjectIds = types.SetNull(types.Int64Type)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationPolicyObjectGroup) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("category"); value.Exists() && !data.Category.IsNull() {
		data.Category = types.StringValue(value.String())
	} else {
		data.Category = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("objectIds"); value.Exists() && !data.ObjectIds.IsNull() {
		data.ObjectIds = helpers.GetInt64Set(value.Array())
	} else {
		data.ObjectIds = types.SetNull(types.Int64Type)
	}
}

// End of section. //template:end fromBodyPartial
