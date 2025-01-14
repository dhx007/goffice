package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_ConditionalFormat struct {
	// Conditional Formatting Scope
	ScopeAttr ST_Scope
	// Conditional Formatting Rule Type
	TypeAttr ST_Type
	// Priority
	PriorityAttr uint32
	// Pivot Areas
	PivotAreas *CT_PivotAreas
	ExtLst     *CT_ExtensionList
}

func NewCT_ConditionalFormat() *CT_ConditionalFormat {
	ret := &CT_ConditionalFormat{}
	ret.PivotAreas = NewCT_PivotAreas()
	return ret
}

func (m *CT_ConditionalFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ScopeAttr != ST_ScopeUnset {
		attr, err := m.ScopeAttr.MarshalXMLAttr(xml.Name{Local: "scope"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TypeAttr != ST_TypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "priority"},
		Value: fmt.Sprintf("%v", m.PriorityAttr)})
	e.EncodeToken(start)
	sepivotAreas := xml.StartElement{Name: xml.Name{Local: "ma:pivotAreas"}}
	e.EncodeElement(m.PivotAreas, sepivotAreas)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConditionalFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotAreas = NewCT_PivotAreas()
	for _, attr := range start.Attr {
		if attr.Name.Local == "scope" {
			m.ScopeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "priority" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PriorityAttr = uint32(parsed)
			continue
		}
	}
lCT_ConditionalFormat:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotAreas"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotAreas"}:
				if err := d.DecodeElement(m.PivotAreas, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_ConditionalFormat %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConditionalFormat
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConditionalFormat and its children
func (m *CT_ConditionalFormat) Validate() error {
	return m.ValidateWithPath("CT_ConditionalFormat")
}

// ValidateWithPath validates the CT_ConditionalFormat and its children, prefixing error messages with path
func (m *CT_ConditionalFormat) ValidateWithPath(path string) error {
	if err := m.ScopeAttr.ValidateWithPath(path + "/ScopeAttr"); err != nil {
		return err
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.PivotAreas.ValidateWithPath(path + "/PivotAreas"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
