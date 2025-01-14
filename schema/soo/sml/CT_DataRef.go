package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_DataRef struct {
	// Reference
	RefAttr *string
	// Named Range
	NameAttr *string
	// Sheet Name
	SheetAttr *string
	IdAttr    *string
}

func NewCT_DataRef() *CT_DataRef {
	ret := &CT_DataRef{}
	return ret
}

func (m *CT_DataRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
			Value: fmt.Sprintf("%v", *m.RefAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.SheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheet"},
			Value: fmt.Sprintf("%v", *m.SheetAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "sheet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SheetAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DataRef: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DataRef and its children
func (m *CT_DataRef) Validate() error {
	return m.ValidateWithPath("CT_DataRef")
}

// ValidateWithPath validates the CT_DataRef and its children, prefixing error messages with path
func (m *CT_DataRef) ValidateWithPath(path string) error {
	return nil
}
