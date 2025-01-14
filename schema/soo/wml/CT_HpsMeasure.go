package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_HpsMeasure struct {
	// Half Point Measurement
	ValAttr ST_HpsMeasure
}

func NewCT_HpsMeasure() *CT_HpsMeasure {
	ret := &CT_HpsMeasure{}
	return ret
}

func (m *CT_HpsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HpsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_HpsMeasure(attr.Value)
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
			return fmt.Errorf("parsing CT_HpsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_HpsMeasure and its children
func (m *CT_HpsMeasure) Validate() error {
	return m.ValidateWithPath("CT_HpsMeasure")
}

// ValidateWithPath validates the CT_HpsMeasure and its children, prefixing error messages with path
func (m *CT_HpsMeasure) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
