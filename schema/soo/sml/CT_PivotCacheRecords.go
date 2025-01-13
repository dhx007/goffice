// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"goffice"
)

type CT_PivotCacheRecords struct {
	// PivotCache Records Count
	CountAttr *uint32
	// PivotCache Record
	R []*CT_Record
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotCacheRecords() *CT_PivotCacheRecords {
	ret := &CT_PivotCacheRecords{}
	return ret
}

func (m *CT_PivotCacheRecords) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "ma:r"}}
		for _, c := range m.R {
			e.EncodeElement(c, ser)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotCacheRecords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PivotCacheRecords:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "r"}:
				tmp := NewCT_Record()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.R = append(m.R, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_PivotCacheRecords %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotCacheRecords
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotCacheRecords and its children
func (m *CT_PivotCacheRecords) Validate() error {
	return m.ValidateWithPath("CT_PivotCacheRecords")
}

// ValidateWithPath validates the CT_PivotCacheRecords and its children, prefixing error messages with path
func (m *CT_PivotCacheRecords) ValidateWithPath(path string) error {
	for i, v := range m.R {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/R[%d]", path, i)); err != nil {
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
