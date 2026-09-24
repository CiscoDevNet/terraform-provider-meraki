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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

// ActionConnectApplianceUmbrellaAccount models the `config` block of an action invocation. Unlike a
// resource/data-source model, this is never persisted: it exists only for the duration of
// a single Invoke call, so it has no computed/read-only fields and no identity.
type ActionConnectApplianceUmbrellaAccount struct {
	NetworkId types.String `tfsdk:"network_id"`
	ApiKey    types.String `tfsdk:"api_key"`
	ApiSecret types.String `tfsdk:"api_secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ActionConnectApplianceUmbrellaAccount) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/umbrella/account/connect", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ActionConnectApplianceUmbrellaAccount) toBody(ctx context.Context) string {
	body := ""
	if !data.ApiKey.IsNull() {
		body, _ = sjson.Set(body, "api.key", data.ApiKey.ValueString())
	}
	if !data.ApiSecret.IsNull() {
		body, _ = sjson.Set(body, "api.secret", data.ApiSecret.ValueString())
	}
	return body
}

// End of section. //template:end toBody
