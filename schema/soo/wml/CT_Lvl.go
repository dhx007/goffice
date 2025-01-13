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
	"strconv"

	"goffice"
	"goffice/schema/soo/ofc/sharedTypes"
)

type CT_Lvl struct {
	// Numbering Level
	IlvlAttr int64
	// Template Code
	TplcAttr *string
	// Tentative Numbering
	TentativeAttr *sharedTypes.ST_OnOff
	// Starting Value
	Start *CT_DecimalNumber
	// Numbering Format
	NumFmt *CT_NumFmt
	// Restart Numbering Level Symbol
	LvlRestart *CT_DecimalNumber
	// Paragraph Style's Associated Numbering Level
	PStyle *CT_String
	// Display All Levels Using Arabic Numerals
	IsLgl *CT_OnOff
	// Content Between Numbering Symbol and Paragraph Text
	Suff *CT_LevelSuffix
	// Numbering Level Text
	LvlText *CT_LevelText
	// Picture Numbering Symbol Definition Reference
	LvlPicBulletId *CT_DecimalNumber
	// Legacy Numbering Level Properties
	Legacy *CT_LvlLegacy
	// Justification
	LvlJc *CT_Jc
	// Numbering Level Associated Paragraph Properties
	PPr *CT_PPrGeneral
	// Numbering Symbol Run Properties
	RPr *CT_RPr
}

func NewCT_Lvl() *CT_Lvl {
	ret := &CT_Lvl{}
	return ret
}

func (m *CT_Lvl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:ilvl"},
		Value: fmt.Sprintf("%v", m.IlvlAttr)})
	if m.TplcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tplc"},
			Value: fmt.Sprintf("%v", *m.TplcAttr)})
	}
	if m.TentativeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tentative"},
			Value: fmt.Sprintf("%v", *m.TentativeAttr)})
	}
	e.EncodeToken(start)
	if m.Start != nil {
		sestart := xml.StartElement{Name: xml.Name{Local: "w:start"}}
		e.EncodeElement(m.Start, sestart)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "w:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.LvlRestart != nil {
		selvlRestart := xml.StartElement{Name: xml.Name{Local: "w:lvlRestart"}}
		e.EncodeElement(m.LvlRestart, selvlRestart)
	}
	if m.PStyle != nil {
		sepStyle := xml.StartElement{Name: xml.Name{Local: "w:pStyle"}}
		e.EncodeElement(m.PStyle, sepStyle)
	}
	if m.IsLgl != nil {
		seisLgl := xml.StartElement{Name: xml.Name{Local: "w:isLgl"}}
		e.EncodeElement(m.IsLgl, seisLgl)
	}
	if m.Suff != nil {
		sesuff := xml.StartElement{Name: xml.Name{Local: "w:suff"}}
		e.EncodeElement(m.Suff, sesuff)
	}
	if m.LvlText != nil {
		selvlText := xml.StartElement{Name: xml.Name{Local: "w:lvlText"}}
		e.EncodeElement(m.LvlText, selvlText)
	}
	if m.LvlPicBulletId != nil {
		selvlPicBulletId := xml.StartElement{Name: xml.Name{Local: "w:lvlPicBulletId"}}
		e.EncodeElement(m.LvlPicBulletId, selvlPicBulletId)
	}
	if m.Legacy != nil {
		selegacy := xml.StartElement{Name: xml.Name{Local: "w:legacy"}}
		e.EncodeElement(m.Legacy, selegacy)
	}
	if m.LvlJc != nil {
		selvlJc := xml.StartElement{Name: xml.Name{Local: "w:lvlJc"}}
		e.EncodeElement(m.LvlJc, selvlJc)
	}
	if m.PPr != nil {
		sepPr := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}
		e.EncodeElement(m.PPr, sepPr)
	}
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Lvl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tplc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TplcAttr = &parsed
			continue
		}
		if attr.Name.Local == "ilvl" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IlvlAttr = parsed
			continue
		}
		if attr.Name.Local == "tentative" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.TentativeAttr = &parsed
			continue
		}
	}
lCT_Lvl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "start"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "start"}:
				m.Start = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "numFmt"}:
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lvlRestart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lvlRestart"}:
				m.LvlRestart = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.LvlRestart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pStyle"}:
				m.PStyle = NewCT_String()
				if err := d.DecodeElement(m.PStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "isLgl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "isLgl"}:
				m.IsLgl = NewCT_OnOff()
				if err := d.DecodeElement(m.IsLgl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "suff"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "suff"}:
				m.Suff = NewCT_LevelSuffix()
				if err := d.DecodeElement(m.Suff, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lvlText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lvlText"}:
				m.LvlText = NewCT_LevelText()
				if err := d.DecodeElement(m.LvlText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lvlPicBulletId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lvlPicBulletId"}:
				m.LvlPicBulletId = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.LvlPicBulletId, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "legacy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "legacy"}:
				m.Legacy = NewCT_LvlLegacy()
				if err := d.DecodeElement(m.Legacy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lvlJc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lvlJc"}:
				m.LvlJc = NewCT_Jc()
				if err := d.DecodeElement(m.LvlJc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pPr"}:
				m.PPr = NewCT_PPrGeneral()
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Lvl %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Lvl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Lvl and its children
func (m *CT_Lvl) Validate() error {
	return m.ValidateWithPath("CT_Lvl")
}

// ValidateWithPath validates the CT_Lvl and its children, prefixing error messages with path
func (m *CT_Lvl) ValidateWithPath(path string) error {
	if m.TentativeAttr != nil {
		if err := m.TentativeAttr.ValidateWithPath(path + "/TentativeAttr"); err != nil {
			return err
		}
	}
	if m.Start != nil {
		if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.LvlRestart != nil {
		if err := m.LvlRestart.ValidateWithPath(path + "/LvlRestart"); err != nil {
			return err
		}
	}
	if m.PStyle != nil {
		if err := m.PStyle.ValidateWithPath(path + "/PStyle"); err != nil {
			return err
		}
	}
	if m.IsLgl != nil {
		if err := m.IsLgl.ValidateWithPath(path + "/IsLgl"); err != nil {
			return err
		}
	}
	if m.Suff != nil {
		if err := m.Suff.ValidateWithPath(path + "/Suff"); err != nil {
			return err
		}
	}
	if m.LvlText != nil {
		if err := m.LvlText.ValidateWithPath(path + "/LvlText"); err != nil {
			return err
		}
	}
	if m.LvlPicBulletId != nil {
		if err := m.LvlPicBulletId.ValidateWithPath(path + "/LvlPicBulletId"); err != nil {
			return err
		}
	}
	if m.Legacy != nil {
		if err := m.Legacy.ValidateWithPath(path + "/Legacy"); err != nil {
			return err
		}
	}
	if m.LvlJc != nil {
		if err := m.LvlJc.ValidateWithPath(path + "/LvlJc"); err != nil {
			return err
		}
	}
	if m.PPr != nil {
		if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
			return err
		}
	}
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
