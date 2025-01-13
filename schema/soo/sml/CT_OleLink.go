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

	"goffice"
)

type CT_OleLink struct {
	IdAttr string
	// Object Link Identifier
	ProgIdAttr string
	// Object Link Items
	OleItems *CT_OleItems
}

func NewCT_OleLink() *CT_OleLink {
	ret := &CT_OleLink{}
	return ret
}

func (m *CT_OleLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "progId"},
		Value: fmt.Sprintf("%v", m.ProgIdAttr)})
	e.EncodeToken(start)
	if m.OleItems != nil {
		seoleItems := xml.StartElement{Name: xml.Name{Local: "ma:oleItems"}}
		e.EncodeElement(m.OleItems, seoleItems)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "progId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIdAttr = parsed
			continue
		}
	}
lCT_OleLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleItems"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oleItems"}:
				m.OleItems = NewCT_OleItems()
				if err := d.DecodeElement(m.OleItems, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_OleLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleLink and its children
func (m *CT_OleLink) Validate() error {
	return m.ValidateWithPath("CT_OleLink")
}

// ValidateWithPath validates the CT_OleLink and its children, prefixing error messages with path
func (m *CT_OleLink) ValidateWithPath(path string) error {
	if m.OleItems != nil {
		if err := m.OleItems.ValidateWithPath(path + "/OleItems"); err != nil {
			return err
		}
	}
	return nil
}
