package sml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_DdeValue struct {
	// DDE Value Type
	TAttr ST_DdeValueType
	// DDE Link Value
	Val string
}

func NewCT_DdeValue() *CT_DdeValue {
	ret := &CT_DdeValue{}
	return ret
}

func (m *CT_DdeValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TAttr != ST_DdeValueTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	seval := xml.StartElement{Name: xml.Name{Local: "ma:val"}}
	goffice.AddPreserveSpaceAttr(&seval, m.Val)
	e.EncodeElement(m.Val, seval)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DdeValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_DdeValue:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "val"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "val"}:
				if err := d.DecodeElement(&m.Val, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_DdeValue %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeValue
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DdeValue and its children
func (m *CT_DdeValue) Validate() error {
	return m.ValidateWithPath("CT_DdeValue")
}

// ValidateWithPath validates the CT_DdeValue and its children, prefixing error messages with path
func (m *CT_DdeValue) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	return nil
}
