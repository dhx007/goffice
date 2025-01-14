package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_CustomColorList struct {
	CustClr []*CT_CustomColor
}

func NewCT_CustomColorList() *CT_CustomColorList {
	ret := &CT_CustomColorList{}
	return ret
}

func (m *CT_CustomColorList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CustClr != nil {
		secustClr := xml.StartElement{Name: xml.Name{Local: "a:custClr"}}
		for _, c := range m.CustClr {
			e.EncodeElement(c, secustClr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomColorList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomColorList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "custClr"}:
				tmp := NewCT_CustomColor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustClr = append(m.CustClr, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CustomColorList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomColorList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomColorList and its children
func (m *CT_CustomColorList) Validate() error {
	return m.ValidateWithPath("CT_CustomColorList")
}

// ValidateWithPath validates the CT_CustomColorList and its children, prefixing error messages with path
func (m *CT_CustomColorList) ValidateWithPath(path string) error {
	for i, v := range m.CustClr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustClr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
