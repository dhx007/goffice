package schemaLibrary

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_Schema", NewCT_Schema)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_SchemaLibrary", NewCT_SchemaLibrary)
	goffice.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "schemaLibrary", NewSchemaLibrary)
}
