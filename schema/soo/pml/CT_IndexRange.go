package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_IndexRange struct {
	// Start
	StAttr uint32
	// End
	EndAttr uint32
}

func NewCT_IndexRange() *CT_IndexRange {
	ret := &CT_IndexRange{}
	return ret
}

func (m *CT_IndexRange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
		Value: fmt.Sprintf("%v", m.StAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "end"},
		Value: fmt.Sprintf("%v", m.EndAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IndexRange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "st" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.StAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "end" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.EndAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_IndexRange: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_IndexRange and its children
func (m *CT_IndexRange) Validate() error {
	return m.ValidateWithPath("CT_IndexRange")
}

// ValidateWithPath validates the CT_IndexRange and its children, prefixing error messages with path
func (m *CT_IndexRange) ValidateWithPath(path string) error {
	return nil
}
