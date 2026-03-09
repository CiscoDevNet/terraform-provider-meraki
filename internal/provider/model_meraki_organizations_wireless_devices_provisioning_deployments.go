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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationsWirelessDevicesProvisioningDeployments struct {
	Id                       types.String                                               `tfsdk:"id"`
	OrganizationId           types.String                                               `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                                                `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                                                `tfsdk:"meta_counts_items_total"`
	Items                    []OrganizationsWirelessDevicesProvisioningDeploymentsItems `tfsdk:"items"`
}

type OrganizationsWirelessDevicesProvisioningDeploymentsItems struct {
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

func (data OrganizationsWirelessDevicesProvisioningDeployments) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/devices/provisioning/deployments", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationsWirelessDevicesProvisioningDeployments) toBody(ctx context.Context, state OrganizationsWirelessDevicesProvisioningDeployments) string {
	body := ""
	if !data.MetaCountsItemsRemaining.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.remaining", data.MetaCountsItemsRemaining.ValueInt64())
	}
	if !data.MetaCountsItemsTotal.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.total", data.MetaCountsItemsTotal.ValueInt64())
	}
	{
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.CompletedAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "completedAt", item.CompletedAt.ValueString())
			}
			if !item.CreatedAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "createdAt", item.CreatedAt.ValueString())
			}
			if !item.DeploymentId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deploymentId", item.DeploymentId.ValueString())
			}
			if !item.LastUpdatedAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lastUpdatedAt", item.LastUpdatedAt.ValueString())
			}
			if !item.RequestedAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "requestedAt", item.RequestedAt.ValueString())
			}
			if !item.Status.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "status", item.Status.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.DevicesNewMac.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.mac", item.DevicesNewMac.ValueString())
			}
			if !item.DevicesNewModel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.model", item.DevicesNewModel.ValueString())
			}
			if !item.DevicesNewName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.name", item.DevicesNewName.ValueString())
			}
			if !item.DevicesNewSerial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.serial", item.DevicesNewSerial.ValueString())
			}
			if !item.DevicesNewRfProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.rfProfile.id", item.DevicesNewRfProfileId.ValueString())
			}
			if !item.DevicesNewRfProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.new.rfProfile.name", item.DevicesNewRfProfileName.ValueString())
			}
			if !item.DevicesNewTags.IsNull() {
				var values []string
				item.DevicesNewTags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "devices.new.tags", values)
			}
			if !item.DevicesOldAfterAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.afterAction", item.DevicesOldAfterAction.ValueString())
			}
			if !item.DevicesOldMac.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.mac", item.DevicesOldMac.ValueString())
			}
			if !item.DevicesOldModel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.model", item.DevicesOldModel.ValueString())
			}
			if !item.DevicesOldName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.name", item.DevicesOldName.ValueString())
			}
			if !item.DevicesOldSerial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.serial", item.DevicesOldSerial.ValueString())
			}
			if !item.DevicesOldRfProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.rfProfile.id", item.DevicesOldRfProfileId.ValueString())
			}
			if !item.DevicesOldRfProfileName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devices.old.rfProfile.name", item.DevicesOldRfProfileName.ValueString())
			}
			if !item.DevicesOldTags.IsNull() {
				var values []string
				item.DevicesOldTags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "devices.old.tags", values)
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "network.id", item.NetworkId.ValueString())
			}
			if !item.NetworkName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "network.name", item.NetworkName.ValueString())
			}
			if !item.Errors.IsNull() {
				var values []string
				item.Errors.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "errors", values)
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationsWirelessDevicesProvisioningDeployments) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Items = make([]OrganizationsWirelessDevicesProvisioningDeploymentsItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationsWirelessDevicesProvisioningDeploymentsItems{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationsWirelessDevicesProvisioningDeployments) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && !data.MetaCountsItemsRemaining.IsNull() {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && !data.MetaCountsItemsTotal.IsNull() {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"completedAt", "createdAt", "deploymentId", "lastUpdatedAt", "requestedAt", "status", "type", "devices.new.mac", "devices.new.model", "devices.new.name", "devices.new.serial", "devices.new.rfProfile.id", "devices.new.rfProfile.name", "devices.old.afterAction", "devices.old.mac", "devices.old.model", "devices.old.name", "devices.old.serial", "devices.old.rfProfile.id", "devices.old.rfProfile.name", "network.id", "network.name"}
		keyValues := [...]string{data.Items[i].CompletedAt.ValueString(), data.Items[i].CreatedAt.ValueString(), data.Items[i].DeploymentId.ValueString(), data.Items[i].LastUpdatedAt.ValueString(), data.Items[i].RequestedAt.ValueString(), data.Items[i].Status.ValueString(), data.Items[i].Type.ValueString(), data.Items[i].DevicesNewMac.ValueString(), data.Items[i].DevicesNewModel.ValueString(), data.Items[i].DevicesNewName.ValueString(), data.Items[i].DevicesNewSerial.ValueString(), data.Items[i].DevicesNewRfProfileId.ValueString(), data.Items[i].DevicesNewRfProfileName.ValueString(), data.Items[i].DevicesOldAfterAction.ValueString(), data.Items[i].DevicesOldMac.ValueString(), data.Items[i].DevicesOldModel.ValueString(), data.Items[i].DevicesOldName.ValueString(), data.Items[i].DevicesOldSerial.ValueString(), data.Items[i].DevicesOldRfProfileId.ValueString(), data.Items[i].DevicesOldRfProfileName.ValueString(), data.Items[i].NetworkId.ValueString(), data.Items[i].NetworkName.ValueString()}

		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Items[%d] = %+v",
				i,
				(*parent).Items[i],
			))
			(*parent).Items = slices.Delete((*parent).Items, i, i+1)
			i--

			continue
		}
		if value := res.Get("completedAt"); value.Exists() && !data.CompletedAt.IsNull() {
			data.CompletedAt = types.StringValue(value.String())
		} else {
			data.CompletedAt = types.StringNull()
		}
		if value := res.Get("createdAt"); value.Exists() && !data.CreatedAt.IsNull() {
			data.CreatedAt = types.StringValue(value.String())
		} else {
			data.CreatedAt = types.StringNull()
		}
		if value := res.Get("deploymentId"); value.Exists() && !data.DeploymentId.IsNull() {
			data.DeploymentId = types.StringValue(value.String())
		} else {
			data.DeploymentId = types.StringNull()
		}
		if value := res.Get("lastUpdatedAt"); value.Exists() && !data.LastUpdatedAt.IsNull() {
			data.LastUpdatedAt = types.StringValue(value.String())
		} else {
			data.LastUpdatedAt = types.StringNull()
		}
		if value := res.Get("requestedAt"); value.Exists() && !data.RequestedAt.IsNull() {
			data.RequestedAt = types.StringValue(value.String())
		} else {
			data.RequestedAt = types.StringNull()
		}
		if value := res.Get("status"); value.Exists() && !data.Status.IsNull() {
			data.Status = types.StringValue(value.String())
		} else {
			data.Status = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("devices.new.mac"); value.Exists() && !data.DevicesNewMac.IsNull() {
			data.DevicesNewMac = types.StringValue(value.String())
		} else {
			data.DevicesNewMac = types.StringNull()
		}
		if value := res.Get("devices.new.model"); value.Exists() && !data.DevicesNewModel.IsNull() {
			data.DevicesNewModel = types.StringValue(value.String())
		} else {
			data.DevicesNewModel = types.StringNull()
		}
		if value := res.Get("devices.new.name"); value.Exists() && !data.DevicesNewName.IsNull() {
			data.DevicesNewName = types.StringValue(value.String())
		} else {
			data.DevicesNewName = types.StringNull()
		}
		if value := res.Get("devices.new.serial"); value.Exists() && !data.DevicesNewSerial.IsNull() {
			data.DevicesNewSerial = types.StringValue(value.String())
		} else {
			data.DevicesNewSerial = types.StringNull()
		}
		if value := res.Get("devices.new.rfProfile.id"); value.Exists() && !data.DevicesNewRfProfileId.IsNull() {
			data.DevicesNewRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesNewRfProfileId = types.StringNull()
		}
		if value := res.Get("devices.new.rfProfile.name"); value.Exists() && !data.DevicesNewRfProfileName.IsNull() {
			data.DevicesNewRfProfileName = types.StringValue(value.String())
		} else {
			data.DevicesNewRfProfileName = types.StringNull()
		}
		if value := res.Get("devices.new.tags"); value.Exists() && !data.DevicesNewTags.IsNull() {
			data.DevicesNewTags = helpers.GetStringList(value.Array())
		} else {
			data.DevicesNewTags = types.ListNull(types.StringType)
		}
		if value := res.Get("devices.old.afterAction"); value.Exists() && !data.DevicesOldAfterAction.IsNull() {
			data.DevicesOldAfterAction = types.StringValue(value.String())
		} else {
			data.DevicesOldAfterAction = types.StringNull()
		}
		if value := res.Get("devices.old.mac"); value.Exists() && !data.DevicesOldMac.IsNull() {
			data.DevicesOldMac = types.StringValue(value.String())
		} else {
			data.DevicesOldMac = types.StringNull()
		}
		if value := res.Get("devices.old.model"); value.Exists() && !data.DevicesOldModel.IsNull() {
			data.DevicesOldModel = types.StringValue(value.String())
		} else {
			data.DevicesOldModel = types.StringNull()
		}
		if value := res.Get("devices.old.name"); value.Exists() && !data.DevicesOldName.IsNull() {
			data.DevicesOldName = types.StringValue(value.String())
		} else {
			data.DevicesOldName = types.StringNull()
		}
		if value := res.Get("devices.old.serial"); value.Exists() && !data.DevicesOldSerial.IsNull() {
			data.DevicesOldSerial = types.StringValue(value.String())
		} else {
			data.DevicesOldSerial = types.StringNull()
		}
		if value := res.Get("devices.old.rfProfile.id"); value.Exists() && !data.DevicesOldRfProfileId.IsNull() {
			data.DevicesOldRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesOldRfProfileId = types.StringNull()
		}
		if value := res.Get("devices.old.rfProfile.name"); value.Exists() && !data.DevicesOldRfProfileName.IsNull() {
			data.DevicesOldRfProfileName = types.StringValue(value.String())
		} else {
			data.DevicesOldRfProfileName = types.StringNull()
		}
		if value := res.Get("devices.old.tags"); value.Exists() && !data.DevicesOldTags.IsNull() {
			data.DevicesOldTags = helpers.GetStringList(value.Array())
		} else {
			data.DevicesOldTags = types.ListNull(types.StringType)
		}
		if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("network.name"); value.Exists() && !data.NetworkName.IsNull() {
			data.NetworkName = types.StringValue(value.String())
		} else {
			data.NetworkName = types.StringNull()
		}
		if value := res.Get("errors"); value.Exists() && !data.Errors.IsNull() {
			data.Errors = helpers.GetStringList(value.Array())
		} else {
			data.Errors = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationsWirelessDevicesProvisioningDeployments) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationsWirelessDevicesProvisioningDeployments) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
