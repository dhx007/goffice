package powerpoint

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Empty", NewCT_Empty)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Rel", NewCT_Rel)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "iscomment", NewIscomment)
	goffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "textdata", NewTextdata)
}
