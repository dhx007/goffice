// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"fmt"

	"goffice"
	"goffice/schema/soo/ofc/sharedTypes"
	"goffice/schema/soo/wml"
)

type CT_Textbox struct {
	InsetAttr       *string
	SingleclickAttr sharedTypes.ST_TrueFalse
	InsetmodeAttr   OfcST_InsetMode
	TxbxContent     *wml.TxbxContent
	Any             unioffice.Any
	IdAttr          *string
	StyleAttr       *string
}

func NewCT_Textbox() *CT_Textbox {
	ret := &CT_Textbox{}
	return ret
}

func (m *CT_Textbox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.InsetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "inset"},
			Value: fmt.Sprintf("%v", *m.InsetAttr)})
	}
	if m.SingleclickAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.SingleclickAttr.MarshalXMLAttr(xml.Name{Local: "singleclick"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InsetmodeAttr != OfcST_InsetModeUnset {
		attr, err := m.InsetmodeAttr.MarshalXMLAttr(xml.Name{Local: "insetmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	e.EncodeToken(start)
	if m.TxbxContent != nil {
		setxbxContent := xml.StartElement{Name: xml.Name{Local: "w:txbxContent"}}
		e.EncodeElement(m.TxbxContent, setxbxContent)
	}
	if m.Any != nil {
		m.Any.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Textbox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "singleclick" {
			m.SingleclickAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "inset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InsetAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
	}
lCT_Textbox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "txbxContent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "txbxContent"}:
				m.TxbxContent = wml.NewTxbxContent()
				if err := d.DecodeElement(m.TxbxContent, &el); err != nil {
					return err
				}
			default:
				if anyEl, err := unioffice.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lCT_Textbox
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Textbox and its children
func (m *CT_Textbox) Validate() error {
	return m.ValidateWithPath("CT_Textbox")
}

// ValidateWithPath validates the CT_Textbox and its children, prefixing error messages with path
func (m *CT_Textbox) ValidateWithPath(path string) error {
	if err := m.SingleclickAttr.ValidateWithPath(path + "/SingleclickAttr"); err != nil {
		return err
	}
	if err := m.InsetmodeAttr.ValidateWithPath(path + "/InsetmodeAttr"); err != nil {
		return err
	}
	if m.TxbxContent != nil {
		if err := m.TxbxContent.ValidateWithPath(path + "/TxbxContent"); err != nil {
			return err
		}
	}
	return nil
}
