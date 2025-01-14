package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_Height struct {
	// Table Row Height
	ValAttr *sharedTypes.ST_TwipsMeasure
	// Table Row Height Type
	HRuleAttr ST_HeightRule
}

func NewCT_Height() *CT_Height {
	ret := &CT_Height{}
	return ret
}

func (m *CT_Height) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.HRuleAttr != ST_HeightRuleUnset {
		attr, err := m.HRuleAttr.MarshalXMLAttr(xml.Name{Local: "w:hRule"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Height) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
		if attr.Name.Local == "hRule" {
			m.HRuleAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Height: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Height and its children
func (m *CT_Height) Validate() error {
	return m.ValidateWithPath("CT_Height")
}

// ValidateWithPath validates the CT_Height and its children, prefixing error messages with path
func (m *CT_Height) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
			return err
		}
	}
	if err := m.HRuleAttr.ValidateWithPath(path + "/HRuleAttr"); err != nil {
		return err
	}
	return nil
}
