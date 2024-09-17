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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchStack struct {
	Id        types.String `tfsdk:"id"`
	NetworkId types.String `tfsdk:"network_id"`
	Name      types.String `tfsdk:"name"`
	Serials   types.Set    `tfsdk:"serials"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchStack) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

func (data SwitchStack) getAddPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/add", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Id.ValueString()))
}

func (data SwitchStack) getRemovePath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/remove", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Id.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchStack) toBody(ctx context.Context, state SwitchStack) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchStack) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("serials"); value.Exists() {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchStack) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("serials"); value.Exists() && !data.Serials.IsNull() {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial