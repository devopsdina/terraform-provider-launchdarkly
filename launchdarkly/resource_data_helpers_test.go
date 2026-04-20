package launchdarkly

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestOptionalSchemaSetFromInterface(t *testing.T) {
	t.Parallel()

	require.Nil(t, optionalSchemaSetFromInterface(nil))
	require.Nil(t, optionalSchemaSetFromInterface("not-a-set"))
	require.Nil(t, optionalSchemaSetFromInterface(42))

	valid := schema.NewSet(schema.HashString, []interface{}{"a", "b"})
	got := optionalSchemaSetFromInterface(valid)
	require.NotNil(t, got)
	require.Equal(t, 2, got.Len())
}

func TestInterfaceSliceFromAny(t *testing.T) {
	t.Parallel()

	require.Nil(t, interfaceSliceFromAny(nil))
	require.Nil(t, interfaceSliceFromAny(42))
	require.Nil(t, interfaceSliceFromAny("slice"))

	sl := []interface{}{"a", "b"}
	require.Equal(t, sl, interfaceSliceFromAny(sl))
}

// Simulates embedded provider behavior: optional blocks unset → nil from d.Get (issue #387).
func TestOptionalSetListAndGetOptionalInterfaceSlice_unsetOptionalBlocks(t *testing.T) {
	t.Parallel()

	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"custom_roles": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
		},
		"policy_statements": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"effect": {Type: schema.TypeString, Required: true},
				},
			},
		},
	}, map[string]interface{}{})

	require.Empty(t, optionalSetList(d, "custom_roles"))
	require.Empty(t, getOptionalInterfaceSlice(d, "policy_statements"))
}

func TestPoliciesFromResourceData_nilPolicyNoPanic(t *testing.T) {
	t.Parallel()

	d := schema.TestResourceDataRaw(t, resourceCustomRole().Schema, map[string]interface{}{
		KEY:               "k",
		NAME:              "n",
		BASE_PERMISSIONS:  "reader",
		POLICY_STATEMENTS: []interface{}{},
	})
	// Omit POLICY — Terraform CLI often yields an empty set; helpers must tolerate nil like embedded SDK.
	require.NotPanics(t, func() {
		_ = policiesFromResourceData(d)
	})
}
