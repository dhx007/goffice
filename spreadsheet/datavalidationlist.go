package spreadsheet

import (
	"strings"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/sml"
)

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct {
	x *sml.CT_DataValidation
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (d DataValidationList) SetRange(cellRange string) {
	d.x.Formula1 = goffice.String(cellRange)
	d.x.Formula2 = goffice.String("0")
}

// SetValues sets the possible values. This is incompatible with SetRange.
func (d DataValidationList) SetValues(values []string) {
	d.x.Formula1 = goffice.String("\"" + strings.Join(values, ",") + "\"")
	d.x.Formula2 = goffice.String("0")
}
