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
	"strconv"

	"goffice"
)

type QueryTable struct {
	CT_QueryTable
}

func NewQueryTable() *QueryTable {
	ret := &QueryTable{}
	ret.CT_QueryTable = *NewCT_QueryTable()
	return ret
}

func (m *QueryTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:queryTable"
	return m.CT_QueryTable.MarshalXML(e, start)
}

func (m *QueryTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_QueryTable = *NewCT_QueryTable()
	for _, attr := range start.Attr {
		if attr.Name.Local == "adjustColumnWidth" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AdjustColumnWidthAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "intermediate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IntermediateAttr = &parsed
			continue
		}
		if attr.Name.Local == "rowNumbers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RowNumbersAttr = &parsed
			continue
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ConnectionIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "backgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
			continue
		}
		if attr.Name.Local == "fillFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FillFormulasAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyNumberFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyNumberFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyFontFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyFontFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstBackgroundRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FirstBackgroundRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoFormatId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AutoFormatIdAttr = &pt
			continue
		}
		if attr.Name.Local == "applyAlignmentFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyAlignmentFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "disableRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableRefreshAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyBorderFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyBorderFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "preserveFormatting" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveFormattingAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyPatternFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyPatternFormatsAttr = &parsed
			continue
		}
		if attr.Name.Local == "growShrinkType" {
			m.GrowShrinkTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "disableEdit" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisableEditAttr = &parsed
			continue
		}
		if attr.Name.Local == "headers" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadersAttr = &parsed
			continue
		}
		if attr.Name.Local == "removeDataOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RemoveDataOnSaveAttr = &parsed
			continue
		}
		if attr.Name.Local == "applyWidthHeightFormats" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ApplyWidthHeightFormatsAttr = &parsed
			continue
		}
	}
lQueryTable:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryTableRefresh"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "queryTableRefresh"}:
				m.QueryTableRefresh = NewCT_QueryTableRefresh()
				if err := d.DecodeElement(m.QueryTableRefresh, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on QueryTable %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lQueryTable
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the QueryTable and its children
func (m *QueryTable) Validate() error {
	return m.ValidateWithPath("QueryTable")
}

// ValidateWithPath validates the QueryTable and its children, prefixing error messages with path
func (m *QueryTable) ValidateWithPath(path string) error {
	if err := m.CT_QueryTable.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
