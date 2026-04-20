package launchdarkly

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestCtyObjectGetAttr_missingAttributeReturnsNull(t *testing.T) {
	t.Parallel()

	config := cty.EmptyObjectVal
	v := ctyObjectGetAttr(config, INCLUDE_IN_SNIPPET)
	require.True(t, v.IsNull())

	v2 := ctyObjectGetAttr(config, DEFAULT_CLIENT_SIDE_AVAILABILITY)
	require.True(t, v2.IsNull())
}

func TestCtyObjectGetAttr_presentAttribute(t *testing.T) {
	t.Parallel()

	config := cty.ObjectVal(map[string]cty.Value{
		INCLUDE_IN_SNIPPET: cty.BoolVal(true),
	})
	v := ctyObjectGetAttr(config, INCLUDE_IN_SNIPPET)
	require.False(t, v.IsNull())
	require.True(t, v.True())
}

func TestCtyValueListElements_nullOrEmpty(t *testing.T) {
	t.Parallel()

	require.Nil(t, ctyValueListElements(cty.NullVal(cty.List(cty.String))))
	l := cty.ListVal([]cty.Value{cty.StringVal("a")})
	require.Len(t, ctyValueListElements(l), 1)
}

func TestCtyBoolTrue(t *testing.T) {
	t.Parallel()

	require.False(t, ctyBoolTrue(cty.NullVal(cty.Bool)))
	require.False(t, ctyBoolTrue(cty.StringVal("x")))
	require.True(t, ctyBoolTrue(cty.True))
	require.False(t, ctyBoolTrue(cty.False))
}
