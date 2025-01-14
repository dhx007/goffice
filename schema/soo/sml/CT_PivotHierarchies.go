package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_PivotHierarchies struct {
	// OLAP Hierarchy Count
	CountAttr *uint32
	// OLAP Hierarchy
	PivotHierarchy []*CT_PivotHierarchy
}

func NewCT_PivotHierarchies() *CT_PivotHierarchies {
	ret := &CT_PivotHierarchies{}
	return ret
}

func (m *CT_PivotHierarchies) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sepivotHierarchy := xml.StartElement{Name: xml.Name{Local: "ma:pivotHierarchy"}}
	for _, c := range m.PivotHierarchy {
		e.EncodeElement(c, sepivotHierarchy)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotHierarchies) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PivotHierarchies:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotHierarchy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotHierarchy"}:
				tmp := NewCT_PivotHierarchy()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotHierarchy = append(m.PivotHierarchy, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_PivotHierarchies %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotHierarchies
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotHierarchies and its children
func (m *CT_PivotHierarchies) Validate() error {
	return m.ValidateWithPath("CT_PivotHierarchies")
}

// ValidateWithPath validates the CT_PivotHierarchies and its children, prefixing error messages with path
func (m *CT_PivotHierarchies) ValidateWithPath(path string) error {
	for i, v := range m.PivotHierarchy {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotHierarchy[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
