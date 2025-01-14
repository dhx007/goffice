package dml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_TablePartStyle struct {
	TcTxStyle *CT_TableStyleTextStyle
	TcStyle   *CT_TableStyleCellStyle
}

func NewCT_TablePartStyle() *CT_TablePartStyle {
	ret := &CT_TablePartStyle{}
	return ret
}

func (m *CT_TablePartStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.TcTxStyle != nil {
		setcTxStyle := xml.StartElement{Name: xml.Name{Local: "a:tcTxStyle"}}
		e.EncodeElement(m.TcTxStyle, setcTxStyle)
	}
	if m.TcStyle != nil {
		setcStyle := xml.StartElement{Name: xml.Name{Local: "a:tcStyle"}}
		e.EncodeElement(m.TcStyle, setcStyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TablePartStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TablePartStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tcTxStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tcTxStyle"}:
				m.TcTxStyle = NewCT_TableStyleTextStyle()
				if err := d.DecodeElement(m.TcTxStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tcStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tcStyle"}:
				m.TcStyle = NewCT_TableStyleCellStyle()
				if err := d.DecodeElement(m.TcStyle, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_TablePartStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TablePartStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TablePartStyle and its children
func (m *CT_TablePartStyle) Validate() error {
	return m.ValidateWithPath("CT_TablePartStyle")
}

// ValidateWithPath validates the CT_TablePartStyle and its children, prefixing error messages with path
func (m *CT_TablePartStyle) ValidateWithPath(path string) error {
	if m.TcTxStyle != nil {
		if err := m.TcTxStyle.ValidateWithPath(path + "/TcTxStyle"); err != nil {
			return err
		}
	}
	if m.TcStyle != nil {
		if err := m.TcStyle.ValidateWithPath(path + "/TcStyle"); err != nil {
			return err
		}
	}
	return nil
}
