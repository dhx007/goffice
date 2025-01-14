package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_ThemeableEffectStyle struct {
	Effect    *CT_EffectProperties
	EffectRef *CT_StyleMatrixReference
}

func NewEG_ThemeableEffectStyle() *EG_ThemeableEffectStyle {
	ret := &EG_ThemeableEffectStyle{}
	return ret
}

func (m *EG_ThemeableEffectStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "a:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.EffectRef != nil {
		seeffectRef := xml.StartElement{Name: xml.Name{Local: "a:effectRef"}}
		e.EncodeElement(m.EffectRef, seeffectRef)
	}
	return nil
}

func (m *EG_ThemeableEffectStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ThemeableEffectStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effect"}:
				m.Effect = NewCT_EffectProperties()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "effectRef"}:
				m.EffectRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.EffectRef, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_ThemeableEffectStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ThemeableEffectStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ThemeableEffectStyle and its children
func (m *EG_ThemeableEffectStyle) Validate() error {
	return m.ValidateWithPath("EG_ThemeableEffectStyle")
}

// ValidateWithPath validates the EG_ThemeableEffectStyle and its children, prefixing error messages with path
func (m *EG_ThemeableEffectStyle) ValidateWithPath(path string) error {
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.EffectRef != nil {
		if err := m.EffectRef.ValidateWithPath(path + "/EffectRef"); err != nil {
			return err
		}
	}
	return nil
}
