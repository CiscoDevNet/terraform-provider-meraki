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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

// ActionRebootDevice models the `config` block of an action invocation. Unlike a
// resource/data-source model, this is never persisted: it exists only for the duration of
// a single Invoke call, so it has no computed/read-only fields and no identity.
type ActionRebootDevice struct {
	Serial types.String `tfsdk:"serial"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ActionRebootDevice) getPath() string {
	return fmt.Sprintf("/devices/%v/reboot", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ActionRebootDevice) toBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toBody
