package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Controls struct {
	// Embedded Control
	Control []*CT_Control
}

func NewCT_Controls() *CT_Controls {
	ret := &CT_Controls{}
	return ret
}

func (m *CT_Controls) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secontrol := xml.StartElement{Name: xml.Name{Local: "ma:control"}}
	for _, c := range m.Control {
		e.EncodeElement(c, secontrol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Controls) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Controls:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "control"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "control"}:
				tmp := NewCT_Control()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Control = append(m.Control, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Controls %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Controls
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Controls and its children
func (m *CT_Controls) Validate() error {
	return m.ValidateWithPath("CT_Controls")
}

// ValidateWithPath validates the CT_Controls and its children, prefixing error messages with path
func (m *CT_Controls) ValidateWithPath(path string) error {
	for i, v := range m.Control {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Control[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
