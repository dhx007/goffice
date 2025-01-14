package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_CTCategories struct {
	Cat []*CT_CTCategory
}

func NewCT_CTCategories() *CT_CTCategories {
	ret := &CT_CTCategories{}
	return ret
}

func (m *CT_CTCategories) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "cat"}}
		for _, c := range m.Cat {
			e.EncodeElement(c, secat)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CTCategories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CTCategories:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "cat"}:
				tmp := NewCT_CTCategory()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cat = append(m.Cat, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CTCategories %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CTCategories
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CTCategories and its children
func (m *CT_CTCategories) Validate() error {
	return m.ValidateWithPath("CT_CTCategories")
}

// ValidateWithPath validates the CT_CTCategories and its children, prefixing error messages with path
func (m *CT_CTCategories) ValidateWithPath(path string) error {
	for i, v := range m.Cat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
