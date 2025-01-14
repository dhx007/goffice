package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_NumFmt struct {
	// Number Format Id
	NumFmtIdAttr uint32
	// Number Format Code
	FormatCodeAttr string
}

func NewCT_NumFmt() *CT_NumFmt {
	ret := &CT_NumFmt{}
	return ret
}

func (m *CT_NumFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
		Value: fmt.Sprintf("%v", m.NumFmtIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCode"},
		Value: fmt.Sprintf("%v", m.FormatCodeAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NumFmtIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "formatCode" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatCodeAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_NumFmt: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_NumFmt and its children
func (m *CT_NumFmt) Validate() error {
	return m.ValidateWithPath("CT_NumFmt")
}

// ValidateWithPath validates the CT_NumFmt and its children, prefixing error messages with path
func (m *CT_NumFmt) ValidateWithPath(path string) error {
	return nil
}
