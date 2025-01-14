package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PaperSource struct {
	// First Page Printer Tray Code
	FirstAttr *int64
	// Non-First Page Printer Tray Code
	OtherAttr *int64
}

func NewCT_PaperSource() *CT_PaperSource {
	ret := &CT_PaperSource{}
	return ret
}

func (m *CT_PaperSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FirstAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:first"},
			Value: fmt.Sprintf("%v", *m.FirstAttr)})
	}
	if m.OtherAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:other"},
			Value: fmt.Sprintf("%v", *m.OtherAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PaperSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "first" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.FirstAttr = &parsed
			continue
		}
		if attr.Name.Local == "other" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.OtherAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PaperSource: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PaperSource and its children
func (m *CT_PaperSource) Validate() error {
	return m.ValidateWithPath("CT_PaperSource")
}

// ValidateWithPath validates the CT_PaperSource and its children, prefixing error messages with path
func (m *CT_PaperSource) ValidateWithPath(path string) error {
	return nil
}
