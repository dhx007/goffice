package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_TableGrid struct {
	GridCol []*CT_TableCol
}

func NewCT_TableGrid() *CT_TableGrid {
	ret := &CT_TableGrid{}
	return ret
}

func (m *CT_TableGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.GridCol != nil {
		segridCol := xml.StartElement{Name: xml.Name{Local: "a:gridCol"}}
		for _, c := range m.GridCol {
			e.EncodeElement(c, segridCol)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TableGrid:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gridCol"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gridCol"}:
				tmp := NewCT_TableCol()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridCol = append(m.GridCol, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_TableGrid %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableGrid
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableGrid and its children
func (m *CT_TableGrid) Validate() error {
	return m.ValidateWithPath("CT_TableGrid")
}

// ValidateWithPath validates the CT_TableGrid and its children, prefixing error messages with path
func (m *CT_TableGrid) ValidateWithPath(path string) error {
	for i, v := range m.GridCol {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridCol[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
