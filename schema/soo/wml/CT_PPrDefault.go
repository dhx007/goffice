package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_PPrDefault struct {
	// Paragraph Properties
	PPr *CT_PPrGeneral
}

func NewCT_PPrDefault() *CT_PPrDefault {
	ret := &CT_PPrDefault{}
	return ret
}

func (m *CT_PPrDefault) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.PPr != nil {
		sepPr := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}
		e.EncodeElement(m.PPr, sepPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PPrDefault) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PPrDefault:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pPr"}:
				m.PPr = NewCT_PPrGeneral()
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_PPrDefault %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PPrDefault
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PPrDefault and its children
func (m *CT_PPrDefault) Validate() error {
	return m.ValidateWithPath("CT_PPrDefault")
}

// ValidateWithPath validates the CT_PPrDefault and its children, prefixing error messages with path
func (m *CT_PPrDefault) ValidateWithPath(path string) error {
	if m.PPr != nil {
		if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
			return err
		}
	}
	return nil
}
