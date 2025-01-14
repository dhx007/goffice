package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_SaveThroughXslt struct {
	IdAttr *string
	// Local Identifier for XSL Transform
	SolutionIDAttr *string
}

func NewCT_SaveThroughXslt() *CT_SaveThroughXslt {
	ret := &CT_SaveThroughXslt{}
	return ret
}

func (m *CT_SaveThroughXslt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.SolutionIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:solutionID"},
			Value: fmt.Sprintf("%v", *m.SolutionIDAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SaveThroughXslt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "solutionID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SolutionIDAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SaveThroughXslt: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SaveThroughXslt and its children
func (m *CT_SaveThroughXslt) Validate() error {
	return m.ValidateWithPath("CT_SaveThroughXslt")
}

// ValidateWithPath validates the CT_SaveThroughXslt and its children, prefixing error messages with path
func (m *CT_SaveThroughXslt) ValidateWithPath(path string) error {
	return nil
}
