package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TextFont struct {
	TypefaceAttr    string
	PanoseAttr      *string
	PitchFamilyAttr ST_PitchFamily
	CharsetAttr     *int8
}

func NewCT_TextFont() *CT_TextFont {
	ret := &CT_TextFont{}
	return ret
}

func (m *CT_TextFont) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "typeface"},
		Value: fmt.Sprintf("%v", m.TypefaceAttr)})
	if m.PanoseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "panose"},
			Value: fmt.Sprintf("%v", *m.PanoseAttr)})
	}
	if m.PitchFamilyAttr != ST_PitchFamilyUnset {
		attr, err := m.PitchFamilyAttr.MarshalXMLAttr(xml.Name{Local: "pitchFamily"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CharsetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "charset"},
			Value: fmt.Sprintf("%v", *m.CharsetAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextFont) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "typeface" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypefaceAttr = parsed
			continue
		}
		if attr.Name.Local == "panose" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PanoseAttr = &parsed
			continue
		}
		if attr.Name.Local == "pitchFamily" {
			m.PitchFamilyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "charset" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := int8(parsed)
			m.CharsetAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextFont: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextFont and its children
func (m *CT_TextFont) Validate() error {
	return m.ValidateWithPath("CT_TextFont")
}

// ValidateWithPath validates the CT_TextFont and its children, prefixing error messages with path
func (m *CT_TextFont) ValidateWithPath(path string) error {
	if err := m.PitchFamilyAttr.ValidateWithPath(path + "/PitchFamilyAttr"); err != nil {
		return err
	}
	return nil
}
