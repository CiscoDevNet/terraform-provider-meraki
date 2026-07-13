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
// It emits a separate model - DataSourceWirelessLocationScanningReceiver - used only by the data source, always including every
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

type DataSourceWirelessLocationScanningReceiver struct {
	Id                    types.String `tfsdk:"id"`
	OrganizationId        types.String `tfsdk:"organization_id"`
	SharedSecret          types.String `tfsdk:"shared_secret"`
	SharedSecretWo        types.String `tfsdk:"shared_secret_wo"`
	SharedSecretWoVersion types.Int64  `tfsdk:"shared_secret_wo_version"`
	Url                   types.String `tfsdk:"url"`
	Version               types.String `tfsdk:"version"`
	NetworkId             types.String `tfsdk:"network_id"`
	RadioType             types.String `tfsdk:"radio_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessLocationScanningReceiver) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/location/scanning/receivers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessLocationScanningReceiver) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("url"); value.Exists() && value.Value() != nil {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get("version"); value.Exists() && value.Value() != nil {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
	}
	if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("radio.type"); value.Exists() && value.Value() != nil {
		data.RadioType = types.StringValue(value.String())
	} else {
		data.RadioType = types.StringNull()
	}
}

// End of section. //template:end fromBody
