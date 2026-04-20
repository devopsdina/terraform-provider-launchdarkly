package launchdarkly

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getOptionalSet(d *schema.ResourceData, key string) *schema.Set {
	return optionalSchemaSetFromInterface(d.Get(key))
}

func optionalSetList(d *schema.ResourceData, key string) []interface{} {
	s := getOptionalSet(d, key)
	if s == nil {
		return nil
	}
	return s.List()
}

func getOptionalInterfaceSlice(d *schema.ResourceData, key string) []interface{} {
	return interfaceSliceFromAny(d.Get(key))
}

func optionalSchemaSetFromInterface(v interface{}) *schema.Set {
	if v == nil {
		return nil
	}
	s, ok := v.(*schema.Set)
	if !ok || s == nil {
		return nil
	}
	return s
}

func interfaceSliceFromAny(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	s, ok := v.([]interface{})
	if !ok {
		return nil
	}
	return s
}

// stringListFromOptionalSetValue converts a *schema.Set wrapped in interface{} (e.g. diff.Get / GetChange)
// to a []string for LaunchDarkly API calls. Nil or wrong type yields nil.
func stringListFromOptionalSetValue(v interface{}) []string {
	s := optionalSchemaSetFromInterface(v)
	if s == nil {
		return nil
	}
	return interfaceSliceToStringSlice(s.List())
}

func optionalSetListFromAny(v interface{}) []interface{} {
	s := optionalSchemaSetFromInterface(v)
	if s == nil {
		return nil
	}
	return s.List()
}
