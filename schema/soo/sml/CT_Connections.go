package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_Connections struct {
	// Connection
	Connection []*CT_Connection
}

func NewCT_Connections() *CT_Connections {
	ret := &CT_Connections{}
	return ret
}

func (m *CT_Connections) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seconnection := xml.StartElement{Name: xml.Name{Local: "ma:connection"}}
	for _, c := range m.Connection {
		e.EncodeElement(c, seconnection)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Connections) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Connections:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "connection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "connection"}:
				tmp := NewCT_Connection()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Connection = append(m.Connection, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_Connections %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Connections
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Connections and its children
func (m *CT_Connections) Validate() error {
	return m.ValidateWithPath("CT_Connections")
}

// ValidateWithPath validates the CT_Connections and its children, prefixing error messages with path
func (m *CT_Connections) ValidateWithPath(path string) error {
	for i, v := range m.Connection {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Connection[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
