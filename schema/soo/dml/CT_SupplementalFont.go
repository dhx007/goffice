package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_SupplementalFont struct {
	ScriptAttr   string
	TypefaceAttr string
}

func NewCT_SupplementalFont() *CT_SupplementalFont {
	ret := &CT_SupplementalFont{}
	return ret
}

func (m *CT_SupplementalFont) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "script"},
		Value: fmt.Sprintf("%v", m.ScriptAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "typeface"},
		Value: fmt.Sprintf("%v", m.TypefaceAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SupplementalFont) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "script" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ScriptAttr = parsed
			continue
		}
		if attr.Name.Local == "typeface" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypefaceAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SupplementalFont: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SupplementalFont and its children
func (m *CT_SupplementalFont) Validate() error {
	return m.ValidateWithPath("CT_SupplementalFont")
}

// ValidateWithPath validates the CT_SupplementalFont and its children, prefixing error messages with path
func (m *CT_SupplementalFont) ValidateWithPath(path string) error {
	return nil
}
