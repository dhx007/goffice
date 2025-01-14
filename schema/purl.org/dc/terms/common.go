package terms

import "github.com/dhx007/goffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "LCSH", NewLCSH)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "MESH", NewMESH)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "DDC", NewDDC)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "LCC", NewLCC)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "UDC", NewUDC)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "Period", NewPeriod)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "W3CDTF", NewW3CDTF)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "DCMIType", NewDCMIType)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "IMT", NewIMT)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "URI", NewURI)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "ISO639-2", NewISO639_2)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "RFC1766", NewRFC1766)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "RFC3066", NewRFC3066)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "Point", NewPoint)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "ISO3166", NewISO3166)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "Box", NewBox)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "TGN", NewTGN)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "elementOrRefinementContainer", NewElementOrRefinementContainer)
	goffice.RegisterConstructor("http://purl.org/dc/terms/", "elementsAndRefinementsGroup", NewElementsAndRefinementsGroup)
}
