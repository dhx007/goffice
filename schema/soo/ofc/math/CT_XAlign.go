package math

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_XAlign struct {
	ValAttr sharedTypes.ST_XAlign
}

func NewCT_XAlign() *CT_XAlign {
	ret := &CT_XAlign{}
	ret.ValAttr = sharedTypes.ST_XAlign(1)
	return ret
}

func (m *CT_XAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "m:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_XAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = sharedTypes.ST_XAlign(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_XAlign: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_XAlign and its children
func (m *CT_XAlign) Validate() error {
	return m.ValidateWithPath("CT_XAlign")
}

// ValidateWithPath validates the CT_XAlign and its children, prefixing error messages with path
func (m *CT_XAlign) ValidateWithPath(path string) error {
	if m.ValAttr == sharedTypes.ST_XAlignUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
