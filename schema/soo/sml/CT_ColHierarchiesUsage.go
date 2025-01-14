package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_ColHierarchiesUsage struct {
	// Items Count
	CountAttr *uint32
	// Column OLAP Hierarchies
	ColHierarchyUsage []*CT_HierarchyUsage
}

func NewCT_ColHierarchiesUsage() *CT_ColHierarchiesUsage {
	ret := &CT_ColHierarchiesUsage{}
	return ret
}

func (m *CT_ColHierarchiesUsage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	secolHierarchyUsage := xml.StartElement{Name: xml.Name{Local: "ma:colHierarchyUsage"}}
	for _, c := range m.ColHierarchyUsage {
		e.EncodeElement(c, secolHierarchyUsage)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColHierarchiesUsage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_ColHierarchiesUsage:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "colHierarchyUsage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "colHierarchyUsage"}:
				tmp := NewCT_HierarchyUsage()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ColHierarchyUsage = append(m.ColHierarchyUsage, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_ColHierarchiesUsage %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColHierarchiesUsage
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColHierarchiesUsage and its children
func (m *CT_ColHierarchiesUsage) Validate() error {
	return m.ValidateWithPath("CT_ColHierarchiesUsage")
}

// ValidateWithPath validates the CT_ColHierarchiesUsage and its children, prefixing error messages with path
func (m *CT_ColHierarchiesUsage) ValidateWithPath(path string) error {
	for i, v := range m.ColHierarchyUsage {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ColHierarchyUsage[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
