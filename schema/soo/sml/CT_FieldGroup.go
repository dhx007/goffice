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

type CT_FieldGroup struct {
	// Parent
	ParAttr *uint32
	// Field Base
	BaseAttr *uint32
	// Range Grouping Properties
	RangePr *CT_RangePr
	// Discrete Grouping Properties
	DiscretePr *CT_DiscretePr
	// OLAP Group Items
	GroupItems *CT_GroupItems
}

func NewCT_FieldGroup() *CT_FieldGroup {
	ret := &CT_FieldGroup{}
	return ret
}

func (m *CT_FieldGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ParAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "par"},
			Value: fmt.Sprintf("%v", *m.ParAttr)})
	}
	if m.BaseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "base"},
			Value: fmt.Sprintf("%v", *m.BaseAttr)})
	}
	e.EncodeToken(start)
	if m.RangePr != nil {
		serangePr := xml.StartElement{Name: xml.Name{Local: "ma:rangePr"}}
		e.EncodeElement(m.RangePr, serangePr)
	}
	if m.DiscretePr != nil {
		sediscretePr := xml.StartElement{Name: xml.Name{Local: "ma:discretePr"}}
		e.EncodeElement(m.DiscretePr, sediscretePr)
	}
	if m.GroupItems != nil {
		segroupItems := xml.StartElement{Name: xml.Name{Local: "ma:groupItems"}}
		e.EncodeElement(m.GroupItems, segroupItems)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FieldGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "par" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ParAttr = &pt
			continue
		}
		if attr.Name.Local == "base" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BaseAttr = &pt
			continue
		}
	}
lCT_FieldGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rangePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rangePr"}:
				m.RangePr = NewCT_RangePr()
				if err := d.DecodeElement(m.RangePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "discretePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "discretePr"}:
				m.DiscretePr = NewCT_DiscretePr()
				if err := d.DecodeElement(m.DiscretePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "groupItems"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "groupItems"}:
				m.GroupItems = NewCT_GroupItems()
				if err := d.DecodeElement(m.GroupItems, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_FieldGroup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FieldGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FieldGroup and its children
func (m *CT_FieldGroup) Validate() error {
	return m.ValidateWithPath("CT_FieldGroup")
}

// ValidateWithPath validates the CT_FieldGroup and its children, prefixing error messages with path
func (m *CT_FieldGroup) ValidateWithPath(path string) error {
	if m.RangePr != nil {
		if err := m.RangePr.ValidateWithPath(path + "/RangePr"); err != nil {
			return err
		}
	}
	if m.DiscretePr != nil {
		if err := m.DiscretePr.ValidateWithPath(path + "/DiscretePr"); err != nil {
			return err
		}
	}
	if m.GroupItems != nil {
		if err := m.GroupItems.ValidateWithPath(path + "/GroupItems"); err != nil {
			return err
		}
	}
	return nil
}
