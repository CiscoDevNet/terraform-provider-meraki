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

type ResourceOrganizationWirelessDevicesProvisioningDeployments struct {
	Id             types.String                                                      `tfsdk:"id"`
	OrganizationId types.String                                                      `tfsdk:"organization_id"`
	Items          []ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems `tfsdk:"items"`
}

type ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems struct {
	Id                    types.String `tfsdk:"id"`
	Type                  types.String `tfsdk:"type"`
	Status                types.String `tfsdk:"status"`
	NetworkId             types.String `tfsdk:"network_id"`
	DevicesNewSerial      types.String `tfsdk:"devices_new_serial"`
	DevicesNewName        types.String `tfsdk:"devices_new_name"`
	DevicesNewTags        types.Set    `tfsdk:"devices_new_tags"`
	DevicesNewRfProfileId types.String `tfsdk:"devices_new_rf_profile_id"`
	DevicesOldSerial      types.String `tfsdk:"devices_old_serial"`
	DevicesOldAfterAction types.String `tfsdk:"devices_old_after_action"`
	DevicesOldName        types.String `tfsdk:"devices_old_name"`
	DevicesOldTags        types.Set    `tfsdk:"devices_old_tags"`
	DevicesOldRfProfileId types.String `tfsdk:"devices_old_rf_profile_id"`
}

type ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceOrganizationWirelessDevicesProvisioningDeployments) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/devices/provisioning/deployments", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems) toBody(ctx context.Context, state ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems) string {
	body := ""
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.Status.IsNull() {
		body, _ = sjson.Set(body, "status", data.Status.ValueString())
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "network.id", data.NetworkId.ValueString())
	}
	if !data.DevicesNewSerial.IsNull() {
		body, _ = sjson.Set(body, "devices.new.serial", data.DevicesNewSerial.ValueString())
	}
	if !data.DevicesNewName.IsNull() {
		body, _ = sjson.Set(body, "devices.new.name", data.DevicesNewName.ValueString())
	}
	if !data.DevicesNewTags.IsNull() {
		var values []string
		data.DevicesNewTags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "devices.new.tags", values)
	}
	if !data.DevicesNewRfProfileId.IsNull() {
		body, _ = sjson.Set(body, "devices.new.rfProfile.id", data.DevicesNewRfProfileId.ValueString())
	}
	if !data.DevicesOldSerial.IsNull() {
		body, _ = sjson.Set(body, "devices.old.serial", data.DevicesOldSerial.ValueString())
	}
	if !data.DevicesOldAfterAction.IsNull() {
		body, _ = sjson.Set(body, "devices.old.afterAction", data.DevicesOldAfterAction.ValueString())
	}
	if !data.DevicesOldName.IsNull() {
		body, _ = sjson.Set(body, "devices.old.name", data.DevicesOldName.ValueString())
	}
	if !data.DevicesOldTags.IsNull() {
		var values []string
		data.DevicesOldTags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "devices.old.tags", values)
	}
	if !data.DevicesOldRfProfileId.IsNull() {
		body, _ = sjson.Set(body, "devices.old.rfProfile.id", data.DevicesOldRfProfileId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("status"); value.Exists() && value.Value() != nil {
			data.Status = types.StringValue(value.String())
		} else {
			data.Status = types.StringNull()
		}
		if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("devices.new.serial"); value.Exists() && value.Value() != nil {
			data.DevicesNewSerial = types.StringValue(value.String())
		} else {
			data.DevicesNewSerial = types.StringNull()
		}
		if value := res.Get("devices.new.name"); value.Exists() && value.Value() != nil {
			data.DevicesNewName = types.StringValue(value.String())
		} else {
			data.DevicesNewName = types.StringNull()
		}
		if value := res.Get("devices.new.tags"); value.Exists() && value.Value() != nil {
			data.DevicesNewTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesNewTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.new.rfProfile.id"); value.Exists() && value.Value() != nil {
			data.DevicesNewRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesNewRfProfileId = types.StringNull()
		}
		if value := res.Get("devices.old.serial"); value.Exists() && value.Value() != nil {
			data.DevicesOldSerial = types.StringValue(value.String())
		} else {
			data.DevicesOldSerial = types.StringNull()
		}
		if value := res.Get("devices.old.afterAction"); value.Exists() && value.Value() != nil {
			data.DevicesOldAfterAction = types.StringValue(value.String())
		} else {
			data.DevicesOldAfterAction = types.StringNull()
		}
		if value := res.Get("devices.old.name"); value.Exists() && value.Value() != nil {
			data.DevicesOldName = types.StringValue(value.String())
		} else {
			data.DevicesOldName = types.StringNull()
		}
		if value := res.Get("devices.old.tags"); value.Exists() && value.Value() != nil {
			data.DevicesOldTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesOldTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.old.rfProfile.id"); value.Exists() && value.Value() != nil {
			data.DevicesOldRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesOldRfProfileId = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("deploymentId").String())
		index++
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("deploymentId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("status"); value.Exists() && !data.Status.IsNull() {
			data.Status = types.StringValue(value.String())
		} else {
			data.Status = types.StringNull()
		}
		if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("devices.new.serial"); value.Exists() && !data.DevicesNewSerial.IsNull() {
			data.DevicesNewSerial = types.StringValue(value.String())
		} else {
			data.DevicesNewSerial = types.StringNull()
		}
		if value := res.Get("devices.new.name"); value.Exists() && !data.DevicesNewName.IsNull() {
			data.DevicesNewName = types.StringValue(value.String())
		} else {
			data.DevicesNewName = types.StringNull()
		}
		if value := res.Get("devices.new.tags"); value.Exists() && !data.DevicesNewTags.IsNull() {
			data.DevicesNewTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesNewTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.new.rfProfile.id"); value.Exists() && !data.DevicesNewRfProfileId.IsNull() {
			data.DevicesNewRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesNewRfProfileId = types.StringNull()
		}
		if value := res.Get("devices.old.serial"); value.Exists() && !data.DevicesOldSerial.IsNull() {
			data.DevicesOldSerial = types.StringValue(value.String())
		} else {
			data.DevicesOldSerial = types.StringNull()
		}
		if value := res.Get("devices.old.afterAction"); value.Exists() && !data.DevicesOldAfterAction.IsNull() {
			data.DevicesOldAfterAction = types.StringValue(value.String())
		} else {
			data.DevicesOldAfterAction = types.StringNull()
		}
		if value := res.Get("devices.old.name"); value.Exists() && !data.DevicesOldName.IsNull() {
			data.DevicesOldName = types.StringValue(value.String())
		} else {
			data.DevicesOldName = types.StringNull()
		}
		if value := res.Get("devices.old.tags"); value.Exists() && !data.DevicesOldTags.IsNull() {
			data.DevicesOldTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesOldTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.old.rfProfile.id"); value.Exists() && !data.DevicesOldRfProfileId.IsNull() {
			data.DevicesOldRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesOldRfProfileId = types.StringNull()
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("deploymentId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("status"); value.Exists() && value.Value() != nil {
			data.Status = types.StringValue(value.String())
		} else {
			data.Status = types.StringNull()
		}
		if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("devices.new.serial"); value.Exists() && value.Value() != nil {
			data.DevicesNewSerial = types.StringValue(value.String())
		} else {
			data.DevicesNewSerial = types.StringNull()
		}
		if value := res.Get("devices.new.name"); value.Exists() && value.Value() != nil {
			data.DevicesNewName = types.StringValue(value.String())
		} else {
			data.DevicesNewName = types.StringNull()
		}
		if value := res.Get("devices.new.tags"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.DevicesNewTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesNewTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.new.rfProfile.id"); value.Exists() && value.Value() != nil {
			data.DevicesNewRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesNewRfProfileId = types.StringNull()
		}
		if value := res.Get("devices.old.serial"); value.Exists() && value.Value() != nil {
			data.DevicesOldSerial = types.StringValue(value.String())
		} else {
			data.DevicesOldSerial = types.StringNull()
		}
		if value := res.Get("devices.old.afterAction"); value.Exists() && value.Value() != nil {
			data.DevicesOldAfterAction = types.StringValue(value.String())
		} else {
			data.DevicesOldAfterAction = types.StringNull()
		}
		if value := res.Get("devices.old.name"); value.Exists() && value.Value() != nil {
			data.DevicesOldName = types.StringValue(value.String())
		} else {
			data.DevicesOldName = types.StringNull()
		}
		if value := res.Get("devices.old.tags"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.DevicesOldTags = helpers.GetStringSet(value.Array())
		} else {
			data.DevicesOldTags = types.SetNull(types.StringType)
		}
		if value := res.Get("devices.old.rfProfile.id"); value.Exists() && value.Value() != nil {
			data.DevicesOldRfProfileId = types.StringValue(value.String())
		} else {
			data.DevicesOldRfProfileId = types.StringNull()
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity) toIdentity(ctx context.Context, plan *ResourceOrganizationWirelessDevicesProvisioningDeployments) {
	data.OrganizationId = plan.OrganizationId
	if len(data.ItemIds.Elements()) == 0 {
		data.ItemIds = types.ListNull(types.StringType)
	}
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) fromIdentity(ctx context.Context, identity *ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity) {
	data.OrganizationId = identity.OrganizationId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceOrganizationWirelessDevicesProvisioningDeployments) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceOrganizationWirelessDevicesProvisioningDeployments) hasChanges(ctx context.Context, state *ResourceOrganizationWirelessDevicesProvisioningDeployments, id string) bool {
	hasChanges := false

	item := ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Type.Equal(stateItem.Type) {
		hasChanges = true
	}
	if !item.Status.Equal(stateItem.Status) {
		hasChanges = true
	}
	if !item.NetworkId.Equal(stateItem.NetworkId) {
		hasChanges = true
	}
	if !item.DevicesNewSerial.Equal(stateItem.DevicesNewSerial) {
		hasChanges = true
	}
	if !item.DevicesNewName.Equal(stateItem.DevicesNewName) {
		hasChanges = true
	}
	if !item.DevicesNewTags.Equal(stateItem.DevicesNewTags) {
		hasChanges = true
	}
	if !item.DevicesNewRfProfileId.Equal(stateItem.DevicesNewRfProfileId) {
		hasChanges = true
	}
	if !item.DevicesOldSerial.Equal(stateItem.DevicesOldSerial) {
		hasChanges = true
	}
	if !item.DevicesOldAfterAction.Equal(stateItem.DevicesOldAfterAction) {
		hasChanges = true
	}
	if !item.DevicesOldName.Equal(stateItem.DevicesOldName) {
		hasChanges = true
	}
	if !item.DevicesOldTags.Equal(stateItem.DevicesOldTags) {
		hasChanges = true
	}
	if !item.DevicesOldRfProfileId.Equal(stateItem.DevicesOldRfProfileId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
