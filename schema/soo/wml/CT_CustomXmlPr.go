package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_CustomXmlPr struct {
	// Custom XML Element Placeholder Text
	Placeholder *CT_String
	// Custom XML Attribute
	Attr []*CT_Attr
}

func NewCT_CustomXmlPr() *CT_CustomXmlPr {
	ret := &CT_CustomXmlPr{}
	return ret
}

func (m *CT_CustomXmlPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Placeholder != nil {
		seplaceholder := xml.StartElement{Name: xml.Name{Local: "w:placeholder"}}
		e.EncodeElement(m.Placeholder, seplaceholder)
	}
	if m.Attr != nil {
		seattr := xml.StartElement{Name: xml.Name{Local: "w:attr"}}
		for _, c := range m.Attr {
			e.EncodeElement(c, seattr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomXmlPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomXmlPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "placeholder"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "placeholder"}:
				m.Placeholder = NewCT_String()
				if err := d.DecodeElement(m.Placeholder, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "attr"}:
				tmp := NewCT_Attr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Attr = append(m.Attr, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CustomXmlPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomXmlPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomXmlPr and its children
func (m *CT_CustomXmlPr) Validate() error {
	return m.ValidateWithPath("CT_CustomXmlPr")
}

// ValidateWithPath validates the CT_CustomXmlPr and its children, prefixing error messages with path
func (m *CT_CustomXmlPr) ValidateWithPath(path string) error {
	if m.Placeholder != nil {
		if err := m.Placeholder.ValidateWithPath(path + "/Placeholder"); err != nil {
			return err
		}
	}
	for i, v := range m.Attr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Attr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
