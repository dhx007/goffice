package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_AxisUnit struct {
	ValAttr float64
}

func NewCT_AxisUnit() *CT_AxisUnit {
	ret := &CT_AxisUnit{}
	ret.ValAttr = 0 + 1
	return ret
}

func (m *CT_AxisUnit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AxisUnit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = 0 + 1
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
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
			return fmt.Errorf("parsing CT_AxisUnit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AxisUnit and its children
func (m *CT_AxisUnit) Validate() error {
	return m.ValidateWithPath("CT_AxisUnit")
}

// ValidateWithPath validates the CT_AxisUnit and its children, prefixing error messages with path
func (m *CT_AxisUnit) ValidateWithPath(path string) error {
	if m.ValAttr <= 0 {
		return fmt.Errorf("%s/m.ValAttr must be > 0 (have %v)", path, m.ValAttr)
	}
	return nil
}
