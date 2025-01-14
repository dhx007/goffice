package elements

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://purl.org/dc/elements/1.1/", "SimpleLiteral", NewSimpleLiteral)
	goffice.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementContainer", NewElementContainer)
	goffice.RegisterConstructor("http://purl.org/dc/elements/1.1/", "any", NewAny)
	goffice.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementsGroup", NewElementsGroup)
}
