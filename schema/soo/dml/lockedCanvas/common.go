package lockedCanvas

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas", "lockedCanvas", NewLockedCanvas)
}
