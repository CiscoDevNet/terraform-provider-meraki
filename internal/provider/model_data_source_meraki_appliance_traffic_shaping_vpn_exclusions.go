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
// It emits a separate model - DataSourceApplianceTrafficShapingVPNExclusions - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceTrafficShapingVPNExclusions struct {
	Id                types.String                                                      `tfsdk:"id"`
	NetworkId         types.String                                                      `tfsdk:"network_id"`
	Custom            []DataSourceApplianceTrafficShapingVPNExclusionsCustom            `tfsdk:"custom"`
	MajorApplications []DataSourceApplianceTrafficShapingVPNExclusionsMajorApplications `tfsdk:"major_applications"`
}

type DataSourceApplianceTrafficShapingVPNExclusionsCustom struct {
	Destination types.String `tfsdk:"destination"`
	Port        types.String `tfsdk:"port"`
	Protocol    types.String `tfsdk:"protocol"`
}

type DataSourceApplianceTrafficShapingVPNExclusionsMajorApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceTrafficShapingVPNExclusions) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/vpnExclusions", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceTrafficShapingVPNExclusions) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("custom"); value.Exists() && value.Value() != nil {
		data.Custom = make([]DataSourceApplianceTrafficShapingVPNExclusionsCustom, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceTrafficShapingVPNExclusionsCustom{}
			if value := res.Get("destination"); value.Exists() && value.Value() != nil {
				data.Destination = types.StringValue(value.String())
			} else {
				data.Destination = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && value.Value() != nil {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			(*parent).Custom = append((*parent).Custom, data)
			return true
		})
	}
	if value := res.Get("majorApplications"); value.Exists() && value.Value() != nil {
		data.MajorApplications = make([]DataSourceApplianceTrafficShapingVPNExclusionsMajorApplications, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceTrafficShapingVPNExclusionsMajorApplications{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).MajorApplications = append((*parent).MajorApplications, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
