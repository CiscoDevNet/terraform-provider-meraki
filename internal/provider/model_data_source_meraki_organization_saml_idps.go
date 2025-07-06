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

type DataSourceOrganizationSAMLIdPs struct {
	OrganizationId types.String                          `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationSAMLIdPsItems `tfsdk:"items"`
}

type DataSourceOrganizationSAMLIdPsItems struct {
	Id                      types.String `tfsdk:"id"`
	SloLogoutUrl            types.String `tfsdk:"slo_logout_url"`
	X509certSha1Fingerprint types.String `tfsdk:"x509cert_sha1_fingerprint"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationSAMLIdPs) getPath() string {
	return fmt.Sprintf("/organizations/%v/saml/idps", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationSAMLIdPs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationSAMLIdPsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationSAMLIdPsItems{}
		data.Id = types.StringValue(res.Get("idpId").String())
		if value := res.Get("sloLogoutUrl"); value.Exists() && value.Value() != nil {
			data.SloLogoutUrl = types.StringValue(value.String())
		} else {
			data.SloLogoutUrl = types.StringNull()
		}
		if value := res.Get("x509certSha1Fingerprint"); value.Exists() && value.Value() != nil {
			data.X509certSha1Fingerprint = types.StringValue(value.String())
		} else {
			data.X509certSha1Fingerprint = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
