package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_FlatText struct {
	ZAttr *ST_Coordinate
}

func NewCT_FlatText() *CT_FlatText {
	ret := &CT_FlatText{}
	return ret
}

func (m *CT_FlatText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ZAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "z"},
			Value: fmt.Sprintf("%v", *m.ZAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FlatText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "z" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.ZAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_FlatText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_FlatText and its children
func (m *CT_FlatText) Validate() error {
	return m.ValidateWithPath("CT_FlatText")
}

// ValidateWithPath validates the CT_FlatText and its children, prefixing error messages with path
func (m *CT_FlatText) ValidateWithPath(path string) error {
	if m.ZAttr != nil {
		if err := m.ZAttr.ValidateWithPath(path + "/ZAttr"); err != nil {
			return err
		}
	}
	return nil
}
