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

type OrganizationFirmwareUpgrades struct {
	OrganizationId types.String                        `tfsdk:"organization_id"`
	Items          []OrganizationFirmwareUpgradesItems `tfsdk:"items"`
}

type OrganizationFirmwareUpgradesItems struct {
	Id                     types.String `tfsdk:"id"`
	CompletedAt            types.String `tfsdk:"completed_at"`
	ProductTypes           types.String `tfsdk:"product_types"`
	Status                 types.String `tfsdk:"status"`
	Time                   types.String `tfsdk:"time"`
	UpgradeBatchId         types.String `tfsdk:"upgrade_batch_id"`
	UpgradeId              types.String `tfsdk:"upgrade_id"`
	FromVersionFirmware    types.String `tfsdk:"from_version_firmware"`
	FromVersionId          types.String `tfsdk:"from_version_id"`
	FromVersionReleaseDate types.String `tfsdk:"from_version_release_date"`
	FromVersionReleaseType types.String `tfsdk:"from_version_release_type"`
	FromVersionShortName   types.String `tfsdk:"from_version_short_name"`
	NetworkId              types.String `tfsdk:"network_id"`
	NetworkName            types.String `tfsdk:"network_name"`
	ToVersionFirmware      types.String `tfsdk:"to_version_firmware"`
	ToVersionId            types.String `tfsdk:"to_version_id"`
	ToVersionReleaseDate   types.String `tfsdk:"to_version_release_date"`
	ToVersionReleaseType   types.String `tfsdk:"to_version_release_type"`
	ToVersionShortName     types.String `tfsdk:"to_version_short_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationFirmwareUpgrades) getPath() string {
	return fmt.Sprintf("/organizations/%v/firmware/upgrades", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationFirmwareUpgrades) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationFirmwareUpgradesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationFirmwareUpgradesItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("completedAt"); value.Exists() && value.Value() != nil {
			data.CompletedAt = types.StringValue(value.String())
		} else {
			data.CompletedAt = types.StringNull()
		}
		if value := res.Get("productTypes"); value.Exists() && value.Value() != nil {
			data.ProductTypes = types.StringValue(value.String())
		} else {
			data.ProductTypes = types.StringNull()
		}
		if value := res.Get("status"); value.Exists() && value.Value() != nil {
			data.Status = types.StringValue(value.String())
		} else {
			data.Status = types.StringNull()
		}
		if value := res.Get("time"); value.Exists() && value.Value() != nil {
			data.Time = types.StringValue(value.String())
		} else {
			data.Time = types.StringNull()
		}
		if value := res.Get("upgradeBatchId"); value.Exists() && value.Value() != nil {
			data.UpgradeBatchId = types.StringValue(value.String())
		} else {
			data.UpgradeBatchId = types.StringNull()
		}
		if value := res.Get("upgradeId"); value.Exists() && value.Value() != nil {
			data.UpgradeId = types.StringValue(value.String())
		} else {
			data.UpgradeId = types.StringNull()
		}
		if value := res.Get("fromVersion.firmware"); value.Exists() && value.Value() != nil {
			data.FromVersionFirmware = types.StringValue(value.String())
		} else {
			data.FromVersionFirmware = types.StringNull()
		}
		if value := res.Get("fromVersion.id"); value.Exists() && value.Value() != nil {
			data.FromVersionId = types.StringValue(value.String())
		} else {
			data.FromVersionId = types.StringNull()
		}
		if value := res.Get("fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
			data.FromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.FromVersionReleaseDate = types.StringNull()
		}
		if value := res.Get("fromVersion.releaseType"); value.Exists() && value.Value() != nil {
			data.FromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.FromVersionReleaseType = types.StringNull()
		}
		if value := res.Get("fromVersion.shortName"); value.Exists() && value.Value() != nil {
			data.FromVersionShortName = types.StringValue(value.String())
		} else {
			data.FromVersionShortName = types.StringNull()
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
		if value := res.Get("toVersion.firmware"); value.Exists() && value.Value() != nil {
			data.ToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ToVersionFirmware = types.StringNull()
		}
		if value := res.Get("toVersion.id"); value.Exists() && value.Value() != nil {
			data.ToVersionId = types.StringValue(value.String())
		} else {
			data.ToVersionId = types.StringNull()
		}
		if value := res.Get("toVersion.releaseDate"); value.Exists() && value.Value() != nil {
			data.ToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ToVersionReleaseDate = types.StringNull()
		}
		if value := res.Get("toVersion.releaseType"); value.Exists() && value.Value() != nil {
			data.ToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ToVersionReleaseType = types.StringNull()
		}
		if value := res.Get("toVersion.shortName"); value.Exists() && value.Value() != nil {
			data.ToVersionShortName = types.StringValue(value.String())
		} else {
			data.ToVersionShortName = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
