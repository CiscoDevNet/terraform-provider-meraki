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
// It emits a separate model - DataSourceNetworkFirmwareUpgrades - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

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

type DataSourceNetworkFirmwareUpgrades struct {
	Id                                                     types.String `tfsdk:"id"`
	NetworkId                                              types.String `tfsdk:"network_id"`
	Timezone                                               types.String `tfsdk:"timezone"`
	ProductsApplianceParticipateInNextBetaRelease          types.Bool   `tfsdk:"products_appliance_participate_in_next_beta_release"`
	ProductsApplianceNextUpgradeTime                       types.String `tfsdk:"products_appliance_next_upgrade_time"`
	ProductsApplianceNextUpgradeToVersionId                types.String `tfsdk:"products_appliance_next_upgrade_to_version_id"`
	ProductsCameraParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_camera_participate_in_next_beta_release"`
	ProductsCameraNextUpgradeTime                          types.String `tfsdk:"products_camera_next_upgrade_time"`
	ProductsCameraNextUpgradeToVersionId                   types.String `tfsdk:"products_camera_next_upgrade_to_version_id"`
	ProductsCellularGatewayParticipateInNextBetaRelease    types.Bool   `tfsdk:"products_cellular_gateway_participate_in_next_beta_release"`
	ProductsCellularGatewayNextUpgradeTime                 types.String `tfsdk:"products_cellular_gateway_next_upgrade_time"`
	ProductsCellularGatewayNextUpgradeToVersionId          types.String `tfsdk:"products_cellular_gateway_next_upgrade_to_version_id"`
	ProductsSecureConnectParticipateInNextBetaRelease      types.Bool   `tfsdk:"products_secure_connect_participate_in_next_beta_release"`
	ProductsSecureConnectNextUpgradeTime                   types.String `tfsdk:"products_secure_connect_next_upgrade_time"`
	ProductsSecureConnectNextUpgradeToVersionId            types.String `tfsdk:"products_secure_connect_next_upgrade_to_version_id"`
	ProductsSensorParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_sensor_participate_in_next_beta_release"`
	ProductsSensorNextUpgradeTime                          types.String `tfsdk:"products_sensor_next_upgrade_time"`
	ProductsSensorNextUpgradeToVersionId                   types.String `tfsdk:"products_sensor_next_upgrade_to_version_id"`
	ProductsSwitchParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_switch_participate_in_next_beta_release"`
	ProductsSwitchNextUpgradeTime                          types.String `tfsdk:"products_switch_next_upgrade_time"`
	ProductsSwitchNextUpgradeToVersionId                   types.String `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchCatalystParticipateInNextBetaRelease     types.Bool   `tfsdk:"products_switch_catalyst_participate_in_next_beta_release"`
	ProductsSwitchCatalystNextUpgradeTime                  types.String `tfsdk:"products_switch_catalyst_next_upgrade_time"`
	ProductsSwitchCatalystNextUpgradeToVersionId           types.String `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	ProductsWirelessParticipateInNextBetaRelease           types.Bool   `tfsdk:"products_wireless_participate_in_next_beta_release"`
	ProductsWirelessNextUpgradeTime                        types.String `tfsdk:"products_wireless_next_upgrade_time"`
	ProductsWirelessNextUpgradeToVersionId                 types.String `tfsdk:"products_wireless_next_upgrade_to_version_id"`
	ProductsWirelessControllerParticipateInNextBetaRelease types.Bool   `tfsdk:"products_wireless_controller_participate_in_next_beta_release"`
	ProductsWirelessControllerNextUpgradeTime              types.String `tfsdk:"products_wireless_controller_next_upgrade_time"`
	ProductsWirelessControllerNextUpgradeToVersionId       types.String `tfsdk:"products_wireless_controller_next_upgrade_to_version_id"`
	UpgradeWindowDayOfWeek                                 types.String `tfsdk:"upgrade_window_day_of_week"`
	UpgradeWindowHourOfDay                                 types.String `tfsdk:"upgrade_window_hour_of_day"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkFirmwareUpgrades) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkFirmwareUpgrades) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("timezone"); value.Exists() && value.Value() != nil {
		data.Timezone = types.StringValue(value.String())
	} else {
		data.Timezone = types.StringNull()
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.camera.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("upgradeWindow.dayOfWeek"); value.Exists() && value.Value() != nil {
		data.UpgradeWindowDayOfWeek = types.StringValue(value.String())
	} else {
		data.UpgradeWindowDayOfWeek = types.StringNull()
	}
	if value := res.Get("upgradeWindow.hourOfDay"); value.Exists() && value.Value() != nil {
		data.UpgradeWindowHourOfDay = types.StringValue(value.String())
	} else {
		data.UpgradeWindowHourOfDay = types.StringNull()
	}
}

// End of section. //template:end fromBody
