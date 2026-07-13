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
// It emits a separate model - DataSourceNetworkClientSplashAuthorizationStatus - used only by the data source, always including every
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

type DataSourceNetworkClientSplashAuthorizationStatus struct {
	Id                  types.String `tfsdk:"id"`
	NetworkId           types.String `tfsdk:"network_id"`
	ClientId            types.String `tfsdk:"client_id"`
	Ssids0IsAuthorized  types.Bool   `tfsdk:"ssids_0_is_authorized"`
	Ssids1IsAuthorized  types.Bool   `tfsdk:"ssids_1_is_authorized"`
	Ssids10IsAuthorized types.Bool   `tfsdk:"ssids_10_is_authorized"`
	Ssids11IsAuthorized types.Bool   `tfsdk:"ssids_11_is_authorized"`
	Ssids12IsAuthorized types.Bool   `tfsdk:"ssids_12_is_authorized"`
	Ssids13IsAuthorized types.Bool   `tfsdk:"ssids_13_is_authorized"`
	Ssids14IsAuthorized types.Bool   `tfsdk:"ssids_14_is_authorized"`
	Ssids2IsAuthorized  types.Bool   `tfsdk:"ssids_2_is_authorized"`
	Ssids3IsAuthorized  types.Bool   `tfsdk:"ssids_3_is_authorized"`
	Ssids4IsAuthorized  types.Bool   `tfsdk:"ssids_4_is_authorized"`
	Ssids5IsAuthorized  types.Bool   `tfsdk:"ssids_5_is_authorized"`
	Ssids6IsAuthorized  types.Bool   `tfsdk:"ssids_6_is_authorized"`
	Ssids7IsAuthorized  types.Bool   `tfsdk:"ssids_7_is_authorized"`
	Ssids8IsAuthorized  types.Bool   `tfsdk:"ssids_8_is_authorized"`
	Ssids9IsAuthorized  types.Bool   `tfsdk:"ssids_9_is_authorized"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkClientSplashAuthorizationStatus) getPath() string {
	return fmt.Sprintf("/networks/%v/clients/%v/splashAuthorizationStatus", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.ClientId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkClientSplashAuthorizationStatus) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ssids.0.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids0IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids0IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.1.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids1IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids1IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.10.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids10IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids10IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.11.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids11IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids11IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.12.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids12IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids12IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.13.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids13IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids13IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.14.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids14IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids14IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.2.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids2IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids2IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.3.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids3IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids3IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.4.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids4IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids4IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.5.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids5IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids5IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.6.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids6IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids6IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.7.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids7IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids7IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.8.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids8IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids8IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.9.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids9IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids9IsAuthorized = types.BoolNull()
	}
}

// End of section. //template:end fromBody
