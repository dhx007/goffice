package math

import (
	"encoding/xml"
	"fmt"
)

type CT_CtrlPr struct {
}

func NewCT_CtrlPr() *CT_CtrlPr {
	ret := &CT_CtrlPr{}
	return ret
}

func (m *CT_CtrlPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CtrlPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CtrlPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CtrlPr and its children
func (m *CT_CtrlPr) Validate() error {
	return m.ValidateWithPath("CT_CtrlPr")
}

// ValidateWithPath validates the CT_CtrlPr and its children, prefixing error messages with path
func (m *CT_CtrlPr) ValidateWithPath(path string) error {
	return nil
}
