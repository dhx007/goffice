package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_MeasureGroup struct {
	// Measure Group Name
	NameAttr string
	// Measure Group Display Name
	CaptionAttr string
}

func NewCT_MeasureGroup() *CT_MeasureGroup {
	ret := &CT_MeasureGroup{}
	return ret
}

func (m *CT_MeasureGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
		Value: fmt.Sprintf("%v", m.CaptionAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MeasureGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MeasureGroup: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MeasureGroup and its children
func (m *CT_MeasureGroup) Validate() error {
	return m.ValidateWithPath("CT_MeasureGroup")
}

// ValidateWithPath validates the CT_MeasureGroup and its children, prefixing error messages with path
func (m *CT_MeasureGroup) ValidateWithPath(path string) error {
	return nil
}
