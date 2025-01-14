package chart

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_DispUnitsChoice struct {
	CustUnit    *CT_Double
	BuiltInUnit *CT_BuiltInUnit
}

func NewCT_DispUnitsChoice() *CT_DispUnitsChoice {
	ret := &CT_DispUnitsChoice{}
	return ret
}

func (m *CT_DispUnitsChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CustUnit != nil {
		secustUnit := xml.StartElement{Name: xml.Name{Local: "c:custUnit"}}
		e.EncodeElement(m.CustUnit, secustUnit)
	}
	if m.BuiltInUnit != nil {
		sebuiltInUnit := xml.StartElement{Name: xml.Name{Local: "c:builtInUnit"}}
		e.EncodeElement(m.BuiltInUnit, sebuiltInUnit)
	}
	return nil
}

func (m *CT_DispUnitsChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DispUnitsChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "custUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "custUnit"}:
				m.CustUnit = NewCT_Double()
				if err := d.DecodeElement(m.CustUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "builtInUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "builtInUnit"}:
				m.BuiltInUnit = NewCT_BuiltInUnit()
				if err := d.DecodeElement(m.BuiltInUnit, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_DispUnitsChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DispUnitsChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DispUnitsChoice and its children
func (m *CT_DispUnitsChoice) Validate() error {
	return m.ValidateWithPath("CT_DispUnitsChoice")
}

// ValidateWithPath validates the CT_DispUnitsChoice and its children, prefixing error messages with path
func (m *CT_DispUnitsChoice) ValidateWithPath(path string) error {
	if m.CustUnit != nil {
		if err := m.CustUnit.ValidateWithPath(path + "/CustUnit"); err != nil {
			return err
		}
	}
	if m.BuiltInUnit != nil {
		if err := m.BuiltInUnit.ValidateWithPath(path + "/BuiltInUnit"); err != nil {
			return err
		}
	}
	return nil
}
