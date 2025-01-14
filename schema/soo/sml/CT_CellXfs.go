package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/dhx007/goffice"
)

type CT_CellXfs struct {
	// Format Count
	CountAttr *uint32
	// Format
	Xf []*CT_Xf
}

func NewCT_CellXfs() *CT_CellXfs {
	ret := &CT_CellXfs{}
	return ret
}

func (m *CT_CellXfs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sexf := xml.StartElement{Name: xml.Name{Local: "ma:xf"}}
	for _, c := range m.Xf {
		e.EncodeElement(c, sexf)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellXfs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_CellXfs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "xf"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "xf"}:
				tmp := NewCT_Xf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Xf = append(m.Xf, tmp)
			default:
				goffice.Log("skipping unsupported element on CT_CellXfs %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellXfs
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellXfs and its children
func (m *CT_CellXfs) Validate() error {
	return m.ValidateWithPath("CT_CellXfs")
}

// ValidateWithPath validates the CT_CellXfs and its children, prefixing error messages with path
func (m *CT_CellXfs) ValidateWithPath(path string) error {
	for i, v := range m.Xf {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Xf[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
