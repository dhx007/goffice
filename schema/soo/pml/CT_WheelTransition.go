package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_WheelTransition struct {
	// Spokes
	SpokesAttr *uint32
}

func NewCT_WheelTransition() *CT_WheelTransition {
	ret := &CT_WheelTransition{}
	return ret
}

func (m *CT_WheelTransition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpokesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spokes"},
			Value: fmt.Sprintf("%v", *m.SpokesAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WheelTransition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spokes" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpokesAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_WheelTransition: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_WheelTransition and its children
func (m *CT_WheelTransition) Validate() error {
	return m.ValidateWithPath("CT_WheelTransition")
}

// ValidateWithPath validates the CT_WheelTransition and its children, prefixing error messages with path
func (m *CT_WheelTransition) ValidateWithPath(path string) error {
	return nil
}
