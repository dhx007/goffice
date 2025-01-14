package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_colItems struct {
	// Column Item Count
	CountAttr *uint32
	// Column Items
	I []*CT_I
}

func NewCT_colItems() *CT_colItems {
	ret := &CT_colItems{}
	return ret
}

func (m *CT_colItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sei := xml.StartElement{Name: xml.Name{Local: "ma:i"}}
	for _, c := range m.I {
		e.EncodeElement(c, sei)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_colItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_colItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "i"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "i"}:
				tmp := NewCT_I()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.I = append(m.I, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_colItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_colItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_colItems and its children
func (m *CT_colItems) Validate() error {
	return m.ValidateWithPath("CT_colItems")
}

// ValidateWithPath validates the CT_colItems and its children, prefixing error messages with path
func (m *CT_colItems) ValidateWithPath(path string) error {
	for i, v := range m.I {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/I[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
