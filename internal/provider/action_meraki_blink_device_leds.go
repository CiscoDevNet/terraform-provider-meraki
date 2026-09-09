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
	"time"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ action.Action              = &BlinkDeviceLedsAction{}
	_ action.ActionWithConfigure = &BlinkDeviceLedsAction{}
)

func NewBlinkDeviceLedsAction() action.Action {
	return &BlinkDeviceLedsAction{}
}

type BlinkDeviceLedsAction struct {
	client *meraki.Client
}

func (a *BlinkDeviceLedsAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_blink_device_leds"
}

// Schema defines the `config` block accepted by this action. Unlike a resource/data-source
// schema, there are no Computed/Default attributes and no plan modifiers: an action has no
// state or plan to diff against, it only ever reads `Config` once during Invoke.
func (a *BlinkDeviceLedsAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This action can invoke the `Blink Device Leds` operation.").String,

		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial").String,
				Required:            true,
			},
			"duration": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The duration in seconds to blink LEDs.").String,
				Required:            true,
			},
		},
	}
}

func (a *BlinkDeviceLedsAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	a.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Invoke is hand-maintained (the //template:begin/end markers were deliberately removed so
// "make gen" never regenerates this section): this action's endpoint (createDeviceLiveToolsLedsBlink)
// is asynchronous, it only enqueues a job, so a call to pollBlinkLedsJob (defined below) was added
// on top of the generic default to wait for the job to actually finish. If the action's attributes
// ever change, this function needs to be updated by hand to match.
func (a *BlinkDeviceLedsAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ActionBlinkDeviceLeds

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning Invoke")

	body := config.toBody(ctx)
	res, err := a.client.Post(config.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to invoke action, got error: %s", err))
		return
	}

	a.pollBlinkLedsJob(ctx, config.Serial.ValueString(), res.Get("ledsBlinkId").String(), resp)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Invoke finished successfully")
}

// pollBlinkLedsJob polls GET /devices/{serial}/liveTools/leds/blink/{ledsBlinkId} until the job
// reaches a terminal status ("complete" or "failed") or the timeout elapses. Like Invoke above,
// this function is hand-maintained and not subject to template regeneration.
func (a *BlinkDeviceLedsAction) pollBlinkLedsJob(ctx context.Context, serial, ledsBlinkId string, resp *action.InvokeResponse) {
	if ledsBlinkId == "" {
		resp.Diagnostics.AddError("Client Error", "Blink LEDs job did not return a ledsBlinkId to poll")
		return
	}

	const pollInterval = 2 * time.Second
	const pollTimeout = 60 * time.Second

	deadline := time.Now().Add(pollTimeout)
	for {
		res, err := a.client.Get(fmt.Sprintf("/devices/%s/liveTools/leds/blink/%s", serial, ledsBlinkId))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to poll blink LEDs job, got error: %s", err))
			return
		}

		status := res.Get("status").String()
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Blink LEDs job %s status: %s", ledsBlinkId, status)})

		switch status {
		case "complete":
			return
		case "failed":
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Blink LEDs job %s failed: %s", ledsBlinkId, res.Get("error").String()))
			return
		}

		if time.Now().After(deadline) {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Timed out waiting for blink LEDs job %s to complete (last status: %s)", ledsBlinkId, status))
			return
		}

		time.Sleep(pollInterval)
	}
}
