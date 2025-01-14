package vml

import (
	"encoding/xml"
	"fmt"
)

type CT_F struct {
	EqnAttr *string
}

func NewCT_F() *CT_F {
	ret := &CT_F{}
	return ret
}

func (m *CT_F) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EqnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "eqn"},
			Value: fmt.Sprintf("%v", *m.EqnAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_F) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "eqn" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EqnAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_F: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_F and its children
func (m *CT_F) Validate() error {
	return m.ValidateWithPath("CT_F")
}

// ValidateWithPath validates the CT_F and its children, prefixing error messages with path
func (m *CT_F) ValidateWithPath(path string) error {
	return nil
}
