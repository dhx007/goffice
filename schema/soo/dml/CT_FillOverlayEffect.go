package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_FillOverlayEffect struct {
	BlendAttr ST_BlendMode
	NoFill    *CT_NoFillProperties
	SolidFill *CT_SolidColorFillProperties
	GradFill  *CT_GradientFillProperties
	BlipFill  *CT_BlipFillProperties
	PattFill  *CT_PatternFillProperties
	GrpFill   *CT_GroupFillProperties
}

func NewCT_FillOverlayEffect() *CT_FillOverlayEffect {
	ret := &CT_FillOverlayEffect{}
	ret.BlendAttr = ST_BlendMode(1)
	return ret
}

func (m *CT_FillOverlayEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.BlendAttr.MarshalXMLAttr(xml.Name{Local: "blend"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.BlipFill != nil {
		seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
		e.EncodeElement(m.BlipFill, seblipFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.GrpFill != nil {
		segrpFill := xml.StartElement{Name: xml.Name{Local: "a:grpFill"}}
		e.EncodeElement(m.GrpFill, segrpFill)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FillOverlayEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BlendAttr = ST_BlendMode(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "blend" {
			m.BlendAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_FillOverlayEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "noFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "noFill"}:
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "solidFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "solidFill"}:
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gradFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gradFill"}:
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blipFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blipFill"}:
				m.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pattFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pattFill"}:
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "grpFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "grpFill"}:
				m.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(m.GrpFill, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_FillOverlayEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FillOverlayEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FillOverlayEffect and its children
func (m *CT_FillOverlayEffect) Validate() error {
	return m.ValidateWithPath("CT_FillOverlayEffect")
}

// ValidateWithPath validates the CT_FillOverlayEffect and its children, prefixing error messages with path
func (m *CT_FillOverlayEffect) ValidateWithPath(path string) error {
	if m.BlendAttr == ST_BlendModeUnset {
		return fmt.Errorf("%s/BlendAttr is a mandatory field", path)
	}
	if err := m.BlendAttr.ValidateWithPath(path + "/BlendAttr"); err != nil {
		return err
	}
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.BlipFill != nil {
		if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.GrpFill != nil {
		if err := m.GrpFill.ValidateWithPath(path + "/GrpFill"); err != nil {
			return err
		}
	}
	return nil
}
