package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
)

type CT_DiagramDefinitionHeaderLst struct {
	LayoutDefHdr []*CT_DiagramDefinitionHeader
}

func NewCT_DiagramDefinitionHeaderLst() *CT_DiagramDefinitionHeaderLst {
	ret := &CT_DiagramDefinitionHeaderLst{}
	return ret
}

func (m *CT_DiagramDefinitionHeaderLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.LayoutDefHdr != nil {
		selayoutDefHdr := xml.StartElement{Name: xml.Name{Local: "layoutDefHdr"}}
		for _, c := range m.LayoutDefHdr {
			e.EncodeElement(c, selayoutDefHdr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DiagramDefinitionHeaderLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DiagramDefinitionHeaderLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "layoutDefHdr"}:
				tmp := NewCT_DiagramDefinitionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LayoutDefHdr = append(m.LayoutDefHdr, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_DiagramDefinitionHeaderLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DiagramDefinitionHeaderLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DiagramDefinitionHeaderLst and its children
func (m *CT_DiagramDefinitionHeaderLst) Validate() error {
	return m.ValidateWithPath("CT_DiagramDefinitionHeaderLst")
}

// ValidateWithPath validates the CT_DiagramDefinitionHeaderLst and its children, prefixing error messages with path
func (m *CT_DiagramDefinitionHeaderLst) ValidateWithPath(path string) error {
	for i, v := range m.LayoutDefHdr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LayoutDefHdr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
