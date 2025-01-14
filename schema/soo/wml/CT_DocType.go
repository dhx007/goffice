package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_DocType struct {
	// Document Classification Value
	ValAttr string
}

func NewCT_DocType() *CT_DocType {
	ret := &CT_DocType{}
	return ret
}

func (m *CT_DocType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
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
			return fmt.Errorf("parsing CT_DocType: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocType and its children
func (m *CT_DocType) Validate() error {
	return m.ValidateWithPath("CT_DocType")
}

// ValidateWithPath validates the CT_DocType and its children, prefixing error messages with path
func (m *CT_DocType) ValidateWithPath(path string) error {
	return nil
}
