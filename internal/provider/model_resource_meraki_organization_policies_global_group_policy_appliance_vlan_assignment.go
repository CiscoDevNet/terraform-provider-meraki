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

type OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	PolicyId       types.String `tfsdk:"policy_id"`
	InterfaceId    types.String `tfsdk:"interface_id"`
}

type OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) getPath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/group/policies/appliance/vlans/assignments", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) getCreatePath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/group/policies/appliance/vlans/assign", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) getDeletePath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/group/policies/appliance/vlans/remove", url.QueryEscape(data.OrganizationId.ValueString()))
}

// toAssignBody returns the body of the `assign` and `remove` endpoints, which take the policy and
// a list with the appliance VLAN to assign or remove.
func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) toAssignBody() string {
	item, _ := sjson.Set("", "interfaceId", data.InterfaceId.ValueString())
	body, _ := sjson.Set("", "policy.id", data.PolicyId.ValueString())
	body, _ = sjson.SetRaw(body, "vlans.-1", item)
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) toBody(ctx context.Context, state OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) string {
	body := ""
	if !data.PolicyId.IsNull() {
		body, _ = sjson.Set(body, "policyId", data.PolicyId.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "interfaceId", data.InterfaceId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("policyId"); value.Exists() && value.Value() != nil {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	if value := res.Get("interfaceId"); value.Exists() && value.Value() != nil {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("policyId"); value.Exists() && !data.PolicyId.IsNull() {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	if value := res.Get("interfaceId"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentIdentity) toIdentity(ctx context.Context, plan *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) fromIdentity(ctx context.Context, identity *OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignmentIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationPoliciesGlobalGroupPolicyApplianceVLANAssignment) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
