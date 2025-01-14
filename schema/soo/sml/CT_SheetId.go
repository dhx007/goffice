package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_SheetId struct {
	// Sheet Id
	ValAttr uint32
}

func NewCT_SheetId() *CT_SheetId {
	ret := &CT_SheetId{}
	return ret
}

func (m *CT_SheetId) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetId) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = uint32(parsed)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SheetId: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SheetId and its children
func (m *CT_SheetId) Validate() error {
	return m.ValidateWithPath("CT_SheetId")
}

// ValidateWithPath validates the CT_SheetId and its children, prefixing error messages with path
func (m *CT_SheetId) ValidateWithPath(path string) error {
	return nil
}
