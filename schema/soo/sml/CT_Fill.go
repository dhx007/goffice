package sml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_Fill struct {
	// Pattern
	PatternFill *CT_PatternFill
	// Gradient
	GradientFill *CT_GradientFill
}

func NewCT_Fill() *CT_Fill {
	ret := &CT_Fill{}
	return ret
}

func (m *CT_Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.PatternFill != nil {
		sepatternFill := xml.StartElement{Name: xml.Name{Local: "ma:patternFill"}}
		e.EncodeElement(m.PatternFill, sepatternFill)
	}
	if m.GradientFill != nil {
		segradientFill := xml.StartElement{Name: xml.Name{Local: "ma:gradientFill"}}
		e.EncodeElement(m.GradientFill, segradientFill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Fill:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "patternFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "patternFill"}:
				m.PatternFill = NewCT_PatternFill()
				if err := d.DecodeElement(m.PatternFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "gradientFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "gradientFill"}:
				m.GradientFill = NewCT_GradientFill()
				if err := d.DecodeElement(m.GradientFill, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_Fill %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Fill
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Fill and its children
func (m *CT_Fill) Validate() error {
	return m.ValidateWithPath("CT_Fill")
}

// ValidateWithPath validates the CT_Fill and its children, prefixing error messages with path
func (m *CT_Fill) ValidateWithPath(path string) error {
	if m.PatternFill != nil {
		if err := m.PatternFill.ValidateWithPath(path + "/PatternFill"); err != nil {
			return err
		}
	}
	if m.GradientFill != nil {
		if err := m.GradientFill.ValidateWithPath(path + "/GradientFill"); err != nil {
			return err
		}
	}
	return nil
}
