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

type CT_InnerShadowEffect struct {
	BlurRadAttr *int64
	DistAttr    *int64
	DirAttr     *int32
	ScrgbClr    *CT_ScRgbColor
	SrgbClr     *CT_SRgbColor
	HslClr      *CT_HslColor
	SysClr      *CT_SystemColor
	SchemeClr   *CT_SchemeColor
	PrstClr     *CT_PresetColor
}

func NewCT_InnerShadowEffect() *CT_InnerShadowEffect {
	ret := &CT_InnerShadowEffect{}
	return ret
}

func (m *CT_InnerShadowEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BlurRadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blurRad"},
			Value: fmt.Sprintf("%v", *m.BlurRadAttr)})
	}
	if m.DistAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dist"},
			Value: fmt.Sprintf("%v", *m.DistAttr)})
	}
	if m.DirAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dir"},
			Value: fmt.Sprintf("%v", *m.DirAttr)})
	}
	e.EncodeToken(start)
	if m.ScrgbClr != nil {
		sescrgbClr := xml.StartElement{Name: xml.Name{Local: "a:scrgbClr"}}
		e.EncodeElement(m.ScrgbClr, sescrgbClr)
	}
	if m.SrgbClr != nil {
		sesrgbClr := xml.StartElement{Name: xml.Name{Local: "a:srgbClr"}}
		e.EncodeElement(m.SrgbClr, sesrgbClr)
	}
	if m.HslClr != nil {
		sehslClr := xml.StartElement{Name: xml.Name{Local: "a:hslClr"}}
		e.EncodeElement(m.HslClr, sehslClr)
	}
	if m.SysClr != nil {
		sesysClr := xml.StartElement{Name: xml.Name{Local: "a:sysClr"}}
		e.EncodeElement(m.SysClr, sesysClr)
	}
	if m.SchemeClr != nil {
		seschemeClr := xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}}
		e.EncodeElement(m.SchemeClr, seschemeClr)
	}
	if m.PrstClr != nil {
		seprstClr := xml.StartElement{Name: xml.Name{Local: "a:prstClr"}}
		e.EncodeElement(m.PrstClr, seprstClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_InnerShadowEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "blurRad" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.BlurRadAttr = &parsed
			continue
		}
		if attr.Name.Local == "dist" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DistAttr = &parsed
			continue
		}
		if attr.Name.Local == "dir" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.DirAttr = &pt
			continue
		}
	}
lCT_InnerShadowEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scrgbClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "scrgbClr"}:
				m.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(m.ScrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "srgbClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "srgbClr"}:
				m.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(m.SrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hslClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "hslClr"}:
				m.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(m.HslClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sysClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "sysClr"}:
				m.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(m.SysClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "schemeClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "schemeClr"}:
				m.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(m.SchemeClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstClr"}:
				m.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(m.PrstClr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_InnerShadowEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_InnerShadowEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_InnerShadowEffect and its children
func (m *CT_InnerShadowEffect) Validate() error {
	return m.ValidateWithPath("CT_InnerShadowEffect")
}

// ValidateWithPath validates the CT_InnerShadowEffect and its children, prefixing error messages with path
func (m *CT_InnerShadowEffect) ValidateWithPath(path string) error {
	if m.BlurRadAttr != nil {
		if *m.BlurRadAttr < 0 {
			return fmt.Errorf("%s/m.BlurRadAttr must be >= 0 (have %v)", path, *m.BlurRadAttr)
		}
		if *m.BlurRadAttr > 27273042316900 {
			return fmt.Errorf("%s/m.BlurRadAttr must be <= 27273042316900 (have %v)", path, *m.BlurRadAttr)
		}
	}
	if m.DistAttr != nil {
		if *m.DistAttr < 0 {
			return fmt.Errorf("%s/m.DistAttr must be >= 0 (have %v)", path, *m.DistAttr)
		}
		if *m.DistAttr > 27273042316900 {
			return fmt.Errorf("%s/m.DistAttr must be <= 27273042316900 (have %v)", path, *m.DistAttr)
		}
	}
	if m.DirAttr != nil {
		if *m.DirAttr < 0 {
			return fmt.Errorf("%s/m.DirAttr must be >= 0 (have %v)", path, *m.DirAttr)
		}
		if *m.DirAttr >= 21600000 {
			return fmt.Errorf("%s/m.DirAttr must be < 21600000 (have %v)", path, *m.DirAttr)
		}
	}
	if m.ScrgbClr != nil {
		if err := m.ScrgbClr.ValidateWithPath(path + "/ScrgbClr"); err != nil {
			return err
		}
	}
	if m.SrgbClr != nil {
		if err := m.SrgbClr.ValidateWithPath(path + "/SrgbClr"); err != nil {
			return err
		}
	}
	if m.HslClr != nil {
		if err := m.HslClr.ValidateWithPath(path + "/HslClr"); err != nil {
			return err
		}
	}
	if m.SysClr != nil {
		if err := m.SysClr.ValidateWithPath(path + "/SysClr"); err != nil {
			return err
		}
	}
	if m.SchemeClr != nil {
		if err := m.SchemeClr.ValidateWithPath(path + "/SchemeClr"); err != nil {
			return err
		}
	}
	if m.PrstClr != nil {
		if err := m.PrstClr.ValidateWithPath(path + "/PrstClr"); err != nil {
			return err
		}
	}
	return nil
}
