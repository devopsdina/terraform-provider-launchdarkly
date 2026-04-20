package launchdarkly

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsOmittedEmbeddedSchemaAttrErr(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		attr   string
		wantOK bool
	}{
		{
			name:   "SetNew invalid key",
			err:    fmt.Errorf("SetNew: invalid key: include_in_snippet"),
			attr:   "include_in_snippet",
			wantOK: true,
		},
		{
			name:   "wrong attr",
			err:    fmt.Errorf("SetNew: invalid key: other"),
			attr:   "include_in_snippet",
			wantOK: false,
		},
		{
			name:   "Invalid address",
			err:    errors.New(`Invalid address to set: []string{"include_in_snippet"}`),
			attr:   "include_in_snippet",
			wantOK: true,
		},
		{
			name:   "nil",
			err:    nil,
			attr:   "include_in_snippet",
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOmittedEmbeddedSchemaAttrErr(tt.err, tt.attr); got != tt.wantOK {
				t.Fatalf("isOmittedEmbeddedSchemaAttrErr(...) = %v, want %v", got, tt.wantOK)
			}
		})
	}
}
