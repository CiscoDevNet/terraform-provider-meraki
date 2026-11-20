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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationSASESiteAttachment struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	RegionSlug     types.String `tfsdk:"region_slug"`
	SiteId         types.String `tfsdk:"site_id"`
}

type OrganizationSASESiteAttachmentIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationSASESiteAttachment) getPath() string {
	return fmt.Sprintf("/organizations/%v/sase/sites", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

func (data OrganizationSASESiteAttachment) getCreatePath() string {
	return fmt.Sprintf("/organizations/%v/sase/sites/attach", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data OrganizationSASESiteAttachment) getDeletePath() string {
	return fmt.Sprintf("/organizations/%v/sase/sites/detach", url.QueryEscape(data.OrganizationId.ValueString()))
}

// toCreateBody returns the body of the create endpoint, which takes a list of items.
func (data OrganizationSASESiteAttachment) toCreateBody() string {
	item, _ := sjson.Set("", "network.id", data.NetworkId.ValueString())
	item, _ = sjson.Set(item, "region.slug", data.RegionSlug.ValueString())
	body, _ := sjson.SetRaw("{}", "items.-1", item)
	return body
}

// toDeleteBody returns the body of the delete endpoint for the given site ID.
func (data OrganizationSASESiteAttachment) toDeleteBody(id string) string {
	item, _ := sjson.Set("", "siteId", id)
	body, _ := sjson.SetRaw("{}", "items.-1", item)
	return body
}

// findItem returns the site matching this resource from the list response, if any.
func (data OrganizationSASESiteAttachment) findItem(res meraki.Res) (gjson.Result, bool) {
	items := res.Get("items")
	if !items.Exists() {
		items = res.Result
	}
	var found gjson.Result
	items.ForEach(func(_, v gjson.Result) bool {
		if v.Get("network.id").String() == data.NetworkId.ValueString() {
			found = v
			return false
		}
		return true
	})
	return found, found.Exists()
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationSASESiteAttachment) toBody(ctx context.Context, state OrganizationSASESiteAttachment) string {
	body := ""
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "network.id", data.NetworkId.ValueString())
	}
	if !data.RegionSlug.IsNull() {
		body, _ = sjson.Set(body, "region.slug", data.RegionSlug.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationSASESiteAttachment) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("region.slug"); value.Exists() && value.Value() != nil {
		data.RegionSlug = types.StringValue(value.String())
	} else {
		data.RegionSlug = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && value.Value() != nil {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationSASESiteAttachment) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("region.slug"); value.Exists() && !data.RegionSlug.IsNull() {
		data.RegionSlug = types.StringValue(value.String())
	} else {
		data.RegionSlug = types.StringNull()
	}
	if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationSASESiteAttachment) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
	if data.SiteId.IsUnknown() {
		if value := res.Get("siteId"); value.Exists() && !data.SiteId.IsNull() {
			data.SiteId = types.StringValue(value.String())
		} else {
			data.SiteId = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationSASESiteAttachmentIdentity) toIdentity(ctx context.Context, plan *OrganizationSASESiteAttachment) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationSASESiteAttachment) fromIdentity(ctx context.Context, identity *OrganizationSASESiteAttachmentIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationSASESiteAttachment) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
