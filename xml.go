package goffice

import "encoding/xml"

// NeedsSpacePreserve returns true if the string has leading or trailing space.
func NeedsSpacePreserve(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	switch s[len(s)-1] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

// AddPreserveSpaceAttr adds an xml:space="preserve" attribute to a start
// element if it is required for the string s.
func AddPreserveSpaceAttr(se *xml.StartElement, s string) {
	if NeedsSpacePreserve(s) {
		se.Attr = append(se.Attr, xml.Attr{Name: xml.Name{Local: "xml:space"}, Value: "preserve"})
	}
}
