package content_types

import (
	"regexp"

	"github.com/dhx007/goffice"
)

const ST_ContentTypePattern = `^\p{Latin}+/.*$`

var ST_ContentTypePatternRe = regexp.MustCompile(ST_ContentTypePattern)

const ST_ExtensionPattern = `([!$&'\(\)\*\+,:=]|(%[0-9a-fA-F][0-9a-fA-F])|[:@]|[a-zA-Z0-9\-_~])+`

var ST_ExtensionPatternRe = regexp.MustCompile(ST_ExtensionPattern)

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Types", NewCT_Types)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Default", NewCT_Default)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "CT_Override", NewCT_Override)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Types", NewTypes)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Default", NewDefault)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/content-types", "Override", NewOverride)
}
