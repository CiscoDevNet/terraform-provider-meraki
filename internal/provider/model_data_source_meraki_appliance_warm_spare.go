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
// It emits a separate model - DataSourceApplianceWarmSpare - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceWarmSpare struct {
	Id            types.String `tfsdk:"id"`
	NetworkId     types.String `tfsdk:"network_id"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	PrimarySerial types.String `tfsdk:"primary_serial"`
	SpareSerial   types.String `tfsdk:"spare_serial"`
	UplinkMode    types.String `tfsdk:"uplink_mode"`
	VirtualIp1    types.String `tfsdk:"virtual_ip1"`
	VirtualIp2    types.String `tfsdk:"virtual_ip2"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceWarmSpare) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/warmSpare", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// fromBody is de-generated (kept in sync manually with the resource-side model_resource.go's custom fromBody) because
// the API nests the virtual IPs under wan1.ip/wan2.ip, not the virtualIp1/virtualIp2 paths the OpenAPI spec implies.
func (data *DataSourceApplianceWarmSpare) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("primarySerial"); value.Exists() && value.Value() != nil {
		data.PrimarySerial = types.StringValue(value.String())
	} else {
		data.PrimarySerial = types.StringNull()
	}
	if value := res.Get("spareSerial"); value.Exists() && value.Value() != nil {
		data.SpareSerial = types.StringValue(value.String())
	} else {
		data.SpareSerial = types.StringNull()
	}
	if value := res.Get("uplinkMode"); value.Exists() && value.Value() != nil {
		data.UplinkMode = types.StringValue(value.String())
	} else {
		data.UplinkMode = types.StringNull()
	}
	if value := res.Get("wan1.ip"); value.Exists() && value.Value() != nil {
		data.VirtualIp1 = types.StringValue(value.String())
	} else {
		data.VirtualIp1 = types.StringNull()
	}
	if value := res.Get("wan2.ip"); value.Exists() && value.Value() != nil {
		data.VirtualIp2 = types.StringValue(value.String())
	} else {
		data.VirtualIp2 = types.StringNull()
	}
}
