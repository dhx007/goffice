package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_CacheSource struct {
	// Cache Type
	TypeAttr ST_SourceType
	// Connection Index
	ConnectionIdAttr *uint32
	// Worksheet PivotCache Source
	WorksheetSource *CT_WorksheetSource
	// Consolidation Source
	Consolidation *CT_Consolidation
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CacheSource() *CT_CacheSource {
	ret := &CT_CacheSource{}
	ret.TypeAttr = ST_SourceType(1)
	return ret
}

func (m *CT_CacheSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.ConnectionIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
			Value: fmt.Sprintf("%v", *m.ConnectionIdAttr)})
	}
	e.EncodeToken(start)
	if m.WorksheetSource != nil {
		seworksheetSource := xml.StartElement{Name: xml.Name{Local: "ma:worksheetSource"}}
		e.EncodeElement(m.WorksheetSource, seworksheetSource)
	}
	if m.Consolidation != nil {
		seconsolidation := xml.StartElement{Name: xml.Name{Local: "ma:consolidation"}}
		e.EncodeElement(m.Consolidation, seconsolidation)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CacheSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_SourceType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ConnectionIdAttr = &pt
			continue
		}
	}
lCT_CacheSource:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "worksheetSource"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "worksheetSource"}:
				m.WorksheetSource = NewCT_WorksheetSource()
				if err := d.DecodeElement(m.WorksheetSource, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "consolidation"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "consolidation"}:
				m.Consolidation = NewCT_Consolidation()
				if err := d.DecodeElement(m.Consolidation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_CacheSource %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CacheSource
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CacheSource and its children
func (m *CT_CacheSource) Validate() error {
	return m.ValidateWithPath("CT_CacheSource")
}

// ValidateWithPath validates the CT_CacheSource and its children, prefixing error messages with path
func (m *CT_CacheSource) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_SourceTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.WorksheetSource != nil {
		if err := m.WorksheetSource.ValidateWithPath(path + "/WorksheetSource"); err != nil {
			return err
		}
	}
	if m.Consolidation != nil {
		if err := m.Consolidation.ValidateWithPath(path + "/Consolidation"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
