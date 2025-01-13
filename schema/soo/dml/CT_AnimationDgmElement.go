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

	"goffice/schema/soo/ofc/sharedTypes"
)

type CT_AnimationDgmElement struct {
	IdAttr      *string
	BldStepAttr ST_DgmBuildStep
}

func NewCT_AnimationDgmElement() *CT_AnimationDgmElement {
	ret := &CT_AnimationDgmElement{}
	return ret
}

func (m *CT_AnimationDgmElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.BldStepAttr != ST_DgmBuildStepUnset {
		attr, err := m.BldStepAttr.MarshalXMLAttr(xml.Name{Local: "bldStep"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AnimationDgmElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "bldStep" {
			m.BldStepAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnimationDgmElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AnimationDgmElement and its children
func (m *CT_AnimationDgmElement) Validate() error {
	return m.ValidateWithPath("CT_AnimationDgmElement")
}

// ValidateWithPath validates the CT_AnimationDgmElement and its children, prefixing error messages with path
func (m *CT_AnimationDgmElement) ValidateWithPath(path string) error {
	if m.IdAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.IdAttr) {
			return fmt.Errorf(`%s/m.IdAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.IdAttr)
		}
	}
	if err := m.BldStepAttr.ValidateWithPath(path + "/BldStepAttr"); err != nil {
		return err
	}
	return nil
}
