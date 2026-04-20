package launchdarkly

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// isOmittedEmbeddedSchemaAttrErr is true when the Terraform schema in use (e.g. Upjet after
// dropping deprecated attributes) does not define attr, so ResourceDiff.SetNew or ResourceData.Set
// rejects the write.
func isOmittedEmbeddedSchemaAttrErr(err error, attr string) bool {
	for err != nil {
		s := err.Error()
		// SDK messages vary by version and may be wrapped (e.g. "cannot compute the instance diff").
		if strings.Contains(s, "invalid key") && strings.Contains(s, attr) {
			return true
		}
		// MapFieldWriter.WriteField: Invalid address to set: []string{"<attr>"} (and similar)
		if strings.Contains(s, "Invalid address to set") && strings.Contains(s, attr) {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

// resourceDiffSetNewSkipMissingKey runs diff.SetNew and treats a missing schema key as success.
// Use when embedders remove deprecated attributes from the runtime schema while provider code
// still references them for Terraform CLI compatibility.
func resourceDiffSetNewSkipMissingKey(diff *schema.ResourceDiff, key string, value interface{}) error {
	err := diff.SetNew(key, value)
	if err == nil {
		return nil
	}
	if isOmittedEmbeddedSchemaAttrErr(err, key) {
		return nil
	}
	return err
}

// resourceDataSetSkipMissingKey runs d.Set and treats a missing schema key as success.
func resourceDataSetSkipMissingKey(d *schema.ResourceData, key string, value interface{}) error {
	err := d.Set(key, value)
	if err == nil {
		return nil
	}
	if isOmittedEmbeddedSchemaAttrErr(err, key) {
		return nil
	}
	return err
}
