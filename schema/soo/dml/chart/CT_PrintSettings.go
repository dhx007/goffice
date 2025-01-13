// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"goffice"
)

type CT_PrintSettings struct {
	HeaderFooter    *CT_HeaderFooter
	PageMargins     *CT_PageMargins
	PageSetup       *CT_PageSetup
	LegacyDrawingHF *CT_RelId
}

func NewCT_PrintSettings() *CT_PrintSettings {
	ret := &CT_PrintSettings{}
	return ret
}

func (m *CT_PrintSettings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "c:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "c:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "c:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.LegacyDrawingHF != nil {
		selegacyDrawingHF := xml.StartElement{Name: xml.Name{Local: "c:legacyDrawingHF"}}
		e.EncodeElement(m.LegacyDrawingHF, selegacyDrawingHF)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PrintSettings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PrintSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "headerFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "headerFooter"}:
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pageMargins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pageMargins"}:
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pageSetup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pageSetup"}:
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "legacyDrawingHF"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "legacyDrawingHF"}:
				m.LegacyDrawingHF = NewCT_RelId()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_PrintSettings %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PrintSettings
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PrintSettings and its children
func (m *CT_PrintSettings) Validate() error {
	return m.ValidateWithPath("CT_PrintSettings")
}

// ValidateWithPath validates the CT_PrintSettings and its children, prefixing error messages with path
func (m *CT_PrintSettings) ValidateWithPath(path string) error {
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.LegacyDrawingHF != nil {
		if err := m.LegacyDrawingHF.ValidateWithPath(path + "/LegacyDrawingHF"); err != nil {
			return err
		}
	}
	return nil
}
