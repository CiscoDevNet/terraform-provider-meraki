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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessEthernetPortProfile struct {
	Id        types.String                          `tfsdk:"id"`
	NetworkId types.String                          `tfsdk:"network_id"`
	Name      types.String                          `tfsdk:"name"`
	Ports     []WirelessEthernetPortProfilePorts    `tfsdk:"ports"`
	UsbPorts  []WirelessEthernetPortProfileUsbPorts `tfsdk:"usb_ports"`
}

type WirelessEthernetPortProfilePorts struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Name       types.String `tfsdk:"name"`
	PskGroupId types.String `tfsdk:"psk_group_id"`
	Ssid       types.Int64  `tfsdk:"ssid"`
}

type WirelessEthernetPortProfileUsbPorts struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	Ssid    types.Int64  `tfsdk:"ssid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessEthernetPortProfile) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ethernet/ports/profiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessEthernetPortProfile) toBody(ctx context.Context, state WirelessEthernetPortProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	{
		body, _ = sjson.Set(body, "ports", []interface{}{})
		for _, item := range data.Ports {
			itemBody := ""
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.PskGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "pskGroupId", item.PskGroupId.ValueString())
			}
			if !item.Ssid.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ssid", item.Ssid.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "ports.-1", itemBody)
		}
	}
	if len(data.UsbPorts) > 0 {
		body, _ = sjson.Set(body, "usbPorts", []interface{}{})
		for _, item := range data.UsbPorts {
			itemBody := ""
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Ssid.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ssid", item.Ssid.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "usbPorts.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessEthernetPortProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ports"); value.Exists() && value.Value() != nil {
		data.Ports = make([]WirelessEthernetPortProfilePorts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessEthernetPortProfilePorts{}
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
		data.UsbPorts = make([]WirelessEthernetPortProfileUsbPorts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessEthernetPortProfileUsbPorts{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessEthernetPortProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := 0; i < len(data.Ports); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Ports[i].Name.ValueString()}

		parent := &data
		data := (*parent).Ports[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ports").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ports[%d] = %+v",
				i,
				(*parent).Ports[i],
			))
			(*parent).Ports = slices.Delete((*parent).Ports, i, i+1)
			i--

			continue
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("pskGroupId"); value.Exists() && !data.PskGroupId.IsNull() {
			data.PskGroupId = types.StringValue(value.String())
		} else {
			data.PskGroupId = types.StringNull()
		}
		if value := res.Get("ssid"); value.Exists() && !data.Ssid.IsNull() {
			data.Ssid = types.Int64Value(value.Int())
		} else {
			data.Ssid = types.Int64Null()
		}
		(*parent).Ports[i] = data
	}
	for i := 0; i < len(data.UsbPorts); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.UsbPorts[i].Name.ValueString()}

		parent := &data
		data := (*parent).UsbPorts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("usbPorts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing UsbPorts[%d] = %+v",
				i,
				(*parent).UsbPorts[i],
			))
			(*parent).UsbPorts = slices.Delete((*parent).UsbPorts, i, i+1)
			i--

			continue
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("ssid"); value.Exists() && !data.Ssid.IsNull() {
			data.Ssid = types.Int64Value(value.Int())
		} else {
			data.Ssid = types.Int64Null()
		}
		(*parent).UsbPorts[i] = data
	}
}

// End of section. //template:end fromBodyPartial
