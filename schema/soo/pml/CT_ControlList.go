package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_ControlList struct {
	// Embedded Control
	Control []*CT_Control
}

func NewCT_ControlList() *CT_ControlList {
	ret := &CT_ControlList{}
	return ret
}

func (m *CT_ControlList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Control != nil {
		secontrol := xml.StartElement{Name: xml.Name{Local: "p:control"}}
		for _, c := range m.Control {
			e.EncodeElement(c, secontrol)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ControlList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ControlList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "control"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "control"}:
				tmp := NewCT_Control()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Control = append(m.Control, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_ControlList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ControlList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ControlList and its children
func (m *CT_ControlList) Validate() error {
	return m.ValidateWithPath("CT_ControlList")
}

// ValidateWithPath validates the CT_ControlList and its children, prefixing error messages with path
func (m *CT_ControlList) ValidateWithPath(path string) error {
	for i, v := range m.Control {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Control[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
