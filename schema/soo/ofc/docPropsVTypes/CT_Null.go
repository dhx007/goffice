package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
)

type CT_Null struct {
}

func NewCT_Null() *CT_Null {
	ret := &CT_Null{}
	return ret
}

func (m *CT_Null) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Null) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Null: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Null and its children
func (m *CT_Null) Validate() error {
	return m.ValidateWithPath("CT_Null")
}

// ValidateWithPath validates the CT_Null and its children, prefixing error messages with path
func (m *CT_Null) ValidateWithPath(path string) error {
	return nil
}
