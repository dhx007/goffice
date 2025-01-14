package core_properties

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_CoreProperties", NewCT_CoreProperties)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keywords", NewCT_Keywords)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keyword", NewCT_Keyword)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "coreProperties", NewCoreProperties)
}
