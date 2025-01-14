package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_TextAutofit struct {
	NoAutofit   *CT_TextNoAutofit
	NormAutofit *CT_TextNormalAutofit
	SpAutoFit   *CT_TextShapeAutofit
}

func NewEG_TextAutofit() *EG_TextAutofit {
	ret := &EG_TextAutofit{}
	return ret
}

func (m *EG_TextAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NoAutofit != nil {
		senoAutofit := xml.StartElement{Name: xml.Name{Local: "a:noAutofit"}}
		e.EncodeElement(m.NoAutofit, senoAutofit)
	}
	if m.NormAutofit != nil {
		senormAutofit := xml.StartElement{Name: xml.Name{Local: "a:normAutofit"}}
		e.EncodeElement(m.NormAutofit, senormAutofit)
	}
	if m.SpAutoFit != nil {
		sespAutoFit := xml.StartElement{Name: xml.Name{Local: "a:spAutoFit"}}
		e.EncodeElement(m.SpAutoFit, sespAutoFit)
	}
	return nil
}

func (m *EG_TextAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextAutofit:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "noAutofit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "noAutofit"}:
				m.NoAutofit = NewCT_TextNoAutofit()
				if err := d.DecodeElement(m.NoAutofit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "normAutofit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "normAutofit"}:
				m.NormAutofit = NewCT_TextNormalAutofit()
				if err := d.DecodeElement(m.NormAutofit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spAutoFit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spAutoFit"}:
				m.SpAutoFit = NewCT_TextShapeAutofit()
				if err := d.DecodeElement(m.SpAutoFit, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_TextAutofit %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextAutofit
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextAutofit and its children
func (m *EG_TextAutofit) Validate() error {
	return m.ValidateWithPath("EG_TextAutofit")
}

// ValidateWithPath validates the EG_TextAutofit and its children, prefixing error messages with path
func (m *EG_TextAutofit) ValidateWithPath(path string) error {
	if m.NoAutofit != nil {
		if err := m.NoAutofit.ValidateWithPath(path + "/NoAutofit"); err != nil {
			return err
		}
	}
	if m.NormAutofit != nil {
		if err := m.NormAutofit.ValidateWithPath(path + "/NormAutofit"); err != nil {
			return err
		}
	}
	if m.SpAutoFit != nil {
		if err := m.SpAutoFit.ValidateWithPath(path + "/SpAutoFit"); err != nil {
			return err
		}
	}
	return nil
}
