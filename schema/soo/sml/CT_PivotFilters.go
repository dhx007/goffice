package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_PivotFilters struct {
	// Pivot Filter Count
	CountAttr *uint32
	// PivotTable Advanced Filter
	Filter []*CT_PivotFilter
}

func NewCT_PivotFilters() *CT_PivotFilters {
	ret := &CT_PivotFilters{}
	return ret
}

func (m *CT_PivotFilters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.Filter != nil {
		sefilter := xml.StartElement{Name: xml.Name{Local: "ma:filter"}}
		for _, c := range m.Filter {
			e.EncodeElement(c, sefilter)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotFilters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PivotFilters:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "filter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "filter"}:
				tmp := NewCT_PivotFilter()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Filter = append(m.Filter, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_PivotFilters %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotFilters
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotFilters and its children
func (m *CT_PivotFilters) Validate() error {
	return m.ValidateWithPath("CT_PivotFilters")
}

// ValidateWithPath validates the CT_PivotFilters and its children, prefixing error messages with path
func (m *CT_PivotFilters) ValidateWithPath(path string) error {
	for i, v := range m.Filter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Filter[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
