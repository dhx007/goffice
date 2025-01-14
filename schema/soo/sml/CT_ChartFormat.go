package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_ChartFormat struct {
	// Chart Index
	ChartAttr uint32
	// Pivot Format Id
	FormatAttr uint32
	// Series Format
	SeriesAttr *bool
	// Pivot Table Location Rule
	PivotArea *CT_PivotArea
}

func NewCT_ChartFormat() *CT_ChartFormat {
	ret := &CT_ChartFormat{}
	ret.PivotArea = NewCT_PivotArea()
	return ret
}

func (m *CT_ChartFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chart"},
		Value: fmt.Sprintf("%v", m.ChartAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "format"},
		Value: fmt.Sprintf("%v", m.FormatAttr)})
	if m.SeriesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "series"},
			Value: fmt.Sprintf("%d", b2i(*m.SeriesAttr))})
	}
	e.EncodeToken(start)
	sepivotArea := xml.StartElement{Name: xml.Name{Local: "ma:pivotArea"}}
	e.EncodeElement(m.PivotArea, sepivotArea)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PivotArea = NewCT_PivotArea()
	for _, attr := range start.Attr {
		if attr.Name.Local == "chart" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ChartAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "format" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.FormatAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "series" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SeriesAttr = &parsed
			continue
		}
	}
lCT_ChartFormat:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotArea"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotArea"}:
				if err := d.DecodeElement(m.PivotArea, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_ChartFormat %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartFormat
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartFormat and its children
func (m *CT_ChartFormat) Validate() error {
	return m.ValidateWithPath("CT_ChartFormat")
}

// ValidateWithPath validates the CT_ChartFormat and its children, prefixing error messages with path
func (m *CT_ChartFormat) ValidateWithPath(path string) error {
	if err := m.PivotArea.ValidateWithPath(path + "/PivotArea"); err != nil {
		return err
	}
	return nil
}
