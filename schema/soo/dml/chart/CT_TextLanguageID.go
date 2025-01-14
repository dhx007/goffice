package chart

import (
	"encoding/xml"
	"fmt"
)

type CT_TextLanguageID struct {
	ValAttr string
}

func NewCT_TextLanguageID() *CT_TextLanguageID {
	ret := &CT_TextLanguageID{}
	return ret
}

func (m *CT_TextLanguageID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextLanguageID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			return fmt.Errorf("parsing CT_TextLanguageID: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TextLanguageID and its children
func (m *CT_TextLanguageID) Validate() error {
	return m.ValidateWithPath("CT_TextLanguageID")
}

// ValidateWithPath validates the CT_TextLanguageID and its children, prefixing error messages with path
func (m *CT_TextLanguageID) ValidateWithPath(path string) error {
	return nil
}
