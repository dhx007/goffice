package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_TextNoAutofit struct {
}

func NewCT_TextNoAutofit() *CT_TextNoAutofit {
	ret := &CT_TextNoAutofit{}
	return ret
}

func (m *CT_TextNoAutofit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextNoAutofit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextNoAutofit: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextNoAutofit and its children
func (m *CT_TextNoAutofit) Validate() error {
	return m.ValidateWithPath("CT_TextNoAutofit")
}

// ValidateWithPath validates the CT_TextNoAutofit and its children, prefixing error messages with path
func (m *CT_TextNoAutofit) ValidateWithPath(path string) error {
	return nil
}
