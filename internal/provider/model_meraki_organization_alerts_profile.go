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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationAlertsProfile struct {
	Id                       types.String  `tfsdk:"id"`
	OrganizationId           types.String  `tfsdk:"organization_id"`
	Description              types.String  `tfsdk:"description"`
	Type                     types.String  `tfsdk:"type"`
	AlertConditionBitRateBps types.Int64   `tfsdk:"alert_condition_bit_rate_bps"`
	AlertConditionDuration   types.Int64   `tfsdk:"alert_condition_duration"`
	AlertConditionInterface  types.String  `tfsdk:"alert_condition_interface"`
	AlertConditionJitterMs   types.Int64   `tfsdk:"alert_condition_jitter_ms"`
	AlertConditionLatencyMs  types.Int64   `tfsdk:"alert_condition_latency_ms"`
	AlertConditionLossRatio  types.Float64 `tfsdk:"alert_condition_loss_ratio"`
	AlertConditionMos        types.Float64 `tfsdk:"alert_condition_mos"`
	AlertConditionWindow     types.Int64   `tfsdk:"alert_condition_window"`
	RecipientsEmails         types.Set     `tfsdk:"recipients_emails"`
	RecipientsHttpServerIds  types.Set     `tfsdk:"recipients_http_server_ids"`
	NetworkTags              types.Set     `tfsdk:"network_tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationAlertsProfile) getPath() string {
	return fmt.Sprintf("/organizations/%v/alerts/profiles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationAlertsProfile) toBody(ctx context.Context, state OrganizationAlertsProfile) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.AlertConditionBitRateBps.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.bit_rate_bps", data.AlertConditionBitRateBps.ValueInt64())
	}
	if !data.AlertConditionDuration.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.duration", data.AlertConditionDuration.ValueInt64())
	}
	if !data.AlertConditionInterface.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.interface", data.AlertConditionInterface.ValueString())
	}
	if !data.AlertConditionJitterMs.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.jitter_ms", data.AlertConditionJitterMs.ValueInt64())
	}
	if !data.AlertConditionLatencyMs.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.latency_ms", data.AlertConditionLatencyMs.ValueInt64())
	}
	if !data.AlertConditionLossRatio.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.loss_ratio", data.AlertConditionLossRatio.ValueFloat64())
	}
	if !data.AlertConditionMos.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.mos", data.AlertConditionMos.ValueFloat64())
	}
	if !data.AlertConditionWindow.IsNull() {
		body, _ = sjson.Set(body, "alertCondition.window", data.AlertConditionWindow.ValueInt64())
	}
	if !data.RecipientsEmails.IsNull() {
		var values []string
		data.RecipientsEmails.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "recipients.emails", values)
	}
	if !data.RecipientsHttpServerIds.IsNull() {
		var values []string
		data.RecipientsHttpServerIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "recipients.httpServerIds", values)
	}
	if !data.NetworkTags.IsNull() {
		var values []string
		data.NetworkTags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "networkTags", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationAlertsProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && value.Value() != nil {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("alertCondition.bit_rate_bps"); value.Exists() && value.Value() != nil {
		data.AlertConditionBitRateBps = types.Int64Value(value.Int())
	} else {
		data.AlertConditionBitRateBps = types.Int64Null()
	}
	if value := res.Get("alertCondition.duration"); value.Exists() && value.Value() != nil {
		data.AlertConditionDuration = types.Int64Value(value.Int())
	} else {
		data.AlertConditionDuration = types.Int64Null()
	}
	if value := res.Get("alertCondition.interface"); value.Exists() && value.Value() != nil {
		data.AlertConditionInterface = types.StringValue(value.String())
	} else {
		data.AlertConditionInterface = types.StringNull()
	}
	if value := res.Get("alertCondition.jitter_ms"); value.Exists() && value.Value() != nil {
		data.AlertConditionJitterMs = types.Int64Value(value.Int())
	} else {
		data.AlertConditionJitterMs = types.Int64Null()
	}
	if value := res.Get("alertCondition.latency_ms"); value.Exists() && value.Value() != nil {
		data.AlertConditionLatencyMs = types.Int64Value(value.Int())
	} else {
		data.AlertConditionLatencyMs = types.Int64Null()
	}
	if value := res.Get("alertCondition.loss_ratio"); value.Exists() && value.Value() != nil {
		data.AlertConditionLossRatio = types.Float64Value(value.Float())
	} else {
		data.AlertConditionLossRatio = types.Float64Null()
	}
	if value := res.Get("alertCondition.mos"); value.Exists() && value.Value() != nil {
		data.AlertConditionMos = types.Float64Value(value.Float())
	} else {
		data.AlertConditionMos = types.Float64Null()
	}
	if value := res.Get("alertCondition.window"); value.Exists() && value.Value() != nil {
		data.AlertConditionWindow = types.Int64Value(value.Int())
	} else {
		data.AlertConditionWindow = types.Int64Null()
	}
	if value := res.Get("recipients.emails"); value.Exists() && value.Value() != nil {
		data.RecipientsEmails = helpers.GetStringSet(value.Array())
	} else {
		data.RecipientsEmails = types.SetNull(types.StringType)
	}
	if value := res.Get("recipients.httpServerIds"); value.Exists() && value.Value() != nil {
		data.RecipientsHttpServerIds = helpers.GetStringSet(value.Array())
	} else {
		data.RecipientsHttpServerIds = types.SetNull(types.StringType)
	}
	if value := res.Get("networkTags"); value.Exists() && value.Value() != nil {
		data.NetworkTags = helpers.GetStringSet(value.Array())
	} else {
		data.NetworkTags = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationAlertsProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("alertCondition.bit_rate_bps"); value.Exists() && !data.AlertConditionBitRateBps.IsNull() {
		data.AlertConditionBitRateBps = types.Int64Value(value.Int())
	} else {
		data.AlertConditionBitRateBps = types.Int64Null()
	}
	if value := res.Get("alertCondition.duration"); value.Exists() && !data.AlertConditionDuration.IsNull() {
		data.AlertConditionDuration = types.Int64Value(value.Int())
	} else {
		data.AlertConditionDuration = types.Int64Null()
	}
	if value := res.Get("alertCondition.interface"); value.Exists() && !data.AlertConditionInterface.IsNull() {
		data.AlertConditionInterface = types.StringValue(value.String())
	} else {
		data.AlertConditionInterface = types.StringNull()
	}
	if value := res.Get("alertCondition.jitter_ms"); value.Exists() && !data.AlertConditionJitterMs.IsNull() {
		data.AlertConditionJitterMs = types.Int64Value(value.Int())
	} else {
		data.AlertConditionJitterMs = types.Int64Null()
	}
	if value := res.Get("alertCondition.latency_ms"); value.Exists() && !data.AlertConditionLatencyMs.IsNull() {
		data.AlertConditionLatencyMs = types.Int64Value(value.Int())
	} else {
		data.AlertConditionLatencyMs = types.Int64Null()
	}
	if value := res.Get("alertCondition.loss_ratio"); value.Exists() && !data.AlertConditionLossRatio.IsNull() {
		data.AlertConditionLossRatio = types.Float64Value(value.Float())
	} else {
		data.AlertConditionLossRatio = types.Float64Null()
	}
	if value := res.Get("alertCondition.mos"); value.Exists() && !data.AlertConditionMos.IsNull() {
		data.AlertConditionMos = types.Float64Value(value.Float())
	} else {
		data.AlertConditionMos = types.Float64Null()
	}
	if value := res.Get("alertCondition.window"); value.Exists() && !data.AlertConditionWindow.IsNull() {
		data.AlertConditionWindow = types.Int64Value(value.Int())
	} else {
		data.AlertConditionWindow = types.Int64Null()
	}
	if value := res.Get("recipients.emails"); value.Exists() && !data.RecipientsEmails.IsNull() {
		data.RecipientsEmails = helpers.GetStringSet(value.Array())
	} else {
		data.RecipientsEmails = types.SetNull(types.StringType)
	}
	if value := res.Get("recipients.httpServerIds"); value.Exists() && !data.RecipientsHttpServerIds.IsNull() {
		data.RecipientsHttpServerIds = helpers.GetStringSet(value.Array())
	} else {
		data.RecipientsHttpServerIds = types.SetNull(types.StringType)
	}
	if value := res.Get("networkTags"); value.Exists() && !data.NetworkTags.IsNull() {
		data.NetworkTags = helpers.GetStringSet(value.Array())
	} else {
		data.NetworkTags = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationAlertsProfile) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationAlertsProfile) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
