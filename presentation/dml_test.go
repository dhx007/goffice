package presentation_test

import (
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

// Issue #207
func TestParseUnionST_AdjCoordinate(t *testing.T) {
	// this crashed due to a null pointer dereferences when not initializing the
	// returned value correctly
	dml.ParseUnionST_AdjCoordinate("123")
}
