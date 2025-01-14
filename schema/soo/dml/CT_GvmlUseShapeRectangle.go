package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GvmlUseShapeRectangle struct {
}

func NewCT_GvmlUseShapeRectangle() *CT_GvmlUseShapeRectangle {
	ret := &CT_GvmlUseShapeRectangle{}
	return ret
}

func (m *CT_GvmlUseShapeRectangle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlUseShapeRectangle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GvmlUseShapeRectangle: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GvmlUseShapeRectangle and its children
func (m *CT_GvmlUseShapeRectangle) Validate() error {
	return m.ValidateWithPath("CT_GvmlUseShapeRectangle")
}

// ValidateWithPath validates the CT_GvmlUseShapeRectangle and its children, prefixing error messages with path
func (m *CT_GvmlUseShapeRectangle) ValidateWithPath(path string) error {
	return nil
}
