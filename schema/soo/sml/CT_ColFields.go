package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_ColFields struct {
	// Repeated Items Count
	CountAttr *uint32
	// Field
	Field []*CT_Field
}

func NewCT_ColFields() *CT_ColFields {
	ret := &CT_ColFields{}
	return ret
}

func (m *CT_ColFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sefield := xml.StartElement{Name: xml.Name{Local: "ma:field"}}
	for _, c := range m.Field {
		e.EncodeElement(c, sefield)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_ColFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "field"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "field"}:
				tmp := NewCT_Field()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Field = append(m.Field, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_ColFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColFields and its children
func (m *CT_ColFields) Validate() error {
	return m.ValidateWithPath("CT_ColFields")
}

// ValidateWithPath validates the CT_ColFields and its children, prefixing error messages with path
func (m *CT_ColFields) ValidateWithPath(path string) error {
	for i, v := range m.Field {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Field[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
