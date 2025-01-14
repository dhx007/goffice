package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_TextTabStopList struct {
	Tab []*CT_TextTabStop
}

func NewCT_TextTabStopList() *CT_TextTabStopList {
	ret := &CT_TextTabStopList{}
	return ret
}

func (m *CT_TextTabStopList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Tab != nil {
		setab := xml.StartElement{Name: xml.Name{Local: "a:tab"}}
		for _, c := range m.Tab {
			e.EncodeElement(c, setab)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextTabStopList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextTabStopList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tab"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tab"}:
				tmp := NewCT_TextTabStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tab = append(m.Tab, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_TextTabStopList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextTabStopList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextTabStopList and its children
func (m *CT_TextTabStopList) Validate() error {
	return m.ValidateWithPath("CT_TextTabStopList")
}

// ValidateWithPath validates the CT_TextTabStopList and its children, prefixing error messages with path
func (m *CT_TextTabStopList) ValidateWithPath(path string) error {
	for i, v := range m.Tab {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tab[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
