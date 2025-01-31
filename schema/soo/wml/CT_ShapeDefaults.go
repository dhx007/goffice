package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_ShapeDefaults struct {
	Any []goffice.Any
}

func NewCT_ShapeDefaults() *CT_ShapeDefaults {
	ret := &CT_ShapeDefaults{}
	return ret
}

func (m *CT_ShapeDefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Any != nil {
		for _, c := range m.Any {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeDefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ShapeDefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := goffice.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = append(m.Any, anyEl)
				}
			}
		case xml.EndElement:
			break lCT_ShapeDefaults
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShapeDefaults and its children
func (m *CT_ShapeDefaults) Validate() error {
	return m.ValidateWithPath("CT_ShapeDefaults")
}

// ValidateWithPath validates the CT_ShapeDefaults and its children, prefixing error messages with path
func (m *CT_ShapeDefaults) ValidateWithPath(path string) error {
	return nil
}
