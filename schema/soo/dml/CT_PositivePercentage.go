package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_PositivePercentage struct {
	ValAttr ST_PositivePercentage
}

func NewCT_PositivePercentage() *CT_PositivePercentage {
	ret := &CT_PositivePercentage{}
	return ret
}

func (m *CT_PositivePercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PositivePercentage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_PositivePercentage(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PositivePercentage: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PositivePercentage and its children
func (m *CT_PositivePercentage) Validate() error {
	return m.ValidateWithPath("CT_PositivePercentage")
}

// ValidateWithPath validates the CT_PositivePercentage and its children, prefixing error messages with path
func (m *CT_PositivePercentage) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
