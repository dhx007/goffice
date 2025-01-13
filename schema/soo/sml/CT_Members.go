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

type CT_Members struct {
	// Item Count
	CountAttr *uint32
	// Hierarchy Level
	LevelAttr *uint32
	// Member
	Member []*CT_Member
}

func NewCT_Members() *CT_Members {
	ret := &CT_Members{}
	return ret
}

func (m *CT_Members) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.LevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "level"},
			Value: fmt.Sprintf("%v", *m.LevelAttr)})
	}
	e.EncodeToken(start)
	semember := xml.StartElement{Name: xml.Name{Local: "ma:member"}}
	for _, c := range m.Member {
		e.EncodeElement(c, semember)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Members) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "level" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LevelAttr = &pt
			continue
		}
	}
lCT_Members:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "member"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "member"}:
				tmp := NewCT_Member()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Member = append(m.Member, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Members %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Members
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Members and its children
func (m *CT_Members) Validate() error {
	return m.ValidateWithPath("CT_Members")
}

// ValidateWithPath validates the CT_Members and its children, prefixing error messages with path
func (m *CT_Members) ValidateWithPath(path string) error {
	for i, v := range m.Member {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Member[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
