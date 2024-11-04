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

type OrganizationLicenses struct {
	OrganizationId types.String                `tfsdk:"organization_id"`
	Items          []OrganizationLicensesItems `tfsdk:"items"`
}

type OrganizationLicensesItems struct {
	Id                        types.String                                    `tfsdk:"id"`
	ActivationDate            types.String                                    `tfsdk:"activation_date"`
	ClaimDate                 types.String                                    `tfsdk:"claim_date"`
	DeviceSerial              types.String                                    `tfsdk:"device_serial"`
	DurationInDays            types.Int64                                     `tfsdk:"duration_in_days"`
	ExpirationDate            types.String                                    `tfsdk:"expiration_date"`
	HeadLicenseId             types.String                                    `tfsdk:"head_license_id"`
	LicenseKey                types.String                                    `tfsdk:"license_key"`
	LicenseType               types.String                                    `tfsdk:"license_type"`
	NetworkId                 types.String                                    `tfsdk:"network_id"`
	OrderNumber               types.String                                    `tfsdk:"order_number"`
	SeatCount                 types.Int64                                     `tfsdk:"seat_count"`
	State                     types.String                                    `tfsdk:"state"`
	TotalDurationInDays       types.Int64                                     `tfsdk:"total_duration_in_days"`
	PermanentlyQueuedLicenses []OrganizationLicensesPermanentlyQueuedLicenses `tfsdk:"permanently_queued_licenses"`
}

type OrganizationLicensesPermanentlyQueuedLicenses struct {
	DurationInDays types.Int64  `tfsdk:"duration_in_days"`
	Id             types.String `tfsdk:"id"`
	LicenseKey     types.String `tfsdk:"license_key"`
	LicenseType    types.String `tfsdk:"license_type"`
	OrderNumber    types.String `tfsdk:"order_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationLicenses) getPath() string {
	return fmt.Sprintf("/organizations/%v/licenses", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationLicenses) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationLicensesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationLicensesItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("activationDate"); value.Exists() && value.Value() != nil {
			data.ActivationDate = types.StringValue(value.String())
		} else {
			data.ActivationDate = types.StringNull()
		}
		if value := res.Get("claimDate"); value.Exists() && value.Value() != nil {
			data.ClaimDate = types.StringValue(value.String())
		} else {
			data.ClaimDate = types.StringNull()
		}
		if value := res.Get("deviceSerial"); value.Exists() && value.Value() != nil {
			data.DeviceSerial = types.StringValue(value.String())
		} else {
			data.DeviceSerial = types.StringNull()
		}
		if value := res.Get("durationInDays"); value.Exists() && value.Value() != nil {
			data.DurationInDays = types.Int64Value(value.Int())
		} else {
			data.DurationInDays = types.Int64Null()
		}
		if value := res.Get("expirationDate"); value.Exists() && value.Value() != nil {
			data.ExpirationDate = types.StringValue(value.String())
		} else {
			data.ExpirationDate = types.StringNull()
		}
		if value := res.Get("headLicenseId"); value.Exists() && value.Value() != nil {
			data.HeadLicenseId = types.StringValue(value.String())
		} else {
			data.HeadLicenseId = types.StringNull()
		}
		if value := res.Get("licenseKey"); value.Exists() && value.Value() != nil {
			data.LicenseKey = types.StringValue(value.String())
		} else {
			data.LicenseKey = types.StringNull()
		}
		if value := res.Get("licenseType"); value.Exists() && value.Value() != nil {
			data.LicenseType = types.StringValue(value.String())
		} else {
			data.LicenseType = types.StringNull()
		}
		if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("orderNumber"); value.Exists() && value.Value() != nil {
			data.OrderNumber = types.StringValue(value.String())
		} else {
			data.OrderNumber = types.StringNull()
		}
		if value := res.Get("seatCount"); value.Exists() && value.Value() != nil {
			data.SeatCount = types.Int64Value(value.Int())
		} else {
			data.SeatCount = types.Int64Null()
		}
		if value := res.Get("state"); value.Exists() && value.Value() != nil {
			data.State = types.StringValue(value.String())
		} else {
			data.State = types.StringNull()
		}
		if value := res.Get("totalDurationInDays"); value.Exists() && value.Value() != nil {
			data.TotalDurationInDays = types.Int64Value(value.Int())
		} else {
			data.TotalDurationInDays = types.Int64Null()
		}
		if value := res.Get("permanentlyQueuedLicenses"); value.Exists() && value.Value() != nil {
			data.PermanentlyQueuedLicenses = make([]OrganizationLicensesPermanentlyQueuedLicenses, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := OrganizationLicensesPermanentlyQueuedLicenses{}
				if value := res.Get("durationInDays"); value.Exists() && value.Value() != nil {
					data.DurationInDays = types.Int64Value(value.Int())
				} else {
					data.DurationInDays = types.Int64Null()
				}
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				if value := res.Get("licenseKey"); value.Exists() && value.Value() != nil {
					data.LicenseKey = types.StringValue(value.String())
				} else {
					data.LicenseKey = types.StringNull()
				}
				if value := res.Get("licenseType"); value.Exists() && value.Value() != nil {
					data.LicenseType = types.StringValue(value.String())
				} else {
					data.LicenseType = types.StringNull()
				}
				if value := res.Get("orderNumber"); value.Exists() && value.Value() != nil {
					data.OrderNumber = types.StringValue(value.String())
				} else {
					data.OrderNumber = types.StringNull()
				}
				(*parent).PermanentlyQueuedLicenses = append((*parent).PermanentlyQueuedLicenses, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
