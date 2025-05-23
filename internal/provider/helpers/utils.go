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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tidwall/gjson"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func GetStringListFromMapList(result []gjson.Result, key string) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		if value := result[r].Get(key); value.Exists() && value.Value() != nil {
			v[r] = types.StringValue(value.String())
		} else {
			v[r] = types.StringNull()
		}
	}
	return types.ListValueMust(types.StringType, v)
}

func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// ToLower is the same as strings.ToLower, except it cares to not to convert null/unknown strings
// into empty strings.
func ToLower(s basetypes.StringValue) basetypes.StringValue {
	if s.IsUnknown() || s.IsNull() {
		return s
	}

	return types.StringValue(strings.ToLower(s.ValueString()))
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}
