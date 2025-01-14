package extended_properties

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_Properties", NewCT_Properties)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorVariant", NewCT_VectorVariant)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorLpstr", NewCT_VectorLpstr)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_DigSigBlob", NewCT_DigSigBlob)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "Properties", NewProperties)
}
