package launchdarkly

import (
	"github.com/hashicorp/go-cty/cty"
)

func ctyObjectGetAttr(config cty.Value, attr string) cty.Value {
	if config.IsNull() || !config.IsKnown() {
		return cty.NullVal(cty.DynamicPseudoType)
	}
	ty := config.Type()
	if !ty.IsObjectType() || !ty.HasAttribute(attr) {
		return cty.NullVal(cty.DynamicPseudoType)
	}
	return config.GetAttr(attr)
}

func ctyValueListElements(v cty.Value) []cty.Value {
	if v.IsNull() || !v.IsKnown() {
		return nil
	}
	return v.AsValueSlice()
}

func ctyBoolTrue(v cty.Value) bool {
	if v.IsNull() || !v.IsKnown() || v.Type() != cty.Bool {
		return false
	}
	return v.True()
}
