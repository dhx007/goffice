package content_types

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Types struct {
	Default  []*Default
	Override []*Override
}

func NewCT_Types() *CT_Types {
	ret := &CT_Types{}
	return ret
}

func (m *CT_Types) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Default != nil {
		seDefault := xml.StartElement{Name: xml.Name{Local: "Default"}}
		for _, c := range m.Default {
			e.EncodeElement(c, seDefault)
		}
	}
	if m.Override != nil {
		seOverride := xml.StartElement{Name: xml.Name{Local: "Override"}}
		for _, c := range m.Override {
			e.EncodeElement(c, seOverride)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Types) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Types:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/content-types", Local: "Default"}:
				tmp := NewDefault()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Default = append(m.Default, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/content-types", Local: "Override"}:
				tmp := NewOverride()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Override = append(m.Override, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Types %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Types
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Types and its children
func (m *CT_Types) Validate() error {
	return m.ValidateWithPath("CT_Types")
}

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (m *CT_Types) ValidateWithPath(path string) error {
	for i, v := range m.Default {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Default[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Override {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Override[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
