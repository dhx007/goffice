// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"goffice/schema/soo/ofc/sharedTypes"
)

type CT_TwipsMeasure struct {
	// Measurement in Twentieths of a Point
	ValAttr sharedTypes.ST_TwipsMeasure
}

func NewCT_TwipsMeasure() *CT_TwipsMeasure {
	ret := &CT_TwipsMeasure{}
	return ret
}

func (m *CT_TwipsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TwipsMeasure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TwipsMeasure: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TwipsMeasure and its children
func (m *CT_TwipsMeasure) Validate() error {
	return m.ValidateWithPath("CT_TwipsMeasure")
}

// ValidateWithPath validates the CT_TwipsMeasure and its children, prefixing error messages with path
func (m *CT_TwipsMeasure) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	return nil
}
