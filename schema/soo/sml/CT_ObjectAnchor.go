package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/dml/spreadsheetDrawing"
)

type CT_ObjectAnchor struct {
	// Move With Cells
	MoveWithCellsAttr *bool
	// Size With Cells
	SizeWithCellsAttr *bool
	From              *spreadsheetDrawing.From
	To                *spreadsheetDrawing.To
}

func NewCT_ObjectAnchor() *CT_ObjectAnchor {
	ret := &CT_ObjectAnchor{}
	ret.From = spreadsheetDrawing.NewFrom()
	ret.To = spreadsheetDrawing.NewTo()
	return ret
}

func (m *CT_ObjectAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MoveWithCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "moveWithCells"},
			Value: fmt.Sprintf("%d", b2i(*m.MoveWithCellsAttr))})
	}
	if m.SizeWithCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sizeWithCells"},
			Value: fmt.Sprintf("%d", b2i(*m.SizeWithCellsAttr))})
	}
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "xdr:from"}}
	e.EncodeElement(m.From, sefrom)
	seto := xml.StartElement{Name: xml.Name{Local: "xdr:to"}}
	e.EncodeElement(m.To, seto)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ObjectAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = spreadsheetDrawing.NewFrom()
	m.To = spreadsheetDrawing.NewTo()
	for _, attr := range start.Attr {
		if attr.Name.Local == "moveWithCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MoveWithCellsAttr = &parsed
			continue
		}
		if attr.Name.Local == "sizeWithCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SizeWithCellsAttr = &parsed
			continue
		}
	}
lCT_ObjectAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "from"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "from"}:
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "to"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "to"}:
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_ObjectAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ObjectAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ObjectAnchor and its children
func (m *CT_ObjectAnchor) Validate() error {
	return m.ValidateWithPath("CT_ObjectAnchor")
}

// ValidateWithPath validates the CT_ObjectAnchor and its children, prefixing error messages with path
func (m *CT_ObjectAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.To.ValidateWithPath(path + "/To"); err != nil {
		return err
	}
	return nil
}
