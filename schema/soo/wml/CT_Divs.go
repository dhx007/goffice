package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Divs struct {
	// Information About Single HTML div Element
	Div []*CT_Div
}

func NewCT_Divs() *CT_Divs {
	ret := &CT_Divs{}
	return ret
}

func (m *CT_Divs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sediv := xml.StartElement{Name: xml.Name{Local: "w:div"}}
	for _, c := range m.Div {
		e.EncodeElement(c, sediv)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Divs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Divs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "div"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "div"}:
				tmp := NewCT_Div()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Div = append(m.Div, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Divs %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Divs
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Divs and its children
func (m *CT_Divs) Validate() error {
	return m.ValidateWithPath("CT_Divs")
}

// ValidateWithPath validates the CT_Divs and its children, prefixing error messages with path
func (m *CT_Divs) ValidateWithPath(path string) error {
	for i, v := range m.Div {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Div[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
