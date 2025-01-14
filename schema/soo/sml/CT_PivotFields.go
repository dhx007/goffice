package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_PivotFields struct {
	// Field Count
	CountAttr *uint32
	// PivotTable Field
	PivotField []*CT_PivotField
}

func NewCT_PivotFields() *CT_PivotFields {
	ret := &CT_PivotFields{}
	return ret
}

func (m *CT_PivotFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sepivotField := xml.StartElement{Name: xml.Name{Local: "ma:pivotField"}}
	for _, c := range m.PivotField {
		e.EncodeElement(c, sepivotField)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PivotFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotField"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pivotField"}:
				tmp := NewCT_PivotField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotField = append(m.PivotField, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_PivotFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotFields and its children
func (m *CT_PivotFields) Validate() error {
	return m.ValidateWithPath("CT_PivotFields")
}

// ValidateWithPath validates the CT_PivotFields and its children, prefixing error messages with path
func (m *CT_PivotFields) ValidateWithPath(path string) error {
	for i, v := range m.PivotField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
