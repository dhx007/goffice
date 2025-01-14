package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_CellStyles struct {
	// Style Count
	CountAttr *uint32
	// Cell Style
	CellStyle []*CT_CellStyle
}

func NewCT_CellStyles() *CT_CellStyles {
	ret := &CT_CellStyles{}
	return ret
}

func (m *CT_CellStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	secellStyle := xml.StartElement{Name: xml.Name{Local: "ma:cellStyle"}}
	for _, c := range m.CellStyle {
		e.EncodeElement(c, secellStyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_CellStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cellStyle"}:
				tmp := NewCT_CellStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellStyle = append(m.CellStyle, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CellStyles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellStyles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellStyles and its children
func (m *CT_CellStyles) Validate() error {
	return m.ValidateWithPath("CT_CellStyles")
}

// ValidateWithPath validates the CT_CellStyles and its children, prefixing error messages with path
func (m *CT_CellStyles) ValidateWithPath(path string) error {
	for i, v := range m.CellStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
