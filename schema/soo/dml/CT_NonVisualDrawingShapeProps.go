// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"goffice"
)

type CT_NonVisualDrawingShapeProps struct {
	TxBoxAttr *bool
	SpLocks   *CT_ShapeLocking
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_NonVisualDrawingShapeProps() *CT_NonVisualDrawingShapeProps {
	ret := &CT_NonVisualDrawingShapeProps{}
	return ret
}

func (m *CT_NonVisualDrawingShapeProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TxBoxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "txBox"},
			Value: fmt.Sprintf("%d", b2i(*m.TxBoxAttr))})
	}
	e.EncodeToken(start)
	if m.SpLocks != nil {
		sespLocks := xml.StartElement{Name: xml.Name{Local: "a:spLocks"}}
		e.EncodeElement(m.SpLocks, sespLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NonVisualDrawingShapeProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "txBox" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TxBoxAttr = &parsed
			continue
		}
	}
lCT_NonVisualDrawingShapeProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spLocks"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spLocks"}:
				m.SpLocks = NewCT_ShapeLocking()
				if err := d.DecodeElement(m.SpLocks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_NonVisualDrawingShapeProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualDrawingShapeProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NonVisualDrawingShapeProps and its children
func (m *CT_NonVisualDrawingShapeProps) Validate() error {
	return m.ValidateWithPath("CT_NonVisualDrawingShapeProps")
}

// ValidateWithPath validates the CT_NonVisualDrawingShapeProps and its children, prefixing error messages with path
func (m *CT_NonVisualDrawingShapeProps) ValidateWithPath(path string) error {
	if m.SpLocks != nil {
		if err := m.SpLocks.ValidateWithPath(path + "/SpLocks"); err != nil {
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
