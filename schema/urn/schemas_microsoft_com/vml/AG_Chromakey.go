package vml

import (
	"encoding/xml"
	"fmt"
)

type AG_Chromakey struct {
	ChromakeyAttr *string
}

func NewAG_Chromakey() *AG_Chromakey {
	ret := &AG_Chromakey{}
	return ret
}

func (m *AG_Chromakey) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ChromakeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chromakey"},
			Value: fmt.Sprintf("%v", *m.ChromakeyAttr)})
	}
	return nil
}

func (m *AG_Chromakey) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Chromakey: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Chromakey and its children
func (m *AG_Chromakey) Validate() error {
	return m.ValidateWithPath("AG_Chromakey")
}

// ValidateWithPath validates the AG_Chromakey and its children, prefixing error messages with path
func (m *AG_Chromakey) ValidateWithPath(path string) error {
	return nil
}
