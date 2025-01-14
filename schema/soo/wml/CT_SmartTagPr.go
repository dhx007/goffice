package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_SmartTagPr struct {
	// Smart Tag Property
	Attr []*CT_Attr
}

func NewCT_SmartTagPr() *CT_SmartTagPr {
	ret := &CT_SmartTagPr{}
	return ret
}

func (m *CT_SmartTagPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Attr != nil {
		seattr := xml.StartElement{Name: xml.Name{Local: "w:attr"}}
		for _, c := range m.Attr {
			e.EncodeElement(c, seattr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SmartTagPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SmartTagPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "attr"}:
				tmp := NewCT_Attr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Attr = append(m.Attr, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_SmartTagPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SmartTagPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SmartTagPr and its children
func (m *CT_SmartTagPr) Validate() error {
	return m.ValidateWithPath("CT_SmartTagPr")
}

// ValidateWithPath validates the CT_SmartTagPr and its children, prefixing error messages with path
func (m *CT_SmartTagPr) ValidateWithPath(path string) error {
	for i, v := range m.Attr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Attr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
