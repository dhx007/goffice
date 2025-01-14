package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GeomGuide struct {
	NameAttr string
	FmlaAttr string
}

func NewCT_GeomGuide() *CT_GeomGuide {
	ret := &CT_GeomGuide{}
	return ret
}

func (m *CT_GeomGuide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fmla"},
		Value: fmt.Sprintf("%v", m.FmlaAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GeomGuide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "fmla" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FmlaAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GeomGuide: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GeomGuide and its children
func (m *CT_GeomGuide) Validate() error {
	return m.ValidateWithPath("CT_GeomGuide")
}

// ValidateWithPath validates the CT_GeomGuide and its children, prefixing error messages with path
func (m *CT_GeomGuide) ValidateWithPath(path string) error {
	return nil
}
