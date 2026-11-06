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

type RoutingVRF struct {
	Id                 types.String `tfsdk:"id"`
	OrganizationId     types.String `tfsdk:"organization_id"`
	Description        types.String `tfsdk:"description"`
	Name               types.String `tfsdk:"name"`
	RouteDistinguisher types.String `tfsdk:"route_distinguisher"`
	RouteTarget        types.String `tfsdk:"route_target"`
	ApplianceEnabled   types.Bool   `tfsdk:"appliance_enabled"`
	ApplianceIndex     types.String `tfsdk:"appliance_index"`
}

type RoutingVRFIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RoutingVRF) getPath() string {
	return fmt.Sprintf("/organizations/%v/routing/vrfs", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RoutingVRF) toBody(ctx context.Context, state RoutingVRF) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RouteDistinguisher.IsNull() {
		body, _ = sjson.Set(body, "routeDistinguisher", data.RouteDistinguisher.ValueString())
	}
	if !data.RouteTarget.IsNull() {
		body, _ = sjson.Set(body, "routeTarget", data.RouteTarget.ValueString())
	}
	if !data.ApplianceEnabled.IsNull() {
		body, _ = sjson.Set(body, "appliance.enabled", data.ApplianceEnabled.ValueBool())
	}
	if !data.ApplianceIndex.IsNull() {
		body, _ = sjson.Set(body, "appliance.index", data.ApplianceIndex.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RoutingVRF) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("routeDistinguisher"); value.Exists() && value.Value() != nil {
		data.RouteDistinguisher = types.StringValue(value.String())
	} else {
		data.RouteDistinguisher = types.StringNull()
	}
	if value := res.Get("routeTarget"); value.Exists() && value.Value() != nil {
		data.RouteTarget = types.StringValue(value.String())
	} else {
		data.RouteTarget = types.StringNull()
	}
	if value := res.Get("appliance.enabled"); value.Exists() && value.Value() != nil {
		data.ApplianceEnabled = types.BoolValue(value.Bool())
	} else {
		data.ApplianceEnabled = types.BoolNull()
	}
	if value := res.Get("appliance.index"); value.Exists() && value.Value() != nil {
		data.ApplianceIndex = types.StringValue(value.String())
	} else {
		data.ApplianceIndex = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *RoutingVRF) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("routeDistinguisher"); value.Exists() && !data.RouteDistinguisher.IsNull() {
		data.RouteDistinguisher = types.StringValue(value.String())
	} else {
		data.RouteDistinguisher = types.StringNull()
	}
	if value := res.Get("routeTarget"); value.Exists() && !data.RouteTarget.IsNull() {
		data.RouteTarget = types.StringValue(value.String())
	} else {
		data.RouteTarget = types.StringNull()
	}
	if value := res.Get("appliance.enabled"); value.Exists() && !data.ApplianceEnabled.IsNull() {
		data.ApplianceEnabled = types.BoolValue(value.Bool())
	} else {
		data.ApplianceEnabled = types.BoolNull()
	}
	if value := res.Get("appliance.index"); value.Exists() && !data.ApplianceIndex.IsNull() {
		data.ApplianceIndex = types.StringValue(value.String())
	} else {
		data.ApplianceIndex = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RoutingVRF) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *RoutingVRFIdentity) toIdentity(ctx context.Context, plan *RoutingVRF) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *RoutingVRF) fromIdentity(ctx context.Context, identity *RoutingVRFIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data RoutingVRF) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
