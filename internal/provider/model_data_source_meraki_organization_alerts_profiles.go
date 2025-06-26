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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationAlertsProfiles struct {
	OrganizationId types.String                                `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationAlertsProfilesItems `tfsdk:"items"`
}

type DataSourceOrganizationAlertsProfilesItems struct {
	Id                       types.String  `tfsdk:"id"`
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

func (data DataSourceOrganizationAlertsProfiles) getPath() string {
	return fmt.Sprintf("/organizations/%v/alerts/profiles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationAlertsProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationAlertsProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationAlertsProfilesItems{}
		data.Id = types.StringValue(res.Get("id").String())
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
