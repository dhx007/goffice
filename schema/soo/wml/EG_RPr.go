package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_RPr struct {
	// Run Properties
	RPr *CT_RPr
}

func NewEG_RPr() *EG_RPr {
	ret := &EG_RPr{}
	return ret
}

func (m *EG_RPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	return nil
}

func (m *EG_RPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_RPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RPr and its children
func (m *EG_RPr) Validate() error {
	return m.ValidateWithPath("EG_RPr")
}

// ValidateWithPath validates the EG_RPr and its children, prefixing error messages with path
func (m *EG_RPr) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
