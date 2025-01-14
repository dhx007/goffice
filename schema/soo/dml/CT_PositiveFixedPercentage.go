package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_PositiveFixedPercentage struct {
	ValAttr ST_PositiveFixedPercentage
}

func NewCT_PositiveFixedPercentage() *CT_PositiveFixedPercentage {
	ret := &CT_PositiveFixedPercentage{}
	return ret
}

func (m *CT_PositiveFixedPercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PositiveFixedPercentage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
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
			return fmt.Errorf("parsing CT_PositiveFixedPercentage: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PositiveFixedPercentage and its children
func (m *CT_PositiveFixedPercentage) Validate() error {
	return m.ValidateWithPath("CT_PositiveFixedPercentage")
}

// ValidateWithPath validates the CT_PositiveFixedPercentage and its children, prefixing error messages with path
func (m *CT_PositiveFixedPercentage) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
