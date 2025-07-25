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

package helpers

import (
	"fmt"
	"strings"
)

type AttributeDescription struct {
	String string
}

func NewAttributeDescription(s string) *AttributeDescription {
	return &AttributeDescription{s}
}

func (d *AttributeDescription) AddMinimumVersionDescription(minimumVersion string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Minimum API version: `%s`", d.String, minimumVersion)
	return d
}

func (d *AttributeDescription) AddEarlyAccessDescription() *AttributeDescription {
	d.String = fmt.Sprintf("%s\n\n~>Warning: This resource or data source depends on an Early Access API endpoint. These API endpoints are subject to breaking changes without prior notice.", d.String)
	return d
}

func (d *AttributeDescription) AddBulkResourceIds(attributes ...string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n\nThis bulk resource uses the following attributes to uniquely identify each object. Changing any of these attributes will cause the object to be deleted and recreated.\n", d.String)
	for _, attr := range attributes {
		d.String = fmt.Sprintf("%s\n- `%s`", d.String, attr)
	}
	return d
}

func (d *AttributeDescription) AddDefaultValueDescription(defaultValue string) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Default value: `%s`", d.String, defaultValue)
	return d
}

func (d *AttributeDescription) AddStringEnumDescription(values ...string) *AttributeDescription {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s\n  - Choices: %s", d.String, strings.Join(v, ", "))
	return d
}

func (d *AttributeDescription) AddIntegerRangeDescription(min, max int64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}

func (d *AttributeDescription) AddFloatRangeDescription(min, max float64) *AttributeDescription {
	d.String = fmt.Sprintf("%s\n  - Range: `%v`-`%v`", d.String, min, max)
	return d
}
