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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkMQTTBrokers struct {
	NetworkId types.String              `tfsdk:"network_id"`
	Items     []NetworkMQTTBrokersItems `tfsdk:"items"`
}

type NetworkMQTTBrokersItems struct {
	Id                         types.String `tfsdk:"id"`
	Host                       types.String `tfsdk:"host"`
	Name                       types.String `tfsdk:"name"`
	Port                       types.Int64  `tfsdk:"port"`
	AuthenticationPassword     types.String `tfsdk:"authentication_password"`
	AuthenticationUsername     types.String `tfsdk:"authentication_username"`
	SecurityMode               types.String `tfsdk:"security_mode"`
	SecurityTlsCaCertificate   types.String `tfsdk:"security_tls_ca_certificate"`
	SecurityTlsVerifyHostnames types.Bool   `tfsdk:"security_tls_verify_hostnames"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkMQTTBrokers) getPath() string {
	return fmt.Sprintf("/networks/%v/mqttBrokers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkMQTTBrokers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]NetworkMQTTBrokersItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := NetworkMQTTBrokersItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("host"); value.Exists() && value.Value() != nil {
			data.Host = types.StringValue(value.String())
		} else {
			data.Host = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && value.Value() != nil {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		if value := res.Get("authentication.username"); value.Exists() && value.Value() != nil {
			data.AuthenticationUsername = types.StringValue(value.String())
		} else {
			data.AuthenticationUsername = types.StringNull()
		}
		if value := res.Get("security.mode"); value.Exists() && value.Value() != nil {
			data.SecurityMode = types.StringValue(value.String())
		} else {
			data.SecurityMode = types.StringNull()
		}
		if value := res.Get("security.tls.verifyHostnames"); value.Exists() && value.Value() != nil {
			data.SecurityTlsVerifyHostnames = types.BoolValue(value.Bool())
		} else {
			data.SecurityTlsVerifyHostnames = types.BoolNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
