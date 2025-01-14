package vml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type Formulas struct {
	CT_Formulas
}

func NewFormulas() *Formulas {
	ret := &Formulas{}
	ret.CT_Formulas = *NewCT_Formulas()
	return ret
}

func (m *Formulas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Formulas.MarshalXML(e, start)
}

func (m *Formulas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Formulas = *NewCT_Formulas()
lFormulas:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "f"}:
				tmp := NewCT_F()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.F = append(m.F, tmp)
			default:
				goffice.Log("skipping unsupported element on Formulas %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lFormulas
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Formulas and its children
func (m *Formulas) Validate() error {
	return m.ValidateWithPath("Formulas")
}

// ValidateWithPath validates the Formulas and its children, prefixing error messages with path
func (m *Formulas) ValidateWithPath(path string) error {
	if err := m.CT_Formulas.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
