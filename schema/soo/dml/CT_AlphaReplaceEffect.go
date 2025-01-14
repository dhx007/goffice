package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_AlphaReplaceEffect struct {
	AAttr ST_PositiveFixedPercentage
}

func NewCT_AlphaReplaceEffect() *CT_AlphaReplaceEffect {
	ret := &CT_AlphaReplaceEffect{}
	return ret
}

func (m *CT_AlphaReplaceEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "a"},
		Value: fmt.Sprintf("%v", m.AAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AlphaReplaceEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "a" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.AAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AlphaReplaceEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AlphaReplaceEffect and its children
func (m *CT_AlphaReplaceEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaReplaceEffect")
}

// ValidateWithPath validates the CT_AlphaReplaceEffect and its children, prefixing error messages with path
func (m *CT_AlphaReplaceEffect) ValidateWithPath(path string) error {
	if err := m.AAttr.ValidateWithPath(path + "/AAttr"); err != nil {
		return err
	}
	return nil
}
