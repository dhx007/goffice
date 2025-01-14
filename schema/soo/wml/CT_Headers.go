package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Headers struct {
	// Header Cell Reference
	Header []*CT_String
}

func NewCT_Headers() *CT_Headers {
	ret := &CT_Headers{}
	return ret
}

func (m *CT_Headers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seheader := xml.StartElement{Name: xml.Name{Local: "w:header"}}
	for _, c := range m.Header {
		e.EncodeElement(c, seheader)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Headers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Headers:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "header"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "header"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Header = append(m.Header, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Headers %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Headers
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Headers and its children
func (m *CT_Headers) Validate() error {
	return m.ValidateWithPath("CT_Headers")
}

// ValidateWithPath validates the CT_Headers and its children, prefixing error messages with path
func (m *CT_Headers) ValidateWithPath(path string) error {
	for i, v := range m.Header {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Header[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
