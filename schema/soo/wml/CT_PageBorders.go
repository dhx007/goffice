package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type CT_PageBorders struct {
	// Z-Ordering of Page Border
	ZOrderAttr ST_PageBorderZOrder
	// Pages to Display Page Borders
	DisplayAttr ST_PageBorderDisplay
	// Page Border Positioning
	OffsetFromAttr ST_PageBorderOffset
	// Top Border
	Top *CT_TopPageBorder
	// Left Border
	Left *CT_PageBorder
	// Bottom Border
	Bottom *CT_BottomPageBorder
	// Right Border
	Right *CT_PageBorder
}

func NewCT_PageBorders() *CT_PageBorders {
	ret := &CT_PageBorders{}
	return ret
}

func (m *CT_PageBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ZOrderAttr != ST_PageBorderZOrderUnset {
		attr, err := m.ZOrderAttr.MarshalXMLAttr(xml.Name{Local: "w:zOrder"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DisplayAttr != ST_PageBorderDisplayUnset {
		attr, err := m.DisplayAttr.MarshalXMLAttr(xml.Name{Local: "w:display"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OffsetFromAttr != ST_PageBorderOffsetUnset {
		attr, err := m.OffsetFromAttr.MarshalXMLAttr(xml.Name{Local: "w:offsetFrom"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageBorders) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "zOrder" {
			m.ZOrderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "display" {
			m.DisplayAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "offsetFrom" {
			m.OffsetFromAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_PageBorders:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "top"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "top"}:
				m.Top = NewCT_TopPageBorder()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "left"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "left"}:
				m.Left = NewCT_PageBorder()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bottom"}:
				m.Bottom = NewCT_BottomPageBorder()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "right"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "right"}:
				m.Right = NewCT_PageBorder()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_PageBorders %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PageBorders
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PageBorders and its children
func (m *CT_PageBorders) Validate() error {
	return m.ValidateWithPath("CT_PageBorders")
}

// ValidateWithPath validates the CT_PageBorders and its children, prefixing error messages with path
func (m *CT_PageBorders) ValidateWithPath(path string) error {
	if err := m.ZOrderAttr.ValidateWithPath(path + "/ZOrderAttr"); err != nil {
		return err
	}
	if err := m.DisplayAttr.ValidateWithPath(path + "/DisplayAttr"); err != nil {
		return err
	}
	if err := m.OffsetFromAttr.ValidateWithPath(path + "/OffsetFromAttr"); err != nil {
		return err
	}
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	return nil
}
