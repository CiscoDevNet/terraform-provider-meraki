package helpers

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func IsFlagImporting(ctx context.Context, req resource.ReadRequest) (bool, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, "importing")

	return slices.Equal(v, []byte("1")), diags
}

// SetFlagImporting checks the respDiags and if they are error-free it sets the `importing` as a private flag inside
// SetKeyer. It appends its own results to respDiags.
//
// The caller must include in respDiags the result of state modification in the first place, to ensure consistency.
// The SetKeyer is something like resp.Private.
func SetFlagImporting(ctx context.Context, importing bool, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	b := []byte("0")
	if importing {
		b = []byte("1")
	}

	diags := sk.SetKey(ctx, "importing", b)
	respDiags.Append(diags...)
}

func GetJsonInitialState(ctx context.Context, req resource.DeleteRequest) (string, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, "initial_state")
	if diags.HasError() {
		return "", diags
	}

	return string(v), diags
}

// SetJsonInitialState checks the respDiags and if they are error-free it sets the `initial_state` as a private JSON value inside
// SetKeyer. It appends its own results to respDiags.
//
// The caller must include in respDiags the result of state modification in the first place, to ensure consistency.
// The SetKeyer is something like resp.Private.
func SetJsonInitialState(ctx context.Context, jsonValue string, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	diags := sk.SetKey(ctx, "initial_state", []byte(jsonValue))
	respDiags.Append(diags...)
}

func IsFlag(ctx context.Context, req resource.ReadRequest, flagName string) (bool, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, flagName)

	return slices.Equal(v, []byte("1")), diags
}

func SetFlag(ctx context.Context, flagName string, value bool, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	b := []byte("0")
	if value {
		b = []byte("1")
	}

	diags := sk.SetKey(ctx, flagName, b)
	respDiags.Append(diags...)
}

func GetJson(ctx context.Context, req resource.DeleteRequest, keyName string) (string, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, keyName)
	if diags.HasError() {
		return "", diags
	}

	return string(v), diags
}

func SetJson(ctx context.Context, keyName string, jsonValue string, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	diags := sk.SetKey(ctx, keyName, []byte(jsonValue))
	respDiags.Append(diags...)
}

// SetKeyer is something like ReadResponse.Private or ImportStateResponse.Private.
type SetKeyer interface {
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

var (
	rr resource.ReadResponse
	ir resource.ImportStateResponse

	// ensure interface match
	_ SetKeyer = rr.Private
	_ SetKeyer = ir.Private
)
