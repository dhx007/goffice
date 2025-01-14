package spreadsheet

import (
	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/sml"
)

// DefinedName is a named range, formula, etc.
type DefinedName struct {
	x *sml.CT_DefinedName
}

// X returns the inner wrapped XML type.
func (d DefinedName) X() *sml.CT_DefinedName {
	return d.x
}

// Name returns the name of the defined name.
func (d DefinedName) Name() string {
	return d.x.NameAttr
}

// Content returns the content of the defined range (the range in most cases)/
func (d DefinedName) Content() string {
	return d.x.Content
}

// SetContent sets the defined name content.
func (d DefinedName) SetContent(s string) {
	d.x.Content = s
}

// SetHidden marks the defined name as hidden.
func (d DefinedName) SetHidden(b bool) {
	d.x.HiddenAttr = goffice.Bool(b)
}

// SetHidden marks the defined name as hidden.
func (d DefinedName) SetLocalSheetID(id uint32) {
	d.x.LocalSheetIdAttr = goffice.Uint32(id)
}
