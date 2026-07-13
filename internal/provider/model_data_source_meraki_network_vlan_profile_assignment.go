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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceNetworkVLANProfileAssignment - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkVLANProfileAssignment struct {
	Id               types.String `tfsdk:"id"`
	NetworkId        types.String `tfsdk:"network_id"`
	VlanProfileIname types.String `tfsdk:"vlan_profile_iname"`
	Serials          types.Set    `tfsdk:"serials"`
	StackIds         types.Set    `tfsdk:"stack_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkVLANProfileAssignment) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles/assignments/reassign", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkVLANProfileAssignment) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("vlanProfile.iname"); value.Exists() && value.Value() != nil {
		data.VlanProfileIname = types.StringValue(value.String())
	} else {
		data.VlanProfileIname = types.StringNull()
	}
	if value := res.Get("serials"); value.Exists() && value.Value() != nil {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
	if value := res.Get("stackIds"); value.Exists() && value.Value() != nil {
		data.StackIds = helpers.GetStringSet(value.Array())
	} else {
		data.StackIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// getByDevicePath and fromByDeviceBody are custom (kept in sync manually with the resource-side model.go's
// equivalents) - the data source resolves current assignment by querying the byDevice endpoint and filtering by
// vlan_profile_iname, the same approach the resource uses to compute state.
func (data DataSourceNetworkVLANProfileAssignment) getByDevicePath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles/assignments/byDevice", url.QueryEscape(data.NetworkId.ValueString()))
}

func (data *DataSourceNetworkVLANProfileAssignment) fromByDeviceBody(ctx context.Context, res meraki.Res) {
	iname := data.VlanProfileIname.ValueString()
	if iname == "" {
		iname = data.Id.ValueString()
	}
	var serialValues []attr.Value
	var stackIdValues []attr.Value
	stackIdSet := map[string]bool{}
	for _, item := range res.Array() {
		profileIname := item.Get("vlanProfile.iname").String()
		if profileIname != iname {
			continue
		}
		if stackId := item.Get("stack.id").String(); stackId != "" {
			if !stackIdSet[stackId] {
				stackIdSet[stackId] = true
				stackIdValues = append(stackIdValues, types.StringValue(stackId))
			}
			continue
		}
		if serial := item.Get("serial").String(); serial != "" {
			serialValues = append(serialValues, types.StringValue(serial))
		}
	}
	if len(serialValues) > 0 {
		data.Serials = types.SetValueMust(types.StringType, serialValues)
	} else if !data.Serials.IsNull() {
		data.Serials = types.SetValueMust(types.StringType, []attr.Value{})
	}
	if !data.StackIds.IsNull() || len(stackIdValues) > 0 {
		data.StackIds = types.SetValueMust(types.StringType, stackIdValues)
	}
}
