package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_CacheHierarchies struct {
	// Hierarchy Count
	CountAttr *uint32
	// PivotCache Hierarchy
	CacheHierarchy []*CT_CacheHierarchy
}

func NewCT_CacheHierarchies() *CT_CacheHierarchies {
	ret := &CT_CacheHierarchies{}
	return ret
}

func (m *CT_CacheHierarchies) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.CacheHierarchy != nil {
		secacheHierarchy := xml.StartElement{Name: xml.Name{Local: "ma:cacheHierarchy"}}
		for _, c := range m.CacheHierarchy {
			e.EncodeElement(c, secacheHierarchy)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CacheHierarchies) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_CacheHierarchies:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cacheHierarchy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cacheHierarchy"}:
				tmp := NewCT_CacheHierarchy()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CacheHierarchy = append(m.CacheHierarchy, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CacheHierarchies %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CacheHierarchies
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CacheHierarchies and its children
func (m *CT_CacheHierarchies) Validate() error {
	return m.ValidateWithPath("CT_CacheHierarchies")
}

// ValidateWithPath validates the CT_CacheHierarchies and its children, prefixing error messages with path
func (m *CT_CacheHierarchies) ValidateWithPath(path string) error {
	for i, v := range m.CacheHierarchy {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CacheHierarchy[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
