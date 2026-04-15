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

type DataSourceOrganizationWirelessDevicesProvisioningDeployment struct {
	OrganizationId types.String                                                       `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationWirelessDevicesProvisioningDeploymentItems `tfsdk:"items"`
}

type DataSourceOrganizationWirelessDevicesProvisioningDeploymentItems struct {
	Id                       types.String                                                           `tfsdk:"id"`
	MetaCountsItemsRemaining types.Int64                                                            `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                                                            `tfsdk:"meta_counts_items_total"`
	Items                    []DataSourceOrganizationWirelessDevicesProvisioningDeploymentItemsItem `tfsdk:"items"`
}

type DataSourceOrganizationWirelessDevicesProvisioningDeploymentItemsItem struct {
	CompletedAt             types.String `tfsdk:"completed_at"`
	CreatedAt               types.String `tfsdk:"created_at"`
	DeploymentId            types.String `tfsdk:"deployment_id"`
	LastUpdatedAt           types.String `tfsdk:"last_updated_at"`
	RequestedAt             types.String `tfsdk:"requested_at"`
	Status                  types.String `tfsdk:"status"`
	Type                    types.String `tfsdk:"type"`
	DevicesNewMac           types.String `tfsdk:"devices_new_mac"`
	DevicesNewModel         types.String `tfsdk:"devices_new_model"`
	DevicesNewName          types.String `tfsdk:"devices_new_name"`
	DevicesNewSerial        types.String `tfsdk:"devices_new_serial"`
	DevicesNewRfProfileId   types.String `tfsdk:"devices_new_rf_profile_id"`
	DevicesNewRfProfileName types.String `tfsdk:"devices_new_rf_profile_name"`
	DevicesNewTags          types.List   `tfsdk:"devices_new_tags"`
	DevicesOldAfterAction   types.String `tfsdk:"devices_old_after_action"`
	DevicesOldMac           types.String `tfsdk:"devices_old_mac"`
	DevicesOldModel         types.String `tfsdk:"devices_old_model"`
	DevicesOldName          types.String `tfsdk:"devices_old_name"`
	DevicesOldSerial        types.String `tfsdk:"devices_old_serial"`
	DevicesOldRfProfileId   types.String `tfsdk:"devices_old_rf_profile_id"`
	DevicesOldRfProfileName types.String `tfsdk:"devices_old_rf_profile_name"`
	DevicesOldTags          types.List   `tfsdk:"devices_old_tags"`
	NetworkId               types.String `tfsdk:"network_id"`
	NetworkName             types.String `tfsdk:"network_name"`
	Errors                  types.List   `tfsdk:"errors"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationWirelessDevicesProvisioningDeployment) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/devices/provisioning/deployments", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationWirelessDevicesProvisioningDeployment) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationWirelessDevicesProvisioningDeploymentItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationWirelessDevicesProvisioningDeploymentItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("meta.counts.items.remaining"); value.Exists() && value.Value() != nil {
			data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
		} else {
			data.MetaCountsItemsRemaining = types.Int64Null()
		}
		if value := res.Get("meta.counts.items.total"); value.Exists() && value.Value() != nil {
			data.MetaCountsItemsTotal = types.Int64Value(value.Int())
		} else {
			data.MetaCountsItemsTotal = types.Int64Null()
		}
		if value := res.Get("items"); value.Exists() && value.Value() != nil {
			data.Items = make([]DataSourceOrganizationWirelessDevicesProvisioningDeploymentItemsItem, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationWirelessDevicesProvisioningDeploymentItemsItem{}
				if value := res.Get("completedAt"); value.Exists() && value.Value() != nil {
					data.CompletedAt = types.StringValue(value.String())
				} else {
					data.CompletedAt = types.StringNull()
				}
				if value := res.Get("createdAt"); value.Exists() && value.Value() != nil {
					data.CreatedAt = types.StringValue(value.String())
				} else {
					data.CreatedAt = types.StringNull()
				}
				if value := res.Get("deploymentId"); value.Exists() && value.Value() != nil {
					data.DeploymentId = types.StringValue(value.String())
				} else {
					data.DeploymentId = types.StringNull()
				}
				if value := res.Get("lastUpdatedAt"); value.Exists() && value.Value() != nil {
					data.LastUpdatedAt = types.StringValue(value.String())
				} else {
					data.LastUpdatedAt = types.StringNull()
				}
				if value := res.Get("requestedAt"); value.Exists() && value.Value() != nil {
					data.RequestedAt = types.StringValue(value.String())
				} else {
					data.RequestedAt = types.StringNull()
				}
				if value := res.Get("status"); value.Exists() && value.Value() != nil {
					data.Status = types.StringValue(value.String())
				} else {
					data.Status = types.StringNull()
				}
				if value := res.Get("type"); value.Exists() && value.Value() != nil {
					data.Type = types.StringValue(value.String())
				} else {
					data.Type = types.StringNull()
				}
				if value := res.Get("devices.new.mac"); value.Exists() && value.Value() != nil {
					data.DevicesNewMac = types.StringValue(value.String())
				} else {
					data.DevicesNewMac = types.StringNull()
				}
				if value := res.Get("devices.new.model"); value.Exists() && value.Value() != nil {
					data.DevicesNewModel = types.StringValue(value.String())
				} else {
					data.DevicesNewModel = types.StringNull()
				}
				if value := res.Get("devices.new.name"); value.Exists() && value.Value() != nil {
					data.DevicesNewName = types.StringValue(value.String())
				} else {
					data.DevicesNewName = types.StringNull()
				}
				if value := res.Get("devices.new.serial"); value.Exists() && value.Value() != nil {
					data.DevicesNewSerial = types.StringValue(value.String())
				} else {
					data.DevicesNewSerial = types.StringNull()
				}
				if value := res.Get("devices.new.rfProfile.id"); value.Exists() && value.Value() != nil {
					data.DevicesNewRfProfileId = types.StringValue(value.String())
				} else {
					data.DevicesNewRfProfileId = types.StringNull()
				}
				if value := res.Get("devices.new.rfProfile.name"); value.Exists() && value.Value() != nil {
					data.DevicesNewRfProfileName = types.StringValue(value.String())
				} else {
					data.DevicesNewRfProfileName = types.StringNull()
				}
				if value := res.Get("devices.new.tags"); value.Exists() && value.Value() != nil {
					data.DevicesNewTags = helpers.GetStringList(value.Array())
				} else {
					data.DevicesNewTags = types.ListNull(types.StringType)
				}
				if value := res.Get("devices.old.afterAction"); value.Exists() && value.Value() != nil {
					data.DevicesOldAfterAction = types.StringValue(value.String())
				} else {
					data.DevicesOldAfterAction = types.StringNull()
				}
				if value := res.Get("devices.old.mac"); value.Exists() && value.Value() != nil {
					data.DevicesOldMac = types.StringValue(value.String())
				} else {
					data.DevicesOldMac = types.StringNull()
				}
				if value := res.Get("devices.old.model"); value.Exists() && value.Value() != nil {
					data.DevicesOldModel = types.StringValue(value.String())
				} else {
					data.DevicesOldModel = types.StringNull()
				}
				if value := res.Get("devices.old.name"); value.Exists() && value.Value() != nil {
					data.DevicesOldName = types.StringValue(value.String())
				} else {
					data.DevicesOldName = types.StringNull()
				}
				if value := res.Get("devices.old.serial"); value.Exists() && value.Value() != nil {
					data.DevicesOldSerial = types.StringValue(value.String())
				} else {
					data.DevicesOldSerial = types.StringNull()
				}
				if value := res.Get("devices.old.rfProfile.id"); value.Exists() && value.Value() != nil {
					data.DevicesOldRfProfileId = types.StringValue(value.String())
				} else {
					data.DevicesOldRfProfileId = types.StringNull()
				}
				if value := res.Get("devices.old.rfProfile.name"); value.Exists() && value.Value() != nil {
					data.DevicesOldRfProfileName = types.StringValue(value.String())
				} else {
					data.DevicesOldRfProfileName = types.StringNull()
				}
				if value := res.Get("devices.old.tags"); value.Exists() && value.Value() != nil {
					data.DevicesOldTags = helpers.GetStringList(value.Array())
				} else {
					data.DevicesOldTags = types.ListNull(types.StringType)
				}
				if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
					data.NetworkId = types.StringValue(value.String())
				} else {
					data.NetworkId = types.StringNull()
				}
				if value := res.Get("network.name"); value.Exists() && value.Value() != nil {
					data.NetworkName = types.StringValue(value.String())
				} else {
					data.NetworkName = types.StringNull()
				}
				if value := res.Get("errors"); value.Exists() && value.Value() != nil {
					data.Errors = helpers.GetStringList(value.Array())
				} else {
					data.Errors = types.ListNull(types.StringType)
				}
				(*parent).Items = append((*parent).Items, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
