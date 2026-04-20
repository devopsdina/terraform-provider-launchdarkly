package launchdarkly

import (
	"strings"

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

// optionalBoolFromResourceData returns d.Get(key) as bool when it is a non-nil bool value.
// When the key is missing from the schema or Get returns nil (e.g. Upjet-embedded provider), defaultVal is used.
func optionalBoolFromResourceData(d *schema.ResourceData, key string, defaultVal bool) bool {
	v := d.Get(key)
	if v == nil {
		return defaultVal
	}
	b, ok := v.(bool)
	if !ok {
		return defaultVal
	}
	return b
}

// trimmedStringAttr returns strings.TrimSpace(d.Get(key)) for string attributes; wrong or nil type yields "".
// Useful for environment keys where accidental whitespace breaks GetEnvironment(project, key).
func trimmedStringAttr(d *schema.ResourceData, key string) string {
	v := d.Get(key)
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(s)
}
