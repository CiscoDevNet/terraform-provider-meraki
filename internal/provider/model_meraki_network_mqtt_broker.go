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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkMQTTBroker struct {
	Id                         types.String `tfsdk:"id"`
	NetworkId                  types.String `tfsdk:"network_id"`
	Host                       types.String `tfsdk:"host"`
	Name                       types.String `tfsdk:"name"`
	Port                       types.Int64  `tfsdk:"port"`
	AuthenticationPassword     types.String `tfsdk:"authentication_password"`
	AuthenticationUsername     types.String `tfsdk:"authentication_username"`
	SecurityMode               types.String `tfsdk:"security_mode"`
	SecurityTlsCaCertificate   types.String `tfsdk:"security_tls_ca_certificate"`
	SecurityTlsVerifyHostnames types.Bool   `tfsdk:"security_tls_verify_hostnames"`
}

type NetworkMQTTBrokerIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	Id        types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkMQTTBroker) getPath() string {
	return fmt.Sprintf("/networks/%v/mqttBrokers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkMQTTBroker) toBody(ctx context.Context, state NetworkMQTTBroker) string {
	body := ""
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, "host", data.Host.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, "port", data.Port.ValueInt64())
	}
	if !data.AuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "authentication.password", data.AuthenticationPassword.ValueString())
	}
	if !data.AuthenticationUsername.IsNull() {
		body, _ = sjson.Set(body, "authentication.username", data.AuthenticationUsername.ValueString())
	}
	if !data.SecurityMode.IsNull() {
		body, _ = sjson.Set(body, "security.mode", data.SecurityMode.ValueString())
	}
	if !data.SecurityTlsCaCertificate.IsNull() {
		body, _ = sjson.Set(body, "security.tls.caCertificate", data.SecurityTlsCaCertificate.ValueString())
	}
	if !data.SecurityTlsVerifyHostnames.IsNull() {
		body, _ = sjson.Set(body, "security.tls.verifyHostnames", data.SecurityTlsVerifyHostnames.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkMQTTBroker) fromBody(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkMQTTBroker) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
		data.Host = types.StringValue(value.String())
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get("authentication.username"); value.Exists() && !data.AuthenticationUsername.IsNull() {
		data.AuthenticationUsername = types.StringValue(value.String())
	} else {
		data.AuthenticationUsername = types.StringNull()
	}
	if value := res.Get("security.mode"); value.Exists() && !data.SecurityMode.IsNull() {
		data.SecurityMode = types.StringValue(value.String())
	} else {
		data.SecurityMode = types.StringNull()
	}
	if value := res.Get("security.tls.verifyHostnames"); value.Exists() && !data.SecurityTlsVerifyHostnames.IsNull() {
		data.SecurityTlsVerifyHostnames = types.BoolValue(value.Bool())
	} else {
		data.SecurityTlsVerifyHostnames = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkMQTTBroker) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkMQTTBrokerIdentity) toIdentity(ctx context.Context, plan *NetworkMQTTBroker) {
	data.NetworkId = plan.NetworkId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkMQTTBroker) fromIdentity(ctx context.Context, identity *NetworkMQTTBrokerIdentity) {
	data.NetworkId = identity.NetworkId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkMQTTBroker) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
