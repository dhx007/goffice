package powerpoint

import (
	"encoding/xml"
	"fmt"
)

type CT_Empty struct {
}

func NewCT_Empty() *CT_Empty {
	ret := &CT_Empty{}
	return ret
}

func (m *CT_Empty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Empty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Empty: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Empty and its children
func (m *CT_Empty) Validate() error {
	return m.ValidateWithPath("CT_Empty")
}

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (m *CT_Empty) ValidateWithPath(path string) error {
	return nil
}
