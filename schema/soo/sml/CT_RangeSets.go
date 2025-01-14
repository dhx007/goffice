package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_RangeSets struct {
	// Reference and Page Item Count
	CountAttr *uint32
	// Range Set
	RangeSet []*CT_RangeSet
}

func NewCT_RangeSets() *CT_RangeSets {
	ret := &CT_RangeSets{}
	return ret
}

func (m *CT_RangeSets) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	serangeSet := xml.StartElement{Name: xml.Name{Local: "ma:rangeSet"}}
	for _, c := range m.RangeSet {
		e.EncodeElement(c, serangeSet)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RangeSets) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_RangeSets:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rangeSet"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rangeSet"}:
				tmp := NewCT_RangeSet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RangeSet = append(m.RangeSet, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_RangeSets %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RangeSets
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RangeSets and its children
func (m *CT_RangeSets) Validate() error {
	return m.ValidateWithPath("CT_RangeSets")
}

// ValidateWithPath validates the CT_RangeSets and its children, prefixing error messages with path
func (m *CT_RangeSets) ValidateWithPath(path string) error {
	for i, v := range m.RangeSet {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RangeSet[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
