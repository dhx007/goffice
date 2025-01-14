package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Boolean struct {
	ValAttr *bool
}

func NewCT_Boolean() *CT_Boolean {
	ret := &CT_Boolean{}
	return ret
}

func (m *CT_Boolean) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%d", b2i(*m.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Boolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Boolean: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Boolean and its children
func (m *CT_Boolean) Validate() error {
	return m.ValidateWithPath("CT_Boolean")
}

// ValidateWithPath validates the CT_Boolean and its children, prefixing error messages with path
func (m *CT_Boolean) ValidateWithPath(path string) error {
	return nil
}
