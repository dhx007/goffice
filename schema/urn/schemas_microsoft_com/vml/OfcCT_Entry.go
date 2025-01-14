package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type OfcCT_Entry struct {
	NewAttr *int32
	OldAttr *int32
}

func NewOfcCT_Entry() *OfcCT_Entry {
	ret := &OfcCT_Entry{}
	return ret
}

func (m *OfcCT_Entry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "new"},
			Value: fmt.Sprintf("%v", *m.NewAttr)})
	}
	if m.OldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "old"},
			Value: fmt.Sprintf("%v", *m.OldAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Entry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "new" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.NewAttr = &pt
			continue
		}
		if attr.Name.Local == "old" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.OldAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Entry: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Entry and its children
func (m *OfcCT_Entry) Validate() error {
	return m.ValidateWithPath("OfcCT_Entry")
}

// ValidateWithPath validates the OfcCT_Entry and its children, prefixing error messages with path
func (m *OfcCT_Entry) ValidateWithPath(path string) error {
	return nil
}
