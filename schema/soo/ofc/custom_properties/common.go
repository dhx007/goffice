package custom_properties

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Properties", NewCT_Properties)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Property", NewCT_Property)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "Properties", NewProperties)
}
