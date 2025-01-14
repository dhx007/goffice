package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_LuminanceEffect struct {
	BrightAttr   *ST_FixedPercentage
	ContrastAttr *ST_FixedPercentage
}

func NewCT_LuminanceEffect() *CT_LuminanceEffect {
	ret := &CT_LuminanceEffect{}
	return ret
}

func (m *CT_LuminanceEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BrightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bright"},
			Value: fmt.Sprintf("%v", *m.BrightAttr)})
	}
	if m.ContrastAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contrast"},
			Value: fmt.Sprintf("%v", *m.ContrastAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LuminanceEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bright" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.BrightAttr = &parsed
			continue
		}
		if attr.Name.Local == "contrast" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ContrastAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LuminanceEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LuminanceEffect and its children
func (m *CT_LuminanceEffect) Validate() error {
	return m.ValidateWithPath("CT_LuminanceEffect")
}

// ValidateWithPath validates the CT_LuminanceEffect and its children, prefixing error messages with path
func (m *CT_LuminanceEffect) ValidateWithPath(path string) error {
	if m.BrightAttr != nil {
		if err := m.BrightAttr.ValidateWithPath(path + "/BrightAttr"); err != nil {
			return err
		}
	}
	if m.ContrastAttr != nil {
		if err := m.ContrastAttr.ValidateWithPath(path + "/ContrastAttr"); err != nil {
			return err
		}
	}
	return nil
}
