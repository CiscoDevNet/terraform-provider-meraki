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

type WirelessEthernetPortProfiles struct {
	NetworkId types.String                        `tfsdk:"network_id"`
	Items     []WirelessEthernetPortProfilesItems `tfsdk:"items"`
}

type WirelessEthernetPortProfilesItems struct {
	Id       types.String                           `tfsdk:"id"`
	Name     types.String                           `tfsdk:"name"`
	Ports    []WirelessEthernetPortProfilesPorts    `tfsdk:"ports"`
	UsbPorts []WirelessEthernetPortProfilesUsbPorts `tfsdk:"usb_ports"`
}

type WirelessEthernetPortProfilesPorts struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Name       types.String `tfsdk:"name"`
	PskGroupId types.String `tfsdk:"psk_group_id"`
	Ssid       types.Int64  `tfsdk:"ssid"`
}

type WirelessEthernetPortProfilesUsbPorts struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	Ssid    types.Int64  `tfsdk:"ssid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessEthernetPortProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ethernet/ports/profiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessEthernetPortProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]WirelessEthernetPortProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := WirelessEthernetPortProfilesItems{}
		data.Id = types.StringValue(res.Get("profileId").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("ports"); value.Exists() && value.Value() != nil {
			data.Ports = make([]WirelessEthernetPortProfilesPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := WirelessEthernetPortProfilesPorts{}
				if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
					data.Enabled = types.BoolValue(value.Bool())
				} else {
					data.Enabled = types.BoolNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("pskGroupId"); value.Exists() && value.Value() != nil {
					data.PskGroupId = types.StringValue(value.String())
				} else {
					data.PskGroupId = types.StringNull()
				}
				if value := res.Get("ssid"); value.Exists() && value.Value() != nil {
					data.Ssid = types.Int64Value(value.Int())
				} else {
					data.Ssid = types.Int64Null()
				}
				(*parent).Ports = append((*parent).Ports, data)
				return true
			})
		}
		if value := res.Get("usbPorts"); value.Exists() && value.Value() != nil {
			data.UsbPorts = make([]WirelessEthernetPortProfilesUsbPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := WirelessEthernetPortProfilesUsbPorts{}
				if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
					data.Enabled = types.BoolValue(value.Bool())
				} else {
					data.Enabled = types.BoolNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("ssid"); value.Exists() && value.Value() != nil {
					data.Ssid = types.Int64Value(value.Int())
				} else {
					data.Ssid = types.Int64Null()
				}
				(*parent).UsbPorts = append((*parent).UsbPorts, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
