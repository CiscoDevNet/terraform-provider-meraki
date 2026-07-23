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
	_ action.Action              = &RebootDeviceAction{}
	_ action.ActionWithConfigure = &RebootDeviceAction{}
)

func NewRebootDeviceAction() action.Action {
	return &RebootDeviceAction{}
}

type RebootDeviceAction struct {
	client *meraki.Client
}

func (a *RebootDeviceAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reboot_device"
}

// Schema defines the `config` block accepted by this action. Unlike a resource/data-source
// schema, there are no Computed/Default attributes and no plan modifiers: an action has no
// state or plan to diff against, it only ever reads `Config` once during Invoke.
func (a *RebootDeviceAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This action can invoke the `Reboot Device` operation.").String,

		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial").String,
				Required:            true,
			},
		},
	}
}

func (a *RebootDeviceAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	a.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin invoke

// Invoke is a generic default: build the request body from Config, issue the request, and
// report any error. It has no way to write outputs back to Terraform (action.InvokeResponse
// carries only Diagnostics and SendProgress) and assumes the operation completes synchronously.
// Actions whose endpoint is asynchronous (returns a job id that must be polled for completion)
// need custom polling logic added by hand, outside of this marked section, since regeneration
// would otherwise overwrite it.
func (a *RebootDeviceAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ActionRebootDevice

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning Invoke")

	body := config.toBody(ctx)
	_, err := a.client.Post(config.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to invoke action, got error: %s", err))
		return
	}

	tflog.Debug(ctx, "Invoke finished successfully")
}

// End of section. //template:end invoke
