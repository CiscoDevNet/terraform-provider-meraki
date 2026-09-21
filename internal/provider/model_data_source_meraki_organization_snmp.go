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
// It emits a separate model - DataSourceOrganizationSNMP - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationSNMP struct {
	Id                  types.String `tfsdk:"id"`
	OrganizationId      types.String `tfsdk:"organization_id"`
	V2cEnabled          types.Bool   `tfsdk:"v2c_enabled"`
	V3AuthMode          types.String `tfsdk:"v3_auth_mode"`
	V3AuthPass          types.String `tfsdk:"v3_auth_pass"`
	V3AuthPassWo        types.String `tfsdk:"v3_auth_pass_wo"`
	V3AuthPassWoVersion types.Int64  `tfsdk:"v3_auth_pass_wo_version"`
	V3Enabled           types.Bool   `tfsdk:"v3_enabled"`
	V3PrivMode          types.String `tfsdk:"v3_priv_mode"`
	V3PrivPass          types.String `tfsdk:"v3_priv_pass"`
	V3PrivPassWo        types.String `tfsdk:"v3_priv_pass_wo"`
	V3PrivPassWoVersion types.Int64  `tfsdk:"v3_priv_pass_wo_version"`
	PeerIps             types.Set    `tfsdk:"peer_ips"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationSNMP) getPath() string {
	return fmt.Sprintf("/organizations/%v/snmp", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationSNMP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("v2cEnabled"); value.Exists() && value.Value() != nil {
		data.V2cEnabled = types.BoolValue(value.Bool())
	} else {
		data.V2cEnabled = types.BoolNull()
	}
	if value := res.Get("v3AuthMode"); value.Exists() && value.Value() != nil {
		data.V3AuthMode = types.StringValue(value.String())
	} else {
		data.V3AuthMode = types.StringNull()
	}
	if value := res.Get("v3Enabled"); value.Exists() && value.Value() != nil {
		data.V3Enabled = types.BoolValue(value.Bool())
	} else {
		data.V3Enabled = types.BoolNull()
	}
	if value := res.Get("v3PrivMode"); value.Exists() && value.Value() != nil {
		data.V3PrivMode = types.StringValue(value.String())
	} else {
		data.V3PrivMode = types.StringNull()
	}
	if value := res.Get("peerIps"); value.Exists() && value.Value() != nil {
		data.PeerIps = helpers.GetStringSet(value.Array())
	} else {
		data.PeerIps = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody
